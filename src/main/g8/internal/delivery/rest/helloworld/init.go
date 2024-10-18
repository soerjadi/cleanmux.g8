package helloworld

import (
	"$module_name$/internal/delivery/rest"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(middleware []mux.MiddlewareFunc) rest.API {
	return &Handler{
		middleware: middleware,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	if h.middleware != nil {
		r.Use(h.middleware...)
	}
	r.HandleFunc("/hello", rest.HandlerFunc(h.hello).Serve).Methods(http.MethodGet)
}
