package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Habit struct {
	UserId uint `json:"user_id"`
	Meta

	Anti             bool `json:"anti"`
	FrequencyType    uint `json:"frequency_type"`     // 1 - Daily, 2 - Weekly, 3 - Monthly
	Target           uint `json:"target"`             // Limit in case of anti
	Mode             uint `json:"mode"`               // Times, Minutes etc.
	ApproxTimeNeeded uint `json:"approx_time_needed"` // time taken in one instance, needed for count mode to make schedule
	Status           uint `json:"status"`             // 0 - Paused, 1 - Active

	ExistingUsage uint `json:"existing_usage"` // based on frequency type update this, to handle multiple logs, insert the log then based on freq update

	PreferredTimePeriod string `json:"preferred_time_period"` // 12:00-14:00
	gorm.Model
}

// will just have daily logs, can use group by for weekly goals
// or can cache or rewrite in table with different type
type HabitLog struct {
	HabitID     uint           `json:"habit_id"`
	ResultCount uint           `json:"result_count"`
	ResultDate  datatypes.Date `json:"result_time"`
	Comment     string         `json:"comment"`

	// Decide whether you want to somewhere like toggl route or just make the schedule for user
	// StartTime   datatypes.Time `json:"start_time"`
	// EndTime     datatypes.Time `json:"end_time"`
	gorm.Model
}

// can do something similar to plan out the exercises
