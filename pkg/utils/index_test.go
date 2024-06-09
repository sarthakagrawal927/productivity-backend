package utils_test

// write test cases for InsertElementsInSliceAfterIdx

import (
	"reflect"
	"testing"
	"todo/pkg/types"
	"todo/pkg/utils"
)

func TestInsertElementsInSliceAfterIdx(t *testing.T) {
	type args struct {
		slice    []types.ScheduleEntry
		after    int
		elements []types.ScheduleEntry
	}
	tests := []struct {
		name string
		args args
		want []types.ScheduleEntry
	}{
		{
			name: "Inserts elements in the middle",
			args: args{
				slice: []types.ScheduleEntry{
					{StartTime: types.HourMinute{Hour: 0, Minute: 0}, EndTime: types.HourMinute{Hour: 1, Minute: 0}},
					{StartTime: types.HourMinute{Hour: 2, Minute: 0}, EndTime: types.HourMinute{Hour: 3, Minute: 0}},
				},
				after: 0,
				elements: []types.ScheduleEntry{
					{StartTime: types.HourMinute{Hour: 1, Minute: 0}, EndTime: types.HourMinute{Hour: 2, Minute: 0}},
				},
			},
			want: []types.ScheduleEntry{
				{StartTime: types.HourMinute{Hour: 0, Minute: 0}, EndTime: types.HourMinute{Hour: 1, Minute: 0}},
				{StartTime: types.HourMinute{Hour: 1, Minute: 0}, EndTime: types.HourMinute{Hour: 2, Minute: 0}},
				{StartTime: types.HourMinute{Hour: 2, Minute: 0}, EndTime: types.HourMinute{Hour: 3, Minute: 0}},
			},
		},
		{
			name: "Inserts elements at the end",
			args: args{
				slice: []types.ScheduleEntry{
					{StartTime: types.HourMinute{Hour: 0, Minute: 0}, EndTime: types.HourMinute{Hour: 1, Minute: 0}},
					{StartTime: types.HourMinute{Hour: 2, Minute: 0}, EndTime: types.HourMinute{Hour: 3, Minute: 0}},
				},
				after: 1,
				elements: []types.ScheduleEntry{
					{StartTime: types.HourMinute{Hour: 3, Minute: 0}, EndTime: types.HourMinute{Hour: 4, Minute: 0}},
				},
			},
			want: []types.ScheduleEntry{
				{StartTime: types.HourMinute{Hour: 0, Minute: 0}, EndTime: types.HourMinute{Hour: 1, Minute: 0}},
				{StartTime: types.HourMinute{Hour: 2, Minute: 0}, EndTime: types.HourMinute{Hour: 3, Minute: 0}},
				{StartTime: types.HourMinute{Hour: 3, Minute: 0}, EndTime: types.HourMinute{Hour: 4, Minute: 0}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.InsertElementsInSliceAfterIdx(tt.args.slice, tt.args.elements, tt.args.after); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertElementsInSliceAfterIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}
