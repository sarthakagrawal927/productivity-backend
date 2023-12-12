package constants

const AllStatus uint = 0

const (
	Todo uint = 1 + iota
	InProgress
	Done
	Backlog
)

const (
	NoPriority uint = iota + 0
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
	NoComplexity uint = iota + 0
	LowComplexity
	MediumComplexity
	HighComplexity
)

const DefaultPageSize = 20

var (
	TaskTypeList       = []uint{uint(Todo), uint(InProgress), uint(Done), uint(Backlog)}
	PriorityTypeList   = []uint{uint(NoPriority), uint(LowPriority), uint(MediumPriority), uint(HighPriority)}
	ComplexityTypeList = []uint{uint(NoComplexity), uint(LowComplexity), uint(MediumComplexity), uint(HighComplexity)}
)
