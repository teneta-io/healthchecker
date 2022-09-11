package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/teneta-io/healthchecker/pkg/handler"
)

type BenchmarkHandler struct {
	api *api
}

func (h *BenchmarkHandler) StressTest() handler.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		stresstest := h.api.stresser.LastResult()

		return json.NewEncoder(w).Encode(stresstest)
	}
}
