package events

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"syncal/database"
	"syncal/internal/events/models"
	"syncal/internal/events/request"
	"syncal/internal/events/service"
)

func Create(c *gin.Context) {
	var req request.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	rta, err := service.CreateEvent(req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"event id": rta})
	}
}

func Get(c *gin.Context) {

}

func Update(c *gin.Context) {
	var req request.UpdateEventRequest
	id_string := c.Param("id")
	id, _ := strconv.Atoi(id_string)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	event, err := service.UpdateEvent(req, uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"events": event})
	}
}

func Delete(c *gin.Context) {
	var id int
	id_string := c.Param("id")
	id, _ = strconv.Atoi(id_string)

	result := database.Database.Unscoped().Delete(&models.Event{}, id)
	var err error
	if result.Error != nil {
		err = result.Error
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if result.RowsAffected == 0 {
		err = fmt.Errorf("No events found with id %d", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"delete": "Delete succesfull"})
	}
}

func GetById(c *gin.Context) {
	var id int
	id_string := c.Param("id")
	id, _ = strconv.Atoi(id_string)
	var event models.Event
	result := database.Database.Preload("CreatedBy").Where("id = ?", id).Find(&event)

	var err error
	if result.Error != nil {
		err = result.Error
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if result.RowsAffected == 0 {
		err = fmt.Errorf("No events found with id %d", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"events": event})
	}
}

func GetByGroup(c *gin.Context) {
	var linked_id int
	linked_id_string := c.Param("id")
	linked_id, _ = strconv.Atoi(linked_id_string)

	rta, err := service.GetGroupEvents(linked_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"events": rta})
	}
}

func DeleteByGroup(c *gin.Context) {
	var linked_id int
	linked_id_string := c.Param("id")
	linked_id, _ = strconv.Atoi(linked_id_string)

	err := service.DeleteGroupEvents(linked_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"delete": "Delete succesfull"})
	}
}
