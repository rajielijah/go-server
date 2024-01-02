package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rajielijah/go-server/views"
)

func ping() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)

			fmt.Println("Request Recieed")
			w.Write([]byte("We are here"))
		}
	}
}
