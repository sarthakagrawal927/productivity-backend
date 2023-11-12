package db

import (
	"fmt"
	"time"
	"todo/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=1234567890 dbname=moves_dev port=5432 sslmode=disable TimeZone=Asia/Shanghai"

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

func Migrations() error {
	fmt.Println("[DB] Migrating models")
	err := DB_CONNECTION.GetDB().AutoMigrate(&models.Task{})
	if err != nil {
		fmt.Println("Error migrating models")
	}
	fmt.Println("[DB] Migrated models")
	return err
}

var DB_CONNECTION *postgresStruct

func SetupDBConnection() {
	var err error
	fmt.Println("[DB] Setting up database connection")
	DB_CONNECTION, err = getNewPostgresDB(dsn)
	if err != nil {
		fmt.Println("[DB] Error setting up database connection")
	} else {
		fmt.Println("[DB] Successfully set up database connection")
	}
}

func (p *postgresStruct) GetDB() *gorm.DB {
	return p.db
}
