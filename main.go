package main

import (
	"fmt"
	"net/http"

	"github.com/rashad-arbab-convictional-engineering-interview/internal/router"
)

func main() {
	port := 8080
	Server := router.Server{
		Port: port,
	}
	fmt.Printf("starting server on port: %d \n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), Server.NewRouter())
	if err != nil {
		panic(err)
	}

}
