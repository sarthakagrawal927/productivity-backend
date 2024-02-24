package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Meta struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Task struct {
	Meta
	UserId      uint            `json:"user_id"`
	Status      uint            `json:"status"`
	Priority    uint            `json:"priority"`
	TimeToSpend uint            `json:"time_to_spend"`
	Deadline    *datatypes.Date `json:"deadline"`
	// SourceEntity uint          `json:"source"` // can be project or regular task
	// SourceId     uint          `json:"source_id"`
	// TagIds pq.Int64Array `gorm:"type:integer[]" json:"tag_ids"`
	gorm.Model
}

// For future
type Tag struct { // can be used for both tasks and journal entries
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	gorm.Model
}

// The goal is your vision for better and brighter future. The projects are your best plans how to achieve your goal. A goal is a collection of projects.
type Project struct { // collection of tasks
	UserId uint `json:"user_id"`
	GoalID uint `json:"goal_id"`
	Meta
	gorm.Model
}

type Goal struct {
	UserId uint `json:"user_id"`
	Meta

	Why string `json:"why"`
	gorm.Model
}
