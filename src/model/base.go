package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"coloum:created_at"`
	UpdatedAt time.Time `gorm:"coloum:updated_at"`
	DeletedAt gorm.DeletedAt
}
