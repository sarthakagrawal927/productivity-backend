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

		if journal.Title, err = ValidateStringFromForm(c, "title"); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if journal.Desc, err = ValidateStringFromForm(c, "desc"); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		c.Set("journal", journal)
		return next(c)
	}
}

func GetJournalValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		pagenum, err := ValidateInt("pagenum", c.QueryParam("pagenum"), 1)
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}

		pagesize, err := ValidateInt("pagesize", c.QueryParam("pagesize"), constants.DefaultPageSize)
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}

		journalType, err := ValidateInt("type", c.QueryParam("type"), 0)
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
