package constants

type TaskStatus int

const AllStatus int = 0

const (
	Todo TaskStatus = 1 + iota
	InProgress
	Done
	Backlog
)

var TaskTypeList = []int{int(Todo), int(InProgress), int(Done), int(Backlog)}
