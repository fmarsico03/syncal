package service

import (
	"fmt"
	"syncal/database"
	"syncal/internal/events/models"
	"syncal/internal/events/request"
	"syncal/internal/participant/service"
	modelsUser "syncal/internal/users/models"
	"syncal/internal/utils/utilEmail"
	"syncal/internal/utils/utilRepeatType"
	"syncal/internal/utils/utilTime"
	"time"
)

func CreateEvent(req request.CreateEventRequest) (uint, error) {
	user, err := utilEmail.FindUserByEmail(req.EmailCreator)
	if err != nil {
		return 0, err
	}
	if !req.Repeat.IsValid() && len(req.Repeat.TypeOf) > 0 {
		return 0, fmt.Errorf("invalid repeat")
	}
	typeOf := req.Repeat.TypeOf
	var id uint

	if typeOf == utilRepeatType.Monthly {
		id, err = CreateMonthlyEvents(user, req, typeOf)
	} else if typeOf == utilRepeatType.Weekly {
		id, err = CreateWeeklyEvents(user, req, typeOf)
	} else if typeOf == utilRepeatType.Daily {
		id, err = CreateDailyEvents(user, req, typeOf)
	} else if typeOf == utilRepeatType.OneTime {
		id = CreateSimpleEvent(user, req, typeOf)
	} else {
		id = CreateSimpleEvent(user, req, typeOf)
	}

	return id, nil
}

func createBaseEvent(user modelsUser.User, req request.CreateEventRequest, typeOf utilRepeatType.RepeatType) models.Event {
	return models.Event{
		Title:       req.Title,
		CreatedBy:   user,
		Description: req.Description,
		Location:    req.Location,
		MeetLink:    req.MeetLink,
		Start:       req.Start,
		End:         req.End,
		Type:        typeOf,
	}
}

func CreateSimpleEvent(user modelsUser.User, req request.CreateEventRequest, typeOf utilRepeatType.RepeatType) uint {
	event := createBaseEvent(user, req, typeOf)
	database.Database.Create(&event)

	for _, email := range req.Participants {
		service.CreateParticipants(event, email, time.Time{})
	}

	return event.ID
}

func CreateRecurringEvents(user modelsUser.User, req request.CreateEventRequest, typeOf utilRepeatType.RepeatType, dateCalculator func(startDate, endDate time.Time, iteration int) (time.Time, time.Time)) (uint, error) {
	var linkedId uint
	var eventOrigin models.Event
	startDate := req.Start
	endDate := req.End

	if err := utilTime.ValidateDates(req.Start, req.End); err != nil {
		return 0, err
	}

	for i := 0; i < req.Repeat.Value; i++ {
		newStart, newEnd := dateCalculator(startDate, endDate, i)

		req.Start = newStart
		req.End = newEnd

		event := createBaseEvent(user, req, typeOf)

		if i == 0 {
			if err := database.Database.Create(&event).Error; err != nil {
				return 0, fmt.Errorf("failed to create the first event: %w", err)
			}
			linkedId = event.ID
			if err := database.Database.Model(&event).Update("linked_id", linkedId).Error; err != nil {
				return 0, fmt.Errorf("failed to update the first event with linked_id: %w", err)
			}
			eventOrigin = event
		} else {
			event.LinkedId = linkedId
			if err := database.Database.Create(&event).Error; err != nil {
				return 0, fmt.Errorf("failed to create recurring event: %w", err)
			}
		}
	}

	for _, email := range req.Participants {
		service.CreateParticipants(eventOrigin, email, req.Start)
	}

	return linkedId, nil
}

func CreateMonthlyEvents(user modelsUser.User, req request.CreateEventRequest, typeOf utilRepeatType.RepeatType) (uint, error) {
	return CreateRecurringEvents(user, req, typeOf, utilTime.CalculateMonthlyDates)
}

func CreateWeeklyEvents(user modelsUser.User, req request.CreateEventRequest, typeOf utilRepeatType.RepeatType) (uint, error) {
	return CreateRecurringEvents(user, req, typeOf, utilTime.CalculateWeeklyDates)
}

func CreateDailyEvents(user modelsUser.User, req request.CreateEventRequest, typeOf utilRepeatType.RepeatType) (uint, error) {
	return CreateRecurringEvents(user, req, typeOf, utilTime.CalculateDailyDates)
}
