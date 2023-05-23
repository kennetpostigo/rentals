package domain

type User struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"first_name" json:"first_name"`
	LastName  string `gorm:"last_name" json:"last_name"`
}

type Price struct {
	Day int `gorm:"column:price_per_day" json:"day"`
}

type Location struct {
	City    string  `gorm:"column:home_city" json:"city"`
	State   string  `gorm:"column:home_state" json:"state"`
	Zip     string  `gorm:"column:home_zip" json:"zip"`
	Country string  `gorm:"column:home_country" json:"country"`
	Lat     float64 `gorm:"column:lat" json:"lat"`
	Lng     float64 `gorm:"column:lng" json:"lng"`
}

type Rental struct {
	Id              int      `json:"id"`
	UserId          int      `json:"-"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Type            string   `json:"type"`
	Make            string   `gorm:"column:vehicle_make" json:"make"`
	Model           string   `gorm:"column:vehicle_model" json:"model"`
	Year            int      `gorm:"column:vehicle_year" json:"year"`
	Length          float64  `gorm:"column:vehicle_length" json:"length"`
	Sleeps          int      `json:"sleeps"`
	PrimaryImageUrl string   `json:"primary_image_url"`
	Price           Price    `gorm:"embedded" json:"price"`
	Location        Location `gorm:"embedded" json:"location"`
	User            User     `json:"user"`
}
