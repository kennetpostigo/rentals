package domain

import (
	"gorm.io/gorm"
)

func FindRentalById(db gorm.DB, id string) (Rental, error) {
	var rental Rental

	err := db.Preload("User").First(&rental, "id = ?", id).Error
	return rental, err
}
