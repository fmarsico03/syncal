package models

type EventMonthly struct {
	Event  Event
	Always bool
	Day    DayOfWeek
	Week   int
	Month  []int
}
