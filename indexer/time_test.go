package indexer

import (
	"reflect"
	"testing"
)

func TestGetTimeRange(t *testing.T) {
	cases := []struct {
		startTime      string
		endTime        string
		expectedResult []string
		description    string
	}{
		{
			"11AM",
			"5PM",
			[]string{"11AM", "12PM", "1PM", "2PM", "3PM", "4PM", "5PM"},
			"start Time in AM, end time in PM",
		},
		{
			"1AM",
			"5AM",
			[]string{"1AM", "2AM", "3AM", "4AM", "5AM"},
			"start time and end time in AM",
		},
		{
			"7PM",
			"10PM",
			[]string{"7PM", "8PM", "9PM", "10PM"},
			"start time and end time in PM",
		},
		{
			"10PM",
			"5AM",
			[]string{"10PM", "11PM", "12AM", "1AM", "2AM", "3AM", "4AM", "5AM"},
			"start Time in PM, end time in AM",
		},
		{
			"7AM",
			"5AM",
			[]string{},
			"invalid time range",
		},
	}
	for _, tc := range cases {
		result, err := getTimeRange(tc.startTime, tc.endTime)
		if len(result) == 0 && len(tc.expectedResult) == 0 && err != nil {
			continue
		}

		if err != nil {
			t.Errorf("Error while gettign time range for startTime %s and endTime %s", tc.startTime, tc.endTime)
		}

		if !reflect.DeepEqual(result, tc.expectedResult) {
			t.Errorf("%s: time range %s-%s got %s, want %s", tc.description, tc.startTime, tc.endTime, result, tc.expectedResult)
		}
	}
}

func TestGetStartEndTimeFromDelivery(t *testing.T) {
	cases := []struct {
		delivery          string
		expectedStartTime string
		expectedEndTime   string
	}{
		{
			"Thursday 6AM - 7PM",
			"6AM",
			"7PM",
		},
		{
			"Friday 12AM - 5PM",
			"12AM",
			"5PM",
		},
	}
	for _, tc := range cases {
		startTime, endTime := getStartEndTimeFromDelivery(tc.delivery)
		if startTime != tc.expectedStartTime {
			t.Errorf("StartTime got %s, want %s", startTime, tc.expectedStartTime)
		}
		if endTime != tc.expectedEndTime {
			t.Errorf("EndTime got %s, want %s", endTime, tc.expectedEndTime)
		}
	}
}
