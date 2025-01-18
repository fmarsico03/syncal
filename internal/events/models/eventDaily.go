package models

type EventDaily struct {
	Event  Event
	Always bool
	Days   []DayOfWeek
}
