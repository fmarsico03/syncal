package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	subRouter := router.Group("/user")

	subRouter.POST("/", Create)
}
