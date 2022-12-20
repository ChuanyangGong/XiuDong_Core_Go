package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"coloum:created_at"`
	UpdatedAt time.Time `gorm:"coloum:updated_at"`
	DeletedAt gorm.DeletedAt
}
