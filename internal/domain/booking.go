package domain

import "time"

type Booking struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	CarID       uint      `json:"car_id"`
	BookingTime time.Time `json:"booking_time"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
}
