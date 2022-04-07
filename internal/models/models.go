package models

import (
	"time"

	"gorm.io/gorm"
)

type Timestamps struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type TodoItem struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Timestamps
}
