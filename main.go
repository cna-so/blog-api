package main

import (
	"backend/cmd/api"
	"log"
	"net/http"

	"backend/initializer"
)

func init() {
	initializer.ConnectToDb()
}

const port = ":8080"

func main() {
	server := http.Server{
		Addr:    port,
		Handler: api.Routes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
