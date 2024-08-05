package service

import (
	"sort"
	"sync"
	"todo/pkg/constants"
	db "todo/pkg/database"
	"todo/pkg/models"
	types "todo/pkg/types"
	"todo/pkg/utils"
)

func createSchedule(userId uint) ([]types.ScheduleEntry, []types.TaskEntry) {
	var (
		Habits []models.Habit
		Tasks  []models.Task
	)

	dbInstance := db.DB_CONNECTION.GetDB()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		dbInstance.Where("anti = ?", false).
			Where("COALESCE(existing_usage, 0) < target").
			Where("status = ? AND user_id = ?", constants.HabitActive, userId).Find(&Habits)
	}()

	go func() {
		defer wg.Done()
		dbInstance.Where("status != ?", constants.Done).Where("user_id = ?", userId).
			Order("deadline DESC").Find(&Tasks)
	}()
	wg.Wait()

	busySlots := []types.ScheduleEntry{constants.SleepSchedule}
	if !utils.IsWeekendToday() {
		busySlots = append(busySlots, constants.OfficeSchedule)
	}

	timeGaps := getTimeGapsFromBusySchedule(busySlots)
	taskEntries := getTaskEntriesFromHabits(Habits)
	taskEntries = append(taskEntries, getTaskEntriesFromTasks(Tasks)...)
	sort.Sort(ByPriority(taskEntries))

	filledSchedule := fillTaskEntriesToAvailableGaps(taskEntries, timeGaps)
	finalSchedule := append(busySlots, filledSchedule...)
	sort.Sort(ByStartTime(finalSchedule))

	return finalSchedule, taskEntries
}

func getFormattedSchedule(userId uint) ([]string, []types.ScheduleEntry, []types.TaskEntry) {
	schedule, taskEntries := createSchedule(userId)
	var formattedSchedule []string
	for _, entry := range schedule {
		formattedSchedule = append(formattedSchedule, getScheduleLabel(entry))
	}
	return formattedSchedule, schedule, taskEntries
}
