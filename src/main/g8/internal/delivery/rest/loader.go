package rest

import "github.com/gorilla/mux"

func RegisterHandlers(r *mux.Router, handlers ...API) {
	for i := (0); i < len(handlers); i++ {
		handlers[i].RegisterRoutes(r)
	}
}
