package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=127.0.0.1 user=root password=root dbname=testingwithrentals port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	println(db, err)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/rentals", getRentals)
	router.Get("/rentals/{id}", getRental)

	http.ListenAndServe(":8080", router)
}

func getRental(w http.ResponseWriter, r *http.Request) {}

func getRentals(w http.ResponseWriter, r *http.Request) {}
