package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Habit struct {
	UserId uint `json:"user_id"`
	Meta

	Anti             bool `json:"anti"`
	FrequencyType    uint `json:"frequency_type"`     // 1 - Daily, 2 - Weekly
	Target           uint `json:"target"`             // Limit in case of anti
	Mode             uint `json:"mode"`               // Times, Minutes etc.
	ApproxTimeNeeded uint `json:"approx_time_needed"` // time taken in one instance, needed for count mode to make schedule
	Status           uint `json:"status"`             // 0 - Paused, 1 - Active
	gorm.Model
}

// will just have daily logs, can use group by for weekly goals
// or can cache or rewrite in table with different type
type HabitLog struct {
	HabitID     uint           `json:"habit_id"`
	ResultCount uint           `json:"result_count"`
	ResultDate  datatypes.Date `json:"result_time"`
}

// TV Series, Movies, Books, etc. With this you can choose how much time you want to spend on your habit and it will let you know what you can watch
type Consumable struct {
	HabitID uint `json:"habit_id"`
	Meta

	SmallestUnitLabel string `json:"smallest_unit_label"` // Episode / Page
	NumTotalUnit      uint   `json:"num_total_unit"`      // 12 episodes / 100 pages
	TimePerUnit       uint   `json:"time_per_unit"`       // 20min / 5min
	NumRemainingUnit  uint   `json:"num_remaining_unit"`
	gorm.Model
}

// can do something similar to plan out the exercises
