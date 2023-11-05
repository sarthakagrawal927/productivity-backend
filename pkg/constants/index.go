package constants

type TaskStatus int

const AllStatus int = 0

const (
	Todo TaskStatus = 1 + iota
	InProgress
	Done
	Backlog
)

const (
	NoPriority int = 0
	LowPriority
	MediumPriority
	HighPriority
)

const (
	Idea int = iota + 1
	Gratitude
	MindClear
	DayPlanning
	DayWrap
	Event
)

const (
	NoComplexity int = 0
	LowComplexity
	MediumComplexity
	HighComplexity
)

var TaskTypeList = []int{int(Todo), int(InProgress), int(Done), int(Backlog)}
