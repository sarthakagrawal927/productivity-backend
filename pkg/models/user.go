package models

import (
	"time"

	"gorm.io/gorm"
)

type TimeType time.Time

type TimePeriod struct {
	StartTime TimeType `json:"start_time,omitempty"`
	EndTime   TimeType `json:"end_time,omitempty"`
}

type Activity struct {
	Name            string     `json:"name"`
	TimePeriod      TimePeriod `json:"period"`
	RelatedEntity   int64      `json:"related_entity"`
	RelatedEntityId int64      `json:"related_entity_id"`
}

// Thinking of storing this in mongoDB.
type User struct {
	gorm.Model
	Email          string     `json:"email"`
	FullName       string     `json:"full_name"`
	SleepSchedule  TimePeriod `json:"sleep_schedule"`
	OfficeSchedule TimePeriod `json:"office_schedule"`
	Activities     []Activity `json:"activities"`
}
