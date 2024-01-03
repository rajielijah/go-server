package main

import (
	"net/http"

	"github.com/rajielijah/go-server/controller"
)

func main() {
	mux := controller.Register()
	db := model.connect()
	defer db.Close()
	http.ListenAndServe("localhost:3000", mux)
}
