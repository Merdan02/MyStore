package routes

//
//import (
//	"github.com/gin-gonic/gin"
//	"mystore/internal/handlers"
//)
//
//func SetupProductRouters(handler *handlers.ProductHandler) *gin.Engine {
//	router := gin.Default()
//
//	productGroup := router.Group("/products")
//	productGroup.POST("/", handler.CreateProduct)
//	productGroup.GET("/", handler.GetAllProduct)
//	productGroup.GET("/id", handler.GetById)
//	productGroup.PUT("/:id", handler.UpdateProduct)
//	productGroup.DELETE("/:id", handler.DeleteProduct)
//	return router
//}
