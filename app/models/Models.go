package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID       uint `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt gorm.DeletedAt `gorm:"index"`
}
