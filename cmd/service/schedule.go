package service

import (
	"strconv"
	"sync"
	"todo/pkg/constants"
	db "todo/pkg/database"
	"todo/pkg/models"
)

type HourMinute struct {
	Hour   int
	Minute int
}

type ScheduleEntry struct {
	Label     string     `json:"label"`
	StartTime HourMinute `json:"start_time"`
	EndTime   HourMinute `json:"end_time"`
}

var sleepSchedule = ScheduleEntry{
	Label: "sleep",
	StartTime: HourMinute{
		Hour:   2,
		Minute: 0,
	},
	EndTime: HourMinute{
		Hour:   9,
		Minute: 0,
	},
}

// 12-19:30
var officeSchedule = ScheduleEntry{
	Label: "Work",
	StartTime: HourMinute{
		Hour:   12,
		Minute: 00,
	},
	EndTime: HourMinute{
		Hour:   19,
		Minute: 30,
	},
}

func getScheduleLabel(entry ScheduleEntry) string {
	return strconv.Itoa(entry.StartTime.Hour) + ":" + strconv.Itoa(entry.StartTime.Minute) + " - " + strconv.Itoa(entry.EndTime.Hour) + ":" + strconv.Itoa(entry.EndTime.Minute) + " " + entry.Label
}

func createSchedule() []ScheduleEntry {
	var (
		Habits []models.Habit
		Tasks  []models.Task
	)

	dbInstance := db.DB_CONNECTION.GetDB()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		dbInstance.Where("anti = ?", false).Where("status = ?", constants.HabitActive).Find(&Habits)
	}()

	go func() {
		defer wg.Done()
		dbInstance.Order("deadline DESC").Find(&Tasks)
	}()

	wg.Wait()
	return []ScheduleEntry{
		sleepSchedule,
		officeSchedule,
	}
}

func getFormattedSchedule() ([]string, []ScheduleEntry) {
	schedule := createSchedule()
	var formattedSchedule []string
	for _, entry := range schedule {
		formattedSchedule = append(formattedSchedule, getScheduleLabel(entry))
	}
	return formattedSchedule, schedule
}
