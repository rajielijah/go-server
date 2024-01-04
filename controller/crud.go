package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rajielijah/go-server/model"
	"github.com/rajielijah/go-server/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				fmt.Println(err)
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			// name := r.URL.Query().Get("name")
			data, err := model.ReadAll()
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("Somme error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{"ITEM Deleted"})
		}
	}
}
