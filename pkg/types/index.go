package types

type HourMinute struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type ScheduleEntry struct {
	Label     string     `json:"label"`
	StartTime HourMinute `json:"start_time"`
	EndTime   HourMinute `json:"end_time"`
	Type      string     `json:"type"`
}

type TaskEntry struct {
	EntityType    uint          `json:"entity_type"`
	EntityId      uint          `json:"entity_id"`
	EntityLabel   string        `json:"entity_label"`
	TimeNeeded    uint          `json:"time_needed"`
	Priority      float64       `json:"priority"` // less number is more
	ScheduleEntry ScheduleEntry `json:"schedule_entry"`
}
