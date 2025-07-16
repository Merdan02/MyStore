package routes

import (
	"github.com/gin-gonic/gin"
	"mystore/internal/handlers"
	"mystore/internal/middleware"
	"os"
)

func SetupRoutes(userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler) *gin.Engine {
	router := gin.Default()
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	userGroup := router.Group("/user")
	userGroup.GET("/", userHandler.GetAllUser)
	userGroup.GET("/id/:id", userHandler.GetUserByID)
	userGroup.GET("/email/:email", userHandler.GetUserByEmail)
	userGroup.GET("/username/:username", userHandler.GetUserByUsername)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.POST("/login", userHandler.Login)
	userGroup.POST("/", userHandler.CreateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)

	protectedUser := router.Group("/protect/user")
	protectedUser.Use(middleware.AuthMiddleware(jwtKey))
	protectedUser.GET("/me", userHandler.GetMe)

	productGroup := router.Group("/products")
	productGroup.GET("/:id", productHandler.GetById)
	productGroup.GET("/", productHandler.GetAllProduct)

	adminGroup := router.Group("/admin/products")
	adminGroup.Use(middleware.AuthMiddleware(jwtKey), middleware.AdminOnly())
	{
		adminGroup.POST("/", productHandler.CreateProduct)
		adminGroup.PUT("/:id", productHandler.UpdateProduct)
		adminGroup.DELETE("/:id", productHandler.DeleteProduct)
	}
	return router
}
