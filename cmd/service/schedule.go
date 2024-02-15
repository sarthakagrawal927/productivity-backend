package service

type HourMinute struct {
	Hour   int
	Minute int
}

type Schedule struct {
	StartTime HourMinute
	EndTime   HourMinute
}

var sleepSchedule = Schedule{
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
var officeSchedule = Schedule{
	StartTime: HourMinute{
		Hour:   12,
		Minute: 00,
	},
	EndTime: HourMinute{
		Hour:   19,
		Minute: 30,
	},
}

func createSchedule() {
	// fetch habits & tasks
}
