package service

import (
	db "todo/pkg/database"
	"todo/pkg/models"

	"github.com/labstack/echo"
)

func migrateDB(c echo.Context) error {
	err := db.DB_CONNECTION.GetDB().AutoMigrate(&models.Task{})
	if err != nil {
		return c.JSON(500, "Error migrating models")
	} else {
		return c.JSON(200, "Migrated models Successfully")
	}
}
