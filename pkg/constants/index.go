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
)

// as it is hard to predict the time a particular task will take, basing it on complexity
const (
	NoComplexity     uint = iota + 1 // 15min
	LowComplexity                    // 30min
	MediumComplexity                 // 60min
	HighComplexity                   // 120min
)

const (
	HabitDailyFreq uint = iota + 1
	HabitWeeklyFreq
)

const (
	HabitActive uint = iota + 1
	HabitPaused
)

const (
	HabitTimeMode  uint = iota + 1
	HabitCountMode      // for example, 10 pushups
	HabitMlMode         // for example, 1 liter of water
)

const DefaultPageSize = 20

var (
	TaskTypeList       = []uint{uint(Todo), uint(InProgress), uint(Done), uint(Backlog)}
	PriorityTypeList   = []uint{uint(NoPriority), uint(LowPriority), uint(MediumPriority), uint(HighPriority)}
	ComplexityTypeList = []uint{uint(NoComplexity), uint(LowComplexity), uint(MediumComplexity), uint(HighComplexity)}

	JournalTypeList = []uint{AllStatus, (Idea), uint(Gratitude), uint(MindClear), uint(DayPlanning), uint(DayWrap), uint(Event), uint(FoodLog)}

	HabitFreqTypeList = []uint{uint(HabitDailyFreq), uint(HabitWeeklyFreq)}
	HabitStatusList   = []uint{uint(HabitPaused), uint(HabitActive)}
	HabitModeList     = []uint{uint(HabitTimeMode), uint(HabitCountMode), uint(HabitMlMode)}
)

const (
	ENTITY_TASK    uint = iota + 1
	ENTITY_HABIT        // 2
	ENTITY_PROJECT      // 3
	ENTITY_GOAL         // 4
)
