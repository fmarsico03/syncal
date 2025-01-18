package models

type EventWeekly struct {
	Event  Event
	Always bool
	Day    DayOfWeek
	Week   []int
}
