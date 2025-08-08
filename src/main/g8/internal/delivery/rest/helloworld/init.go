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
	route := r.NewRoute().Name("hello.world").SubRouter()

	route.HandleFunc("/hello", rest.HandlerFunc(h.hello).Serve).Methods(http.MethodGet)

	if h.middleware != nil {
		route.Use(h.middleware...)
	}
}
