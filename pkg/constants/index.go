package constants

const AllStatus uint = 0

const (
	Todo uint = 1 + iota
	InProgress
	Done
	Backlog
)

const (
	NoPriority uint = 0
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
)

const (
	NoComplexity uint = 0
	LowComplexity
	MediumComplexity
	HighComplexity
)

var TaskTypeList = []uint{uint(Todo), uint(InProgress), uint(Done), uint(Backlog)}
