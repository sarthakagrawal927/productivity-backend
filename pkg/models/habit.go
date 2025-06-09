package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Habit struct {
	UserId uint `json:"user_id"`
	Meta

	FrequencyType    uint  `json:"frequency_type"`               // 1 - Daily, 2 - Weekly, 3 - Monthly
	UpperLimit       uint  `json:"upper_limit"`                  // Limit in case of anti
	LowerLimit       uint  `json:"lower_limit"`                  // Limit in case of anti
	Priority         uint  `json:"priority"`                     // Same as task
	Mode             uint  `json:"mode"`                         // Counts, Minutes etc.
	ApproxTimeNeeded *uint `json:"approx_time_needed,omitempty"` // time taken in one instance, needed for count mode to make schedule
	Status           uint  `json:"status"`                       // 0 - Paused, 1 - Active
	Category         uint  `json:"category"`                     // 1 - Health, 2 - Finance, 3 - Relationship, 4 - Brain, 5 - Productivity

	UpgradeStatus uint `json:"upgrade_status"`  // 0 - Not Upgradable, 1 - Upgradable, 2 - Upgraded, 3 - Downgraded, 4 - NotAllowed
	LatestHabitId uint `json:"latest_habit_id"` // to link all prev habits

	Score         uint `json:"score"`
	ExistingUsage uint `json:"existing_usage"` // based on frequency type update this, to handle multiple logs, insert the log then based on freq update

	// ShouldSchedule     bool     `json:"should_schedule"`
	PreferredStartTime    *datatypes.Time `json:"preferred_start_time"`
	PreferredWeekdaysMask uint8           `json:"preferred_weekdays_mask"`
	PreferredMonthDate    *uint           `json:"preferred_month_date"` // 12th

	gorm.Model
}

// will just have daily logs, can use group by for weekly goals
// or can cache or rewrite in table with different type
type HabitLog struct {
	UserId  uint `json:"user_id"`
	HabitID uint `json:"habit_id"`
	Count   uint `json:"count"`

	Comment        *string         `json:"comment,omitempty"`
	HabitStartTime *datatypes.Time `json:"start_time"`
	LoggedForDate  datatypes.Date  `json:"logged_for_date"`
	MoodRating     uint            `json:"mood_rating"`
	// can get end time by just adding count
	gorm.Model
}

// can do something similar to plan out the exercises

// Once I am happy with the metric calculation
// type HabitMetric struct {
// 	HabitID        uint           `json:"habit_id"`
// 	UserId         uint           `json:"user_id"`
// 	MetricDate     datatypes.Date `json:"metric_date"`
// 	CompletionRate float64        `json:"completion_rate"`
// 	ScoreChange    int            `json:"score_change"` // Can be negative
// 	gorm.Model
// }
