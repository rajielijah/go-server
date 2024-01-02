// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func hellohandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 Not Found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "Method not suppported", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprintf(w, "Hello Motherfucker")

// }

// func formhandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST request Successful\n")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")
// 	fmt.Fprintf(w, "Name = %s\n", name)
// 	fmt.Fprintf(w, "Address = %s\n", address)

// }

// func main() {
// 	fileServer := http.FileServer(http.Dir("./static"))
// 	http.Handle("/", fileServer)
// 	http.HandleFunc("/hello", hellohandler)
// 	http.HandleFunc("/form", formhandler)

// 	fmt.Printf("starting server at port 8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}

// }
