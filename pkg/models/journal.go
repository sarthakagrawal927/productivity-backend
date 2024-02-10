package models

import "gorm.io/gorm"

type JournalEntry struct {
	Meta
	UserId uint `json:"user_id"`

	Type uint `json:"type"` // idea, journal, affirmation etc
	gorm.Model
}

type JournalPrompt struct {
	UserId    uint   `json:"user_id"`
	Question  string `json:"question"`
	PopupTime string `json:"popup_time"`
	gorm.Model
}

// Sample: What 3 things you are proud of? etc,
