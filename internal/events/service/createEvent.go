package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"syncal/database"
	"syncal/internal/events/models"
	"syncal/internal/events/request"
	modelsUser "syncal/internal/users/models"
	"time"
)

// Validaciones
func findUserByEmail(email string) (modelsUser.User, error) {
	var user modelsUser.User
	err := database.Database.Where("mail = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errors.New("user not found")
	}
	return user, err
}

func validateDates(startDate, endDate time.Time) error {
	if !endDate.After(startDate) {
		return errors.New("end date must be after start date")
	}
	return nil
}

func CreateEvent(req request.CreateEventRequest) (uint, error) {
	user, err := findUserByEmail(req.EmailCreator)
	if err != nil {
		return 0, err
	}
	if !req.Repeat.IsValid() {
		return 0, fmt.Errorf("invalid repeat")
	}
	var id uint
	if req.Repeat.TypeOf == models.Monthly {
		id, err = CreateMonthlyEvents(user, req)
	} else if req.Repeat.TypeOf == models.Weekly {
		id, err = CreateWeeklyEvents(user, req)
	} else if req.Repeat.TypeOf == models.Daily {
		id, err = CreateDailyEvents(user, req)
	} else if req.Repeat.TypeOf == models.OneTime {
		id = CreateSimpleEvent(user, req)
	} else {
		id = CreateSimpleEvent(user, req)
	}

	return id, nil
}

func createBaseEvent(user modelsUser.User, req request.CreateEventRequest) models.Event {
	return models.Event{
		Title:       req.Title,
		CreatedBy:   user,
		Description: req.Description,
		Location:    req.Location,
		MeetLink:    req.MeetLink,
		Start:       req.Start,
		End:         req.End,
	}
}

func CreateSimpleEvent(user modelsUser.User, req request.CreateEventRequest) uint {
	event := createBaseEvent(user, req)
	database.Database.Create(&event)

	return event.ID
}

func CreateRecurringEvents(user modelsUser.User, req request.CreateEventRequest, dateCalculator func(startDate, endDate time.Time, iteration int) (time.Time, time.Time)) (uint, error) {
	var linkedId uint
	startDate := req.Start
	endDate := req.End

	if err := validateDates(req.Start, req.End); err != nil {
		return 0, err
	}

	for i := 0; i < req.Repeat.Value; i++ {
		newStart, newEnd := dateCalculator(startDate, endDate, i)

		req.Start = newStart
		req.End = newEnd

		event := createBaseEvent(user, req)

		if i == 0 {
			if err := database.Database.Create(&event).Error; err != nil {
				return 0, fmt.Errorf("failed to create the first event: %w", err)
			}
			linkedId = event.ID
			if err := database.Database.Model(&event).Update("linked_id", linkedId).Error; err != nil {
				return 0, fmt.Errorf("failed to update the first event with linked_id: %w", err)
			}
		} else {
			event.LinkedId = linkedId
			if err := database.Database.Create(&event).Error; err != nil {
				return 0, fmt.Errorf("failed to create recurring event: %w", err)
			}
		}
	}

	return linkedId, nil
}

func CreateMonthlyEvents(user modelsUser.User, req request.CreateEventRequest) (uint, error) {
	return CreateRecurringEvents(user, req, calculateMonthlyDates)
}

func CreateWeeklyEvents(user modelsUser.User, req request.CreateEventRequest) (uint, error) {
	return CreateRecurringEvents(user, req, calculateWeeklyDates)
}

func CreateDailyEvents(user modelsUser.User, req request.CreateEventRequest) (uint, error) {
	return CreateRecurringEvents(user, req, calculateDailyDates)
}
