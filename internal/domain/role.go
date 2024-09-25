package domain

type Role struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
