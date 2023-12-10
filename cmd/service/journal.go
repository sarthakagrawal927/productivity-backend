package service

import (
	db "todo/pkg/database"
	middleware "todo/pkg/middlewares"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func AddJournalEntry(c echo.Context) error {
	journal := c.Get("journal").(models.JournalEntry)
	queryResult := db.DB_CONNECTION.GetDB().Create(&journal)
	return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Created Successfully", Data: journal}, false)
}

func GetJournalEntries(c echo.Context) error {
	var journalEntries []models.JournalEntry

	pagenum := c.Get("pagenum").(int)
	pagesize := c.Get("pagesize").(int)

	queryResult := db.DB_CONNECTION.GetDB().Select("id", "title", "created_at").Limit(pagesize).Offset((pagenum - 1) * pagesize).Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: true}).Find(&journalEntries)
	return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Success", Data: journalEntries}, true)
}

func GetJournalEntry(c echo.Context) error {
	var journalEntry models.JournalEntry
	id := c.Param("id")
	queryResult := db.DB_CONNECTION.GetDB().Where("id = ?", id).First(&journalEntry)
	return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Success", Data: journalEntry}, false)
}
