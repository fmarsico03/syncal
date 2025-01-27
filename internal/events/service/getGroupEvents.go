package service

import (
	"fmt"
	"syncal/database"
	"syncal/internal/events/models"
	"syncal/internal/events/response"
	modelsUser "syncal/internal/users/models"
)

func GetGroupEvents(id int) ([]response.ResponseGetGroups, error) {
	var events []models.Event

	result := database.Database.Where("linked_id = ?", id).Find(&events)

	if result.Error != nil {
		return nil, result.Error
	}

	// Verificar si no se encontraron eventos
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("No events found with id %d", id)
	}
	rta := mapEventsToResponse(events)

	return rta, nil
}

func mapEventsToResponse(events []models.Event) []response.ResponseGetGroups {
	var responseGroups []response.ResponseGetGroups

	mail := GetUser(events[0].CreatedByID).Mail

	for _, event := range events {
		rta := response.ResponseGetGroups{
			Id:          event.ID,
			Title:       event.Title,
			Creator:     mail,
			Start:       event.Start,
			End:         event.End,
			Description: event.Description,
			MeetLink:    event.MeetLink,
		}

		responseGroups = append(responseGroups, rta)
	}

	return responseGroups
}

func GetUser(id uint) modelsUser.User {
	var user modelsUser.User
	database.Database.Model(&modelsUser.User{}).Where("id = ?", id).First(&user)

	return user
}
