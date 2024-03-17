package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getNewPostgresDB(url string) (*postgresStruct, error) {
	pg := &postgresStruct{
		connAttempts: 10,
		connTimeout:  5 * time.Second,
	}

	var err error

	for pg.connAttempts > 0 {
		pg.db, err = gorm.Open(postgres.Open(url), &gorm.Config{
			Logger:                 queryLogger,
			SkipDefaultTransaction: true,
		})
		if err == nil {
			break
		}

		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	return pg, err
}

var DB_CONNECTION *postgresStruct

func SetupDBConnection() {
	var err error

	fmt.Println("[DB] Setting up database connection to URL: ", os.Getenv("DATABASE_URL"))
	DB_CONNECTION, err = getNewPostgresDB("host=localhost user=postgres password=1234567890 dbname=moves_dev port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		fmt.Println("[DB] Error setting up database connection")
	} else {
		fmt.Println("[DB] Successfully set up database connection")
	}
}

func (p *postgresStruct) GetDB() *gorm.DB {
	return p.db
}
