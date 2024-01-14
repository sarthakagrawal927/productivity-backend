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
	UserId     uint          `json:"user_id"`
	Status     uint          `json:"status"`
	DueDate    string        `json:"due_date"`
	Priority   uint          `json:"priority"`
	Complexity uint          `json:"complexity"`
	Source     uint          `json:"source"` // can be habit, project or regular task
	SourceId   uint          `json:"source_id"`
	TagIds     pq.Int64Array `gorm:"type:integer[]" json:"tag_ids"`
}

// type Schedule struct {
// 	gorm.Model
// 	StartTime string `json:"start_time"`
// 	EndTime   string `json:"end_time"`
// }

type Tag struct { // can be used for both tasks and journal entries
	gorm.Model
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
}

// The goal is your vision for better and brighter future. The projects are your best plans how to achieve your goal. A goal is a collection of projects.
type Project struct { // collection of tasks
	gorm.Model
	UserId uint `json:"user_id"`
	GoalID uint `json:"goal_id"`
	Meta
}

type Goal struct {
	gorm.Model
	UserId uint `json:"user_id"`
	Meta

	Why string `json:"why"`
}
