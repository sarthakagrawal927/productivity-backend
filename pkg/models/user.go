package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TimeType datatypes.Time

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
	Email              string         `json:"email"`
	FullName           string         `json:"full_name"`
	SleepStartTime     datatypes.Time `json:"sleep_start_time"`
	SleepEndTime       datatypes.Time `json:"sleep_end_time"`
	WorkStartTime      datatypes.Time `json:"work_start_time"`
	WorkEndTime        datatypes.Time `json:"work_end_time"`
	PasswordAntiHabits string         `json:"password_anti_habits"`
	gorm.Model
}
type Relationship struct {
	UserId uint `json:"user_id"`

	// RelatedPersonId  uint           `json:"related_person_id"` in future to replace name & photoUrl
	Name     string `json:"name"`
	PhotoUrl string `json:"photo_url"`

	ElectronBand     uint           `json:"electron_band"`
	ContactFrequency uint           `json:"contact_frequency"`
	LastContactTime  datatypes.Date `json:"last_contact_time"`
}
