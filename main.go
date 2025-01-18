package main

import (
	"github.com/gin-gonic/gin"
	"syncal/database"
	"syncal/internal/events"
	"syncal/internal/users"
)

func main() {

	router := gin.Default()
	api := router.Group("/api")
	//Migracion
	database.Migrate()

	//events
	events.RegisterRoutes(api)

	//users
	users.RegisterRoutes(api)

	err := router.Run()
	if err != nil {
		return
	}

}
