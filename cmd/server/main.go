package main

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tmaffia/pkg/handlers"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /", handlers.HomeHandler)
	r.HandleFunc("GET /health", handlers.HealthCheckHandler)


	log.Println("API is running!")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Fatalf("could not start server: %v\n", err)
	}
}
