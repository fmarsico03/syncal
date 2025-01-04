package main

import (
	"github.com/gin-gonic/gin"
	"syncal/database"
	"syncal/internal/events"
)

func main() {

	router := gin.Default()

	//Migracion
	database.Migrate()

	//events
	events.RegisterRoutes(router)

	//users
	//Routes and migration

	err := router.Run()
	if err != nil {
		return
	}

}
