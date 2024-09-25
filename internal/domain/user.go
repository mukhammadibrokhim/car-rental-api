package domain

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	RoleID   uint   `json:"role_id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Address  string `json:"address"`
	DOB      string `json:"dob"`
	Password string `json:"password" binding:"required"`
	Role     Role   `gorm:"foreignKey:RoleID"` // One-to-many relation with Role

}
