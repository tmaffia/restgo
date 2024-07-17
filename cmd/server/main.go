package main

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tmaffia/pkg/handlers"
)


func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.HomeHandler)
	r.Get("/health", handlers.HealthCheckHandler)

	log.Println("API is running!")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Fatalf("could not start server: %v\n", err)
	}
}