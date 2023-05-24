package main

import (
	"fmt"
	"net/http"
	"os"
	"rentals/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := makeDBConnection()
	if err != nil {
		panic(err)
	}

	router := makeRouter(db)
	http.ListenAndServe(":8080", router)
}

func makeRouter(db gorm.DB) chi.Router {
	r := chi.NewRouter()
	rs := routes.PgHandle{DB: &db}
	r.Use(middleware.Logger)

	r.Get("/rentals", rs.GetRentals)
	r.Get("/rentals/{id}", rs.GetRental)

	return r
}

func makeDBConnection() (gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	return *db, err
}
