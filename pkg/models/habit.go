package models

import "gorm.io/gorm"

type Habit struct {
	gorm.Model
	Meta

	Status    uint `json:"status"`
	Frequency uint `json:"frequency"`
}

type HabitLog struct {
	gorm.Model
	HabitID uint `json:"habit_id"`
	Amount  uint `json:"amount"`
}
