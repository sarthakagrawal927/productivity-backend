package service

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
	"todo/pkg/constants"
	db "todo/pkg/database"
	"todo/pkg/models"

	"gorm.io/datatypes"
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

type TaskEntry struct {
	EntityType  uint    `json:"entity_type"`
	EntityId    uint    `json:"entity_id"`
	EntityLabel string  `json:"entity_label"`
	TimeNeeded  uint    `json:"time_needed"`
	Priority    float64 `json:"priority"` // less number is more
}

func calculateDifferenceInDays(taskDeadline *datatypes.Date) (float64, error) {
	if taskDeadline == nil {
		return 8, nil
	}

	deadline, err := taskDeadline.Value()
	if err != nil {
		return 0, err
	}

	deadlineTime := deadline.(time.Time)
	deadlineTime = time.Date(deadlineTime.Year(), deadlineTime.Month(), deadlineTime.Day(), 0, 0, 0, 0, time.UTC)
	hrs := time.Until(deadlineTime).Hours()
	// if it runs in the morning, its fine otherwise gives negative as it considers time difference from 00:00
	return math.Max(math.Round((hrs)/24), 0), nil // hack until I figure out the time zone issues & add cron to ensure based on timezones
}

func createSchedule() ([]ScheduleEntry, []TaskEntry) {
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

	busySlots := []ScheduleEntry{sleepSchedule, officeSchedule}
	taskEntries := make([]TaskEntry, len(Habits)+len(Tasks))

	for i, habit := range Habits {
		percentageFulfilled := float64(habit.ExistingUsage) / float64(habit.Target)
		priority := float64(habit.FrequencyType) * percentageFulfilled // habitFreq 1 is daily, 2 is weekly, 3 is monthly

		timeNeeded := habit.Target - habit.ExistingUsage

		if habit.Mode == constants.HabitCountMode {
			timeNeeded = habit.ApproxTimeNeeded * timeNeeded
		}

		if habit.FrequencyType != constants.HabitDailyFreq { // a good aim is to clear any habit tasks in around 5 times itself
			timeNeeded = timeNeeded / 5
			priority = priority + 0.3 // so that urgent tasks & today tasks are given priority
		}

		taskEntries[i] = TaskEntry{
			EntityType:  constants.ENTITY_HABIT,
			EntityId:    habit.ID,
			EntityLabel: habit.Title,
			TimeNeeded:  timeNeeded, // consider habit type, mode etc
			Priority:    priority,   // 0 - 3.3
		}
	}
	for i, task := range Tasks {
		difference, err := calculateDifferenceInDays(task.Deadline)
		if err != nil {
			fmt.Println("error calculating difference", err)
			continue
		}
		taskEntries[i+len(Habits)] = TaskEntry{
			EntityType:  constants.ENTITY_TASK,
			EntityId:    task.ID,
			EntityLabel: task.Title,
			TimeNeeded:  task.TimeToSpend,
			// considering I allow to set deadline just 7 days in future, lowest priority would be 4-1 + 0.3*7 = 5.1, highest: 4-4 + 0.5*0 = 0
			Priority: float64(constants.HighPriority-task.Priority) + 0.3*difference, // 0 - 5.1
		}
	}

	fmt.Println(busySlots, len(taskEntries))

	return []ScheduleEntry{
		sleepSchedule,
		officeSchedule,
	}, taskEntries
}

func getFormattedSchedule() ([]string, []ScheduleEntry, []TaskEntry) {
	schedule, taskEntries := createSchedule()
	var formattedSchedule []string
	for _, entry := range schedule {
		formattedSchedule = append(formattedSchedule, getScheduleLabel(entry))
	}
	return formattedSchedule, schedule, taskEntries
}
