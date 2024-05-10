package helloworld

import (
	"main/internal/delivery/rest"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler() rest.API {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/hello", rest.HandlerFunc(h.hello).Serve).Methods(http.MethodGet)
}
