package domain

type Car struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	CarNumber   string `json:"car_number"`
	CarType     string `json:"car_type"`
	Description string `json:"description"`
}
