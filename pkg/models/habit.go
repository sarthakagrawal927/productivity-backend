package models

import "gorm.io/gorm"

type Habit struct {
	gorm.Model
	Meta

	Status    int `json:"status"`
	Frequency int `json:"frequency"`
}

type HabitLog struct {
	gorm.Model
	HabitID uint `json:"habit_id"`
	Amount  int  `json:"amount"`
}
