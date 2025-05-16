package routes

import (
	"github.com/gin-gonic/gin"
	"mystore/internal/handlers"
)

func SetupRoutes(userHandler *handlers.UserHandler, handler *handlers.ProductHandler) *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/user")
	userGroup.GET("/", userHandler.GetAllUser)
	userGroup.GET("/id/:id", userHandler.GetUserByID)
	userGroup.GET("/email/:email", userHandler.GetUserByEmail)
	userGroup.GET("/username/:username", userHandler.GetUserByUsername)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.POST("/login", userHandler.Login)
	userGroup.POST("/", userHandler.CreateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)

	productGroup := router.Group("/products")
	productGroup.POST("/", handler.CreateProduct)
	productGroup.GET("/", handler.GetAllProduct)
	productGroup.GET("/:id", handler.GetById)
	productGroup.PUT("/:id", handler.UpdateProduct)
	productGroup.DELETE("/:id", handler.DeleteProduct)

	return router
}
