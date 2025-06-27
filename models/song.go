package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	RoomID    uint   `gorm:"index"`
	VideoID   string // YouTube video ID
	Title     string
	Thumbnail string
	AddedBy   string
}