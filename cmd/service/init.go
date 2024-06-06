package service

import (
	"net/http"
	"todo/pkg/metrics"
	validators "todo/pkg/validators"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func CreateService() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.Logger.SetLevel(1)
	logger, _ := zap.NewProduction()

	s := metrics.NewStats()
	e.Use(s.Process)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	e.GET("/api/user/today/logs", GetDailyLogs)
	e.GET("/api/user/today/schedule", GetTodaySchedule)

	e.POST("/api/todo", CreateTodo, validators.CreateTaskValidator)
	e.GET("/api/todo", GetTodo, validators.GetTasksValidator)
	e.DELETE("/api/todo", DeleteTodo, validators.DeleteTaskValidator)
	e.PATCH("/api/todo", UpdateTodo, validators.UpdateTaskValidator)

	e.POST("/api/journal", AddJournalEntry, validators.CreateJournalValidator)
	e.GET("/api/journal", GetJournalEntries, validators.GetJournalValidator)
	e.GET("/api/journal/:id", GetJournalEntry, validators.GetJournalEntryValidator)

	e.POST("/api/habit", CreateHabit, validators.CreateHabitValidator)
	e.GET("/api/habit", GetHabits, validators.GetHabitsValidator)
	e.POST("/api/habit/log", AddHabitLog, validators.CreateHabitLogValidator)
	e.GET("/api/habit/logs/:id", GetHabitWithLogs, validators.GetSingleHabitValidator)

	e.POST("/api/consumable/book", CreateBookConsumable, validators.CreateBookValidator)
	e.POST("/api/consumable/food", CreateFoodConsumable, validators.CreateFoodValidator)
	e.POST("/api/consumable/food/log", CreateFoodConsumed, validators.FoodConsumedValidator)
	e.GET("/api/consumable/food", GetFoodItems)
	e.GET("/api/consumable/food/log", GetDailyFoodLogs)
	e.GET("/api/consumable/food/consumption_items", GetFoodConsumed, validators.FoodConsumptionByDateValidator)

	e.POST("/api/admin/db_migrate", migrateDB)
	e.POST("/api/admin/db_delete_all", deleteAllTasks)
	e.GET("/api/admin/metrics", s.Handle)
	e.POST("/api/admin/db_seed", seedTasks)

	e.Logger.Fatal(e.Start(":1323"))
}
