package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Code      string `gorm:"uniqueIndex"`
	Name      string
	IsPrivate bool
}