package apiserver

import (
	"fmt"

	"github.com/gorilla/mux"
)

func configureRoutes(a *api) *mux.Router {
	router := mux.NewRouter()

	route := fmt.Sprintf("/%s", a.keystore.GetPublicKey())
	router.Handle(route, a.Benchmark().StressTest())

	return router
}
