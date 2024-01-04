package main

import (
	"fmt"
	"net/http"

	"github.com/rajielijah/go-server/controller"
	"github.com/rajielijah/go-server/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("serving")
	http.ListenAndServe("localhost:3000", mux)
}
