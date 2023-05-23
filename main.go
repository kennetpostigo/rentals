package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/rentals", getRentals)
	router.Get("/rentals/{id}", getRental)

	http.ListenAndServe(":8080", router)
}

func getRental(w http.ResponseWriter, r *http.Request) {}

func getRentals(w http.ResponseWriter, r *http.Request) {}
