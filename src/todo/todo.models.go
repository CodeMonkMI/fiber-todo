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
}
