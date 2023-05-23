package domain

import (
	"errors"
	"math"
	"net/url"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type RentalFilter struct {
	PriceMax *int
	PriceMin *int
	Limit    int
	Offset   int
	IDs      []int
	Near     []float64
	Sort     string
}

const MILE = 1609.344

var SORT_DB_MAP = map[string]string{
	"id":      "id",
	"name":    "name",
	"type":    "type",
	"make":    "vehicle_make",
	"model":   "vehicle_model",
	"year":    "vehicle_year",
	"length":  "vehicle_length",
	"sleeps":  "sleeps",
	"price":   "price_per_day",
	"city":    "home_city",
	"state":   "home_state",
	"zip":     "home_zip",
	"country": "home_country",
}

func stringToIds(str string) ([]int, error) {
	strIDs := strings.Split(str, ",")
	ids := make([]int, len(strIDs))

	for i, s := range strIDs {
		v, err := strconv.Atoi(s)

		if err != nil {
			return nil, err
		}

		ids[i] = v
	}

	return ids, nil
}

func stringToCoordintates(str string) ([]float64, error) {
	strCoords := strings.Split(str, ",")
	coords := make([]float64, 2)

	if len(strCoords) == 2 {
		lat, err := strconv.ParseFloat(strCoords[0], 64)
		if err != nil {
			return nil, err
		} else {
			coords[0] = lat
		}

		lng, err := strconv.ParseFloat(strCoords[1], 64)
		if err != nil {
			return nil, err
		} else {
			coords[1] = lng
		}

		return coords, nil
	}

	return nil, errors.New("invalid Lat,Lng coordinates")
}

func MakeRentalFilter(qp url.Values) (*RentalFilter, error) {
	filters := RentalFilter{
		PriceMin: nil,
		PriceMax: nil,
		Limit:    10,
		Offset:   0,
		IDs:      make([]int, 0),
		Near:     make([]float64, 0),
		Sort:     "id",
	}

	if qp.Get("price_min") != "" {
		priceMin, err := strconv.Atoi(qp.Get("price_min"))
		if err != nil {
			return nil, err
		} else {
			filters.PriceMin = &priceMin
		}
	}

	if qp.Get("price_max") != "" {
		priceMax, err := strconv.Atoi(qp.Get("price_max"))
		if err != nil {
			return nil, err

		} else {
			filters.PriceMax = &priceMax
		}
	}

	if qp.Get("limit") != "" {
		limit, err := strconv.Atoi(qp.Get("limit"))
		if limit > 0 {
			if err != nil {
				return nil, err
			} else {
				filters.Limit = limit
			}
		}
	}

	if qp.Get("offset") != "" {
		offset, err := strconv.Atoi(qp.Get("offset"))
		offset = offset * filters.Limit
		if offset > 0 {
			if err != nil {
				return nil, err
			} else {
				filters.Offset = offset
			}
		}
	}

	if qp.Get("ids") != "" {
		ids, err := stringToIds(qp.Get("ids"))
		if err != nil {
			return nil, err
		} else {
			filters.IDs = ids
		}
	}

	if qp.Get("near") != "" {
		near, err := stringToCoordintates(qp.Get("near"))
		if err != nil {
			return nil, err
		} else {
			filters.Near = near
		}
	}

	if qp.Get("sort") != "" {
		sort := qp.Get("sort")
		sort, exists := SORT_DB_MAP[sort]

		if sort != "" && exists {
			filters.Sort = sort
		} else {
			return nil, errors.New("invalid sort key, allowed sort keys: name, type, make, model, year, length, sleeps, price, city, state, zip, country")
		}
	}

	return &filters, nil
}

func FindRentalById(db gorm.DB, id string) (Rental, error) {
	var rental Rental

	err := db.Preload("User").First(&rental, "id = ?", id).Error
	return rental, err
}

func FindRentalsWithFilter(db gorm.DB, filters RentalFilter) (Pagination, error) {
	var rentals []Rental

	query := db.Preload("User")

	if filters.PriceMax != nil {
		query = query.Where("price_per_day < ?", &filters.PriceMax)
	}

	if filters.PriceMin != nil {
		query = query.Where("price_per_day > ?", &filters.PriceMin)
	}

	query = query.Limit(filters.Limit)

	if filters.Offset != 0 {
		query = query.Offset(filters.Offset)
	}

	if len(filters.IDs) > 0 {
		query = query.Where("id IN (?)", filters.IDs)
	}

	if len(filters.Near) == 2 {
		// Note: I'm not crazy familiar with PostGis but I found this on StackOverflow: https://stackoverflow.com/questions/51889155/getting-all-buildings-in-range-of-5-miles-from-specified-coordinates/51889638
		query = query.Where("ST_Distance(ST_MakePoint(lng, lat), ST_MakePoint(?, ?)) <= ?", &filters.Near[1], &filters.Near[0], MILE*100)
	}

	query = query.Order(filters.Sort)

	err := query.Find(&rentals).Error

	var totalRows int64
	db.Model(rentals).Count(&totalRows)

	page := (filters.Offset / filters.Limit)
	if page == 0 {
		page = 1
	}

	res := Pagination{
		Limit:      filters.Limit,
		Offset:     filters.Offset,
		Page:       page,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(filters.Limit))),
		Rows:       rentals,
	}

	return res, err
}
