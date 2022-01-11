package entity

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id      uint `gorm:"primaryKey"`
	UserID  string
	Content string
}
