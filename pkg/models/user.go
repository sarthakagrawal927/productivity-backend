package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Will add models related to user here, in future.
type User struct {
	gorm.Model
	Email      string        `json:"email"`
	FullName   string        `json:"full_name"`
	TodayTasks pq.Int64Array `gorm:"type:integer[]" json:"today_tasks"`
}
