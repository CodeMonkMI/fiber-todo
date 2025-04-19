package todo

import (
	"time"

	"gorm.io/gorm"
)

type TodoModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string         `json:"title"`
	Completed bool           `json:"completed"`
	CreatedBy uint           `json:"created_by"`
}
type TodoCreateRequest struct {
	Title     string `validate:"required,min=4"`
	Completed bool
}

type TodoUpdateRequest struct {
	Title     string `validate:"required,min=4"`
	Completed bool
}
