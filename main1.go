package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Recieed")
		w.Write([]byte("We are here"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
