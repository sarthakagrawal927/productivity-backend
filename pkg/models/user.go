package models

import "gorm.io/gorm"

// Will add models related to user here, in future.
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
