package service

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
	"todo/pkg/constants"
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
	Type      string     `json:"type"`
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

func getMinutesFromHourMinute(hourMinute HourMinute) int {
	return hourMinute.Hour*60 + hourMinute.Minute
}

// custom sort functions
type ByPriority []TaskEntry

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

type ByStartTime []ScheduleEntry

func (a ByStartTime) Len() int      { return len(a) }
func (a ByStartTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByStartTime) Less(i, j int) bool {
	return getMinutesFromHourMinute(a[i].StartTime) < getMinutesFromHourMinute(a[j].StartTime)
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

func getTaskEntriesFromHabits(Habits []models.Habit) []TaskEntry {
	taskEntries := make([]TaskEntry, len(Habits))
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
			EntityLabel: "(Habit) " + habit.Title + " - " + habit.Desc,
			TimeNeeded:  timeNeeded, // consider habit type, mode etc
			Priority:    priority,   // 0 - 3.3
		}
	}
	return taskEntries
}

func getTaskEntriesFromTasks(Tasks []models.Task) []TaskEntry {
	taskEntries := make([]TaskEntry, len(Tasks))
	for i, task := range Tasks {
		difference, err := calculateDifferenceInDays(task.Deadline)
		if err != nil {
			fmt.Println("error calculating difference", err)
			continue
		}
		taskEntries[i] = TaskEntry{
			EntityType:  constants.ENTITY_TASK,
			EntityId:    task.ID,
			EntityLabel: "(Task) " + task.Title + " - " + task.Desc,
			TimeNeeded:  task.TimeToSpend,
			// considering I allow to set deadline just 7 days in future, lowest priority would be 4-1 + 0.3*7 = 5.1, highest: 4-4 + 0.5*0 = 0
			Priority: float64(constants.HighPriority-task.Priority) + 0.3*difference, // 0 - 5.1
		}
	}
	return taskEntries
}

// can improve this but works for now.
func getTimeGapsFromBusySchedule(busy []ScheduleEntry) []ScheduleEntry {
	sort.Sort(ByStartTime(busy))

	var timeGaps []ScheduleEntry

	// Function to add a gap entry
	addGapEntry := func(start HourMinute, end HourMinute) {
		timeGaps = append(timeGaps, ScheduleEntry{
			Label:     "Free",
			StartTime: start,
			EndTime:   end,
			Type:      "gap",
		})
	}

	// Add gap before first busy
	if busy[0].StartTime.Hour != 0 || busy[0].StartTime.Minute != 0 {
		addGapEntry(HourMinute{0, 0}, busy[0].StartTime)
	}

	// Add gaps between busy slots
	for i := 0; i < len(busy)-1; i++ {
		addGapEntry(busy[i].EndTime, busy[i+1].StartTime)
	}

	// Add gap after last busy
	if busy[len(busy)-1].EndTime.Hour != 24 || busy[len(busy)-1].EndTime.Minute != 0 {
		addGapEntry(busy[len(busy)-1].EndTime, HourMinute{24, 0})
	}

	return timeGaps
}

func addTimeToHourMinute(hourMinute HourMinute, timeToAdd uint) HourMinute {
	hourMinute.Minute += int(timeToAdd % 60)
	hourMinute.Hour += int(timeToAdd / 60)
	if hourMinute.Minute >= 60 {
		hourMinute.Minute -= 60
		hourMinute.Hour++
	}
	return hourMinute
}

func fillTaskEntriesToAvailableGaps(taskEntries []TaskEntry, gaps []ScheduleEntry) []ScheduleEntry {
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

	insertTaskInGap := func(task TaskEntry, gapIdx int) {
		scheduledEntry := ScheduleEntry{
			Label: task.EntityLabel,
			StartTime: HourMinute{
				Hour:   gaps[gapIdx].StartTime.Hour,
				Minute: gaps[gapIdx].StartTime.Minute,
			},
			EndTime: addTimeToHourMinute(gaps[gapIdx].StartTime, task.TimeNeeded),
			Type:    "task",
		}
		gaps = append(gaps, ScheduleEntry{})
		copy(gaps[gapIdx+1:], gaps[gapIdx:])
		gaps[gapIdx] = scheduledEntry
		gaps[gapIdx+1].StartTime = addTimeToHourMinute(scheduledEntry.EndTime, 5)
	}

	for _, task := range taskEntries {
		gapIdx := findClosestGap(task.TimeNeeded)
		if gapIdx == -1 {
			// fmt.Println("No gap found for task", task.EntityLabel)
			continue
		}
		insertTaskInGap(task, gapIdx)
	}

	// filter out the gaps that are not used
	var filledSchedule []ScheduleEntry
	for idx, scheduleEntry := range gaps {
		if scheduleEntry.Type != "gap" {
			filledSchedule = append(filledSchedule, gaps[idx])
		}
	}

	return filledSchedule
}
