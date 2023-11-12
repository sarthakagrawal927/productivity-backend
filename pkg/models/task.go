package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Meta struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Task struct {
	gorm.Model
	Meta
	Status     uint          `json:"status"`
	DueDate    string        `json:"due_date"`
	Priority   uint          `json:"priority"`
	Complexity uint          `json:"complexity"`
	Source     uint          `json:"source"` // can be habit, goal or regular task
	SourceId   uint          `json:"source_id"`
	TagIds     pq.Int64Array `gorm:"type:integer[]"`
}

type Schedule struct {
	gorm.Model
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// TODO Consider merging both goal & project or tags, unsure whether all 3 are needed
type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

type Goal struct {
	gorm.Model
	Meta

	Why string `json:"why"`
}

type Project struct {
	gorm.Model
	Meta
}
