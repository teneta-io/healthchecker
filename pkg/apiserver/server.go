package apiserver

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/teneta-io/healthchecker/pkg/config"
	"github.com/teneta-io/healthchecker/pkg/handler"
	"github.com/teneta-io/healthchecker/pkg/keystore"
	"github.com/teneta-io/healthchecker/pkg/logger"
	"github.com/teneta-io/healthchecker/pkg/stresser"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server represent main server configurations
type Server struct {
	*http.Server
}

type api struct {
	router   *mux.Router
	keystore keystore.KeyStore
	stresser stresser.StressTester

	benchmarkAPI *BenchmarkHandler
}

// NewServer returns new http server
func NewServer(
	addr string,
	conf *config.Server,
	keystore keystore.KeyStore,
	stresser stresser.StressTester,
) *Server {
	handler := newAPI(keystore, stresser)

	srv := &http.Server{
		Addr:           addr,
		Handler:        handlers.CustomLoggingHandler(os.Stdout, handler, logFormatter),
		ReadTimeout:    conf.ReadTimeout.Duration,
		WriteTimeout:   conf.WriteTimeout.Duration,
		IdleTimeout:    conf.IdleTimeout.Duration,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		Server: srv,
	}
}

func logFormatter(w io.Writer, params handlers.LogFormatterParams) {
	logger.WithFields(logger.Fields{
		"method": params.Request.Method,
		"path":   params.URL.Path,
		"status": params.StatusCode,
		"size":   params.Size,
		"addr":   params.Request.RemoteAddr,
	}).Infof("http")
}

func newAPI(
	keystore keystore.KeyStore,
	stresser stresser.StressTester,
) *api {
	api := &api{
		keystore: keystore,
		stresser: stresser,
		router:   mux.NewRouter(),
	}

	api.router = configureRoutes(api)

	return api
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin"})

	cors := handlers.CORS(origins, methods, headers)(a.router)
	cors.ServeHTTP(w, r)
}

func (*api) HealthCheck() handler.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := json.NewEncoder(w).Encode("Success"); err != nil {
			return err
		}

		return nil
	}
}

func (a *api) Benchmark() *BenchmarkHandler {
	if a.benchmarkAPI == nil {
		a.benchmarkAPI = &BenchmarkHandler{
			api: a,
		}
	}

	return a.benchmarkAPI
}
