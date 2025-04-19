package auth

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	FullName string `json:"fullname"`
}
type RegisterResponse struct {
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	FullName string `json:"fullname"`
}

type RegisterRequest struct {
	Username string `json:"username" gorm:"unique" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	FullName string `json:"fullname" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Password string `json:"password" validate:"required"`
}
