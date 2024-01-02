package main

import (
	"net/http"

	"github.com/rajielijah/go-server/controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe("localhost:3000", mux)
}
