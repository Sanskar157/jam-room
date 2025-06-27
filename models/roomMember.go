package models

import (
	"gorm.io/gorm"
)

type RoomMember struct {
	gorm.Model
	RoomID   uint   `gorm:"index"`
	Alias    string
	IsAdmin  bool   `gorm:"default:false"`  
}