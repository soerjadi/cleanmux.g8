package helloworld

import "net/http"

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "hello world", nil
}
