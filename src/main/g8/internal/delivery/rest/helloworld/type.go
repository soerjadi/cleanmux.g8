package helloworld

import "github.com/gorilla/mux"

type Handler struct {
	middleware []mux.MiddlewareFunc
}
