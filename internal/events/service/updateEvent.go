package service

import (
	"fmt"
	"syncal/database"
	"syncal/internal/events/models"
	"syncal/internal/events/request"
)

func UpdateEvent(req request.UpdateEventRequest, id uint) (models.Event, error) {
	// Buscar el evento por ID
	var event models.Event
	result := database.Database.Preload("CreatedBy").First(&event, id)

	if result.Error != nil {
		return models.Event{}, fmt.Errorf("error finding event: %w", result.Error)
	}

	updates := map[string]interface{}{}

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if !req.Start.IsZero() {
		updates["start"] = req.Start
	}
	if !req.End.IsZero() {
		updates["end"] = req.End
	}
	if req.Location != "" {
		updates["location"] = req.Location
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.MeetLink != "" {
		updates["meet_link"] = req.MeetLink
	}

	saveResult := database.Database.Model(&event).Updates(updates)

	if saveResult.Error != nil {
		return models.Event{}, fmt.Errorf("Error saving updated event: %w", saveResult.Error)
	}

	return event, nil
}
