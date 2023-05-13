package main

import (
	"backend/controllers"
	"backend/initializer"
	"net/http"
)

func init() {
	initializer.ConnectToDb()
}

const port = ":8080"

func main() {
	server := http.Server{
		Addr:    port,
		Handler: controllers.Routes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
