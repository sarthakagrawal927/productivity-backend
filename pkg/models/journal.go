package models

import "gorm.io/gorm"

type JournalEntry struct {
	gorm.Model
	Meta

	Type int `json:"type"` // idea, journal, affirmation etc
}

type JournalPrompt struct {
	gorm.Model
	Question  string `json:"question"`
	PopupTime string `json:"popup_time"`
}

// Sample: What 3 things you are proud of? etc,
