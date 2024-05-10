package rest

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)

func (fn HandlerFunc) Serve(w http.ResponseWriter, r *http.Request) {
	response := Response{}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	data, err := fn(w, r)

	if data != nil && err == nil {
		response.Data = data
		response.Status = "success"

		if buff, err := json.Marshal(response); err == nil {
			_, err := w.Write(buff)
			if err != nil {
				return
			}
		}
	}

	if err != nil {
		response.Data = data
		response.Message = err.Error()
		response.Status = "error"
	} else {
		return
	}

	buf, _ := json.Marshal(response)

	_, _ = w.Write(buf)
}
