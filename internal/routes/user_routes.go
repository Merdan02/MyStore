package routes

import (
	"github.com/gin-gonic/gin"
	"mystore/internal/handlers"
)

func SetupRoutes(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/user")
	userGroup.GET("/", userHandler.GetAllUser)
	userGroup.GET("/id/:id", userHandler.GetUserByID)
	userGroup.GET("/email/:email", userHandler.GetUserByEmail)
	userGroup.GET("/username/:username", userHandler.GetUserByUsername)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.POST("/", userHandler.CreateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)

	return router
}
