package models

import "gorm.io/gorm"

type SongVote struct {
	gorm.Model
	SongID    uint      `gorm:"index"`     // Foreign key to Song
	Alias     string                         // RoomMember alias
	Vote      int                            // +1 for upvote, -1 for downvote
}