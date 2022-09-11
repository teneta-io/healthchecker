package handler

import (
	"encoding/json"
	"net/http"

	"github.com/teneta-io/healthchecker/pkg/logger"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

type Wrap struct {
	Handler
}

func (wrap Wrap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := wrap.Handle(w, r); err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific HTTP status code.
			w.WriteHeader(e.Status())
			if err := json.NewEncoder(w).Encode(e); err != nil {
				panic(err)
			}
		default:
			// Any error types we don't specifically look out for default to serving a HTTP 50
			logger.Errorf("unhealthy: %s", err)
			w.WriteHeader(ErrUnhealthy.Status())
			if err := json.NewEncoder(w).Encode(ErrUnhealthy); err != nil {
				panic(err)
			}
		}
	}
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (f HandlerFunc) Handle(w http.ResponseWriter, r *http.Request) error {
	return f(w, r)
}

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Wrap{f}.ServeHTTP(w, r)
}
