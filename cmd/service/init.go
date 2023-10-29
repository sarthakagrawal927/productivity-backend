package service

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateService() {
	e := echo.New()

	e.Logger.SetLevel(1)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	e.POST("/api/todo", CreateTodo)
	e.GET("/api/todo", GetTodo)
	e.POST("/api/admin/db_migrate", migrateDB)

	e.Logger.Fatal(e.Start(":1323"))
}
