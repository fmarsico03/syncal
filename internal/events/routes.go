package events

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	subRouter := router.Group("/events")

	subRouter.POST("/", Create)
	subRouter.GET("/", Get)
	subRouter.DELETE("/:id", Delete)
	subRouter.PUT("/:id", Update)
	subRouter.GET("/:id", GetById)
	subRouter.GET("/group/:id", GetByGroup)
	subRouter.DELETE("/group/:id", DeleteByGroup)
}
