package service

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
	"todo/pkg/constants"
	"todo/pkg/models"
	"todo/pkg/types"
	utils "todo/pkg/utils"

	"gorm.io/datatypes"
)

func getMinutesFromHourMinute(hourMinute types.HourMinute) int {
	return hourMinute.Hour*60 + hourMinute.Minute
}

// custom sort functions
type ByPriority []types.TaskEntry

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

type ByStartTime []types.ScheduleEntry

func (a ByStartTime) Len() int      { return len(a) }
func (a ByStartTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByStartTime) Less(i, j int) bool {
	return getMinutesFromHourMinute(a[i].StartTime) < getMinutesFromHourMinute(a[j].StartTime)
}

func getScheduleLabel(entry types.ScheduleEntry) string {
	return strconv.Itoa(entry.StartTime.Hour) + ":" + strconv.Itoa(entry.StartTime.Minute) + " - " + strconv.Itoa(entry.EndTime.Hour) + ":" + strconv.Itoa(entry.EndTime.Minute) + " " + entry.Label
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

func getTaskEntriesFromHabits(Habits []models.Habit) []types.TaskEntry {
	taskEntries := make([]types.TaskEntry, len(Habits))
	for i, habit := range Habits {
		percentageFulfilled := float64(habit.ExistingUsage) / float64(habit.Target)
		priority := float64(habit.FrequencyType) * percentageFulfilled // habitFreq 1 is daily, 2 is weekly, 3 is monthly

		timeNeeded := habit.Target - habit.ExistingUsage

		if habit.Mode == constants.HabitCountMode {
			timeNeeded = habit.ApproxTimeNeeded * timeNeeded
		}

		if habit.FrequencyType != constants.HabitDailyFreq { // a good aim is to clear any habit tasks in around 5 times itself
			if !utils.IsWeekendToday() {
				timeNeeded = timeNeeded / 5
			}
			priority = priority + 0.2 // so that urgent tasks & today tasks are given priority
		}

		taskEntries[i] = types.TaskEntry{
			EntityType:  constants.ENTITY_HABIT,
			EntityId:    habit.ID,
			EntityLabel: "(Habit) " + habit.Title + " - " + habit.Desc,
			TimeNeeded:  timeNeeded, // consider habit type, mode etc
			Priority:    priority,   // 0 - 3.3
		}

		if habit.FrequencyType == constants.HabitDailyFreq && habit.PreferredTimePeriod != "" {
			taskEntries[i].Priority -= 0.3 // so that preferred tasks are given priority
			taskEntries[i].ScheduleEntry, _ = utils.ConvertToScheduleEntry(habit.PreferredTimePeriod)
		}
	}
	return taskEntries
}

func getTaskEntriesFromTasks(Tasks []models.Task) []types.TaskEntry {
	taskEntries := make([]types.TaskEntry, len(Tasks))
	for i, task := range Tasks {
		difference, err := calculateDifferenceInDays(task.Deadline)
		if err != nil {
			fmt.Println("error calculating difference", err)
			continue
		}
		taskEntries[i] = types.TaskEntry{
			EntityType:  constants.ENTITY_TASK,
			EntityId:    task.ID,
			EntityLabel: "(Task) " + task.Title + " - " + task.Desc,
			TimeNeeded:  task.TimeToSpend,
			Priority:    float64(constants.HighPriority-task.Priority) + 0.2*difference,
		}
	}
	return taskEntries
}

// can improve this but works for now.
func getTimeGapsFromBusySchedule(busy []types.ScheduleEntry) []types.ScheduleEntry {
	sort.Sort(ByStartTime(busy))

	var timeGaps []types.ScheduleEntry

	// Function to add a gap entry
	addGapEntry := func(start types.HourMinute, end types.HourMinute) {
		timeGaps = append(timeGaps, types.ScheduleEntry{
			Label:     "Free",
			StartTime: start,
			EndTime:   end,
			Type:      "gap",
		})
	}

	// Add gap before first busy
	if busy[0].StartTime.Hour != 0 || busy[0].StartTime.Minute != 0 {
		addGapEntry(types.HourMinute{Hour: 0, Minute: 0}, busy[0].StartTime)
	}

	// Add gaps between busy slots
	for i := 0; i < len(busy)-1; i++ {
		addGapEntry(busy[i].EndTime, busy[i+1].StartTime)
	}

	// Add gap after last busy
	if busy[len(busy)-1].EndTime.Hour != 24 || busy[len(busy)-1].EndTime.Minute != 0 {
		addGapEntry(busy[len(busy)-1].EndTime, types.HourMinute{Hour: 24, Minute: 0})
	}

	return timeGaps
}

func addTimeToHourMinute(hourMinute types.HourMinute, timeToAdd uint) types.HourMinute {
	hourMinute.Minute += int(timeToAdd % 60)
	hourMinute.Hour += int(timeToAdd / 60)
	if hourMinute.Minute >= 60 {
		hourMinute.Minute -= 60
		hourMinute.Hour++
	}
	return hourMinute
}

func fillTaskEntriesToAvailableGaps(taskEntries []types.TaskEntry, gaps []types.ScheduleEntry) []types.ScheduleEntry {
	// assuming gaps are sorted by startTime & taskEntries are sorted by priority
	// also want to have 5min buffer before & after each task
	// since this is a difficult dynamic programming problem, I will use a greedy approach for now, fill max priority task in smallest gap it can fit in
	// in future, can suggest different schedules based on different assumptions & tradeoffs, can also consider splitting the task: either here or in before process.
	// can consider user's energy, quadrants, recalibrate based on every schedule movement
	findClosestGap := func(timeNeeded uint) int {
		smallestGapIdx := -1
		smallestGapTime := 24 * 60
		for i, gap := range gaps {
			if gap.Type != "gap" {
				continue
			}
			gapTime := getMinutesFromHourMinute(gap.EndTime) - getMinutesFromHourMinute(gap.StartTime)
			if gapTime >= int(timeNeeded) && gapTime < smallestGapTime {
				smallestGapTime = gapTime
				smallestGapIdx = i
			}
		}
		return smallestGapIdx
	}

	// returns gapIdx if habit has preferred time, -1 otherwise
	getIdxForHabitHasPreferredTime := func(taskScheduleEntry types.ScheduleEntry) int {
		for i, gap := range gaps {
			if gap.Type != "gap" {
				continue
			}
			if getMinutesFromHourMinute(gap.StartTime) <= getMinutesFromHourMinute(taskScheduleEntry.StartTime) &&
				getMinutesFromHourMinute(gap.EndTime) >= getMinutesFromHourMinute(taskScheduleEntry.EndTime) {
				return i
			}
		}
		return -1
	}

	insertTaskInGap := func(task types.TaskEntry, gapIdx int) {
		var scheduledEntry types.ScheduleEntry
		if task.ScheduleEntry != (types.ScheduleEntry{}) {
			scheduledEntry = types.ScheduleEntry{
				StartTime: task.ScheduleEntry.StartTime,
				EndTime:   task.ScheduleEntry.EndTime,
			}
		} else {
			scheduledEntry = types.ScheduleEntry{
				StartTime: types.HourMinute{
					Hour:   gaps[gapIdx].StartTime.Hour,
					Minute: gaps[gapIdx].StartTime.Minute,
				},
				EndTime: addTimeToHourMinute(gaps[gapIdx].StartTime, task.TimeNeeded),
			}
		}
		scheduledEntry.Label = task.EntityLabel
		scheduledEntry.Type = "task"

		startMinuteGap := getMinutesFromHourMinute(gaps[gapIdx].StartTime)
		endMinuteGap := getMinutesFromHourMinute(gaps[gapIdx].EndTime)

		startMinuteTask := getMinutesFromHourMinute(scheduledEntry.StartTime)
		endMinuteTask := getMinutesFromHourMinute(scheduledEntry.EndTime)

		if startMinuteGap == startMinuteTask && endMinuteGap == endMinuteTask { // task fits perfectly in gap
			gaps[gapIdx] = scheduledEntry
		} else if startMinuteGap == startMinuteTask { // task is in the beginning of gap
			gaps = utils.InsertElementsInSliceAfterIdx(gaps, []types.ScheduleEntry{scheduledEntry}, gapIdx-1)
			gaps[gapIdx+1].StartTime = scheduledEntry.EndTime
		} else if endMinuteGap == endMinuteTask { // task is in the end of gap
			gaps = utils.InsertElementsInSliceAfterIdx(gaps, []types.ScheduleEntry{scheduledEntry}, gapIdx)
			gaps[gapIdx].EndTime = scheduledEntry.StartTime
		} else { // task is in the middle of gap
			gaps = utils.InsertElementsInSliceAfterIdx(gaps, []types.ScheduleEntry{scheduledEntry, {
				StartTime: scheduledEntry.EndTime,
				EndTime:   gaps[gapIdx].EndTime,
				Label:     "Free",
				Type:      "gap",
			}}, gapIdx)
			gaps[gapIdx].EndTime = scheduledEntry.StartTime
		}
	}

	var gapIdx int
	for _, task := range taskEntries {
		gapIdx = -1
		if task.EntityType == constants.ENTITY_HABIT && task.ScheduleEntry != (types.ScheduleEntry{}) {
			gapIdx = getIdxForHabitHasPreferredTime(task.ScheduleEntry)
		}
		if gapIdx == -1 {
			gapIdx = findClosestGap(task.TimeNeeded)
			task.ScheduleEntry = types.ScheduleEntry{}
		}
		if gapIdx == -1 {
			continue
		}
		insertTaskInGap(task, gapIdx)
	}

	// filter out the gaps that are not used
	var filledSchedule []types.ScheduleEntry
	for idx, scheduleEntry := range gaps {
		if scheduleEntry.Type != "gap" {
			filledSchedule = append(filledSchedule, gaps[idx])
		}
	}

	return filledSchedule
}
