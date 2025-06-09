package constants

const AllStatus uint = 0

const (
	Todo uint = 1 + iota
	// Scheduled
	InProgress
	Done
	Backlog
)

const (
	NoPriority uint = iota + 1
	LowPriority
	MediumPriority
	HighPriority
)

const (
	Idea uint = iota + 1
	Gratitude
	MindClear
	DayPlanning
	DayWrap
	Event
	FoodLog
	HighlightOfTheDay
)

const (
	HabitDailyFreq uint = iota + 1
	HabitWeeklyFreq
	HabitMonthlyFreq
)

const (
	HabitActive uint = iota + 1
	HabitArchived
)

const (
	HabitTimeMode  uint = iota + 1
	HabitCountMode      // for example, 10 pushups
	HabitMlMode         // for example, 1 liter of water
)

const (
	HabitCategoryHealth uint = iota + 1
	HabitCategoryFinance
	HabitCategoryRelationship
	HabitCategoryBrain
	HabitCategoryProductivity
)

const (
	MoodRatingVeryBad  uint = iota + 1 // 1
	MoodRatingBad                      // 2
	MoodRatingNeutral                  // 3
	MoodRatingGood                     // 4
	MoodRatingVeryGood                 // 5
)

const DefaultPageSize = 20

var (
	TaskTypeList     = []uint{uint(Todo), uint(InProgress), uint(Done), uint(Backlog)}
	PriorityTypeList = []uint{uint(NoPriority), uint(LowPriority), uint(MediumPriority), uint(HighPriority)}

	JournalTypeList = []uint{
		AllStatus, uint(Idea), uint(Gratitude),
		uint(MindClear), uint(DayPlanning), uint(DayWrap),
		uint(Event), uint(FoodLog), uint(HighlightOfTheDay),
	}

	HabitFreqTypeList = []uint{uint(HabitDailyFreq), uint(HabitWeeklyFreq), uint(HabitMonthlyFreq)}
	HabitStatusList   = []uint{uint(HabitArchived), uint(HabitActive)}
	HabitModeList     = []uint{uint(HabitTimeMode), uint(HabitCountMode), uint(HabitMlMode)}
	HabitCategoryList = []uint{uint(HabitCategoryHealth), uint(HabitCategoryFinance), uint(HabitCategoryRelationship), uint(HabitCategoryBrain), uint(HabitCategoryProductivity)}
	MoodRatingList    = []uint{uint(MoodRatingVeryBad), uint(MoodRatingBad), uint(MoodRatingNeutral), uint(MoodRatingGood), uint(MoodRatingVeryGood)}
)

const (
	ACTIVITY_TIME_CHANGED uint = iota + 1
	ACTIVITY_SKIPPED_FOR_TIME
)

// for every activity (based on habit) user skips or time changes, we log to activity table. Using this data we can suggest users to modify their habits.
// based on the reason of skipping or changing time, we can suggest users to improvise. Do not store successes, only failures.

const (
	ENTITY_TASK    uint = iota + 1
	ENTITY_HABIT        // 2
	ENTITY_PROJECT      // 3
	ENTITY_GOAL         // 4
)

const (
	BOOK_READING  uint = iota + 1
	BOOK_FINISHED      // 2
	BOOK_TO_READ       // 3
)

const (
	FOOD_LOG_DAY_MODE uint = iota + 1
	FOOD_LOG_WEEK_MODE
)

const (
	REGULAR_FOOD uint = iota + 1
	NOT_REGULAR_FOOD
)
