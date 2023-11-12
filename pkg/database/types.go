package db

import (
	"time"

	"gorm.io/gorm"
)

type postgresStruct struct {
	connAttempts uint
	connTimeout  time.Duration

	db *gorm.DB
}
