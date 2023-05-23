package routes

import (
	"errors"
	"net/http"
	"rentals/domain"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type PgHandle struct {
	DB *gorm.DB
}

func (pg *PgHandle) GetRental(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rental, err := domain.FindRentalById(*pg.DB, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, rental)
}

func (pg *PgHandle) GetRentals(w http.ResponseWriter, r *http.Request) {
	filters, err := domain.MakeRentalFilter(r.URL.Query())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rentals, err := domain.FindRentalsWithFilter(*pg.DB, *filters)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, rentals)
}
