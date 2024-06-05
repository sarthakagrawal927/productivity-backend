package constants

import "todo/pkg/types"

// values in this file are only till user is not created
var SleepSchedule = types.ScheduleEntry{
	Label: "sleep",
	StartTime: types.HourMinute{
		Hour:   2,
		Minute: 0,
	},
	EndTime: types.HourMinute{
		Hour:   9,
		Minute: 0,
	},
}

// 12-19:30
var OfficeSchedule = types.ScheduleEntry{
	Label: "Work",
	StartTime: types.HourMinute{
		Hour:   12,
		Minute: 00,
	},
	EndTime: types.HourMinute{
		Hour:   19,
		Minute: 30,
	},
}

var FastingSchedule = types.ScheduleEntry{
	Label: "Fasting",
	StartTime: types.HourMinute{
		Hour:   19,
		Minute: 30,
	},
	EndTime: types.HourMinute{
		Hour:   12,
		Minute: 0,
	},
}
