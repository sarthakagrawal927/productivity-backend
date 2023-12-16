package validators

import (
	"todo/pkg/constants"
	middleware "todo/pkg/middlewares"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

func CreateJournalValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		journal := models.JournalEntry{}
		var err error

		if journal.Title, err = validateStringFromForm(c, "title"); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if journal.Desc, err = validateStringFromForm(c, "desc"); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if journal.Type, err = validateIntFromArray(c.FormValue("type"), constants.JournalTypeList); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		c.Set("journal", journal)
		return next(c)
	}
}

func GetJournalValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		pagenum, err := validateInt("pagenum", c.QueryParam("pagenum"), 1)
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}

		pagesize, err := validateInt("pagesize", c.QueryParam("pagesize"), constants.DefaultPageSize)
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}

		journalType, err := validateIntFromArray(c.QueryParam("type"), constants.JournalTypeList, 0)
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}

		c.Set("pagenum", int(pagenum))
		c.Set("pagesize", int(pagesize))
		c.Set("type", int(journalType))
		return next(c)
	}
}

func GetJournalEntryValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := validateAndGetId(c.Param("id"))
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}

		c.Set("id", int(id))
		return next(c)
	}
}
