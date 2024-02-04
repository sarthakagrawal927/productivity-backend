package service

import (
	"todo/pkg/constants"
	db "todo/pkg/database"
	"todo/pkg/models"
	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func AddJournalEntry(c echo.Context) error {
	journal := c.Get("journal").(models.JournalEntry)
	queryResult := db.DB_CONNECTION.GetDB().Create(&journal)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Created Successfully", Data: journal}, false)
}

func GetJournalEntries(c echo.Context) error {
	var journalEntries []models.JournalEntry

	pagenum := int(c.Get("pagenum").(uint))
	pagesize := int(c.Get("pagesize").(uint))
	journalType := c.Get("type").(uint)

	journalTypes := constants.JournalTypeList
	if journalType != 0 {
		journalTypes = []uint{journalType}
	}

	queryResult := db.DB_CONNECTION.GetDB().Select("id", "title", "created_at", "desc", "type").Where("type in ?", journalTypes).Limit(pagesize).Offset((pagenum - 1) * pagesize).Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: true}).Find(&journalEntries)

	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: journalEntries}, true)
}

func GetJournalEntry(c echo.Context) error {
	var journalEntry models.JournalEntry
	id := c.Param("id")
	queryResult := db.DB_CONNECTION.GetDB().Where("id = ?", id).First(&journalEntry)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: journalEntry}, false)
}
