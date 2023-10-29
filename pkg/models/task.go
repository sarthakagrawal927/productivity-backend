package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Status   int    `json:"status"`
	DueDate  string `json:"due_date"`
	Priority int    `json:"priority"`
}
