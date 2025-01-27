package service

import (
	"fmt"
	"syncal/database"
	"syncal/internal/events/models"
)

func DeleteGroupEvents(id int) error {
	var events []models.Event

	result := database.Database.Unscoped().Where("linked_id = ?", id).Delete(&events)

	if result.Error != nil {
		return result.Error
	}

	// Verificar si no se encontraron eventos
	if result.RowsAffected == 0 {
		return fmt.Errorf("No events found with id %d", id)
	}
	return nil
}
