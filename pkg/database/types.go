package db

import (
	"time"

	"gorm.io/gorm"
)

type postgresStruct struct {
	connAttempts int
	connTimeout  time.Duration

	db *gorm.DB
}
