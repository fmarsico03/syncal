package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"syncal/database"
	"syncal/internal/events/models"
	"syncal/internal/events/request"
	modelsUser "syncal/internal/users/models"
)

func findUserByEmail(email string) (modelsUser.User, error) {
	var user modelsUser.User
	err := database.Database.Where("mail = ?", email).First(&user).Error
	if err != nil {
		// Manejar error de usuario no encontrado o fallo en la consulta
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, fmt.Errorf("error retrieving user: %w", err)
	}
	return user, nil
}

func CreateEvent(req request.CreateEventRequest) (uint, error) {
	user, err := findUserByEmail(req.EmailCreator)
	if err != nil {
		return 0, err
	}

	var id uint
	if len(req.Months) > 0 && len(req.Weeks) > 0 && len(req.Days) > 0 {
		id = CreateComplexEvent(user, req)
	} else if len(req.Months) > 0 && len(req.Weeks) == 1 && len(req.Days) == 1 {
		id = CreateMonthlyEvent(user, req)
	} else if len(req.Weeks) == 1 && len(req.Days) == 1 {
		id = CreateWeeklyEvent(user, req)
	} else if len(req.Days) > 0 {
		id = CreateDailyEvent(user, req)
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

func CreateDailyEvent(user modelsUser.User, req request.CreateEventRequest) uint {
	dailyEvent := models.EventDaily{
		Event:  createBaseEvent(user, req),
		Days:   models.ConvertToDaysOfWeek(req.Days),
		Always: req.Always,
	}

	database.Database.Create(&dailyEvent)
	return dailyEvent.Event.ID
}

func CreateWeeklyEvent(user modelsUser.User, req request.CreateEventRequest) uint {
	weeklyEvent := models.EventWeekly{
		Event:  createBaseEvent(user, req),
		Day:    models.ConvertToDaysOfWeek(req.Days)[0],
		Week:   req.Weeks,
		Always: req.Always,
	}

	database.Database.Create(&weeklyEvent)
	return weeklyEvent.Event.ID
}

func CreateMonthlyEvent(user modelsUser.User, req request.CreateEventRequest) uint {
	monthlyEvent := models.EventMonthly{
		Event:  createBaseEvent(user, req),
		Day:    models.ConvertToDaysOfWeek(req.Days)[0],
		Week:   req.Weeks[0],
		Month:  req.Months,
		Always: req.Always,
	}

	database.Database.Create(&monthlyEvent)
	return monthlyEvent.Event.ID
}

func CreateComplexEvent(user modelsUser.User, req request.CreateEventRequest) uint {
	recurrence := models.Recurrence{
		Days:   models.ConvertToDaysOfWeek(req.Days),
		Weeks:  req.Weeks,
		Months: req.Months,
	}
	complexEvent := models.EventComplex{
		Event:      createBaseEvent(user, req),
		Recurrence: recurrence,
	}

	database.Database.Create(&complexEvent)
	return complexEvent.Event.ID
}
