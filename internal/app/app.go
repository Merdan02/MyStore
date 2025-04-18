package app

import (
	"database/sql"
	"go.uber.org/zap"
	"log"
	"mystore/internal/config"
	"mystore/internal/handlers"
	"mystore/internal/repository"
	"mystore/internal/routes"
	"mystore/internal/service"
)

func InitApp(db *sql.DB, logger *zap.Logger) (*handlers.UserHandler, *handlers.ProductHandler) {
	productRepo := repository.NewProductRepo(db, logger)
	productService := service.NewProductService(productRepo, logger)
	productHandler := handlers.NewProductHandler(productService)

	userRepo := repository.NewUserRepository(db, logger)
	userService := service.NewUserService(userRepo, logger)
	userHandler := handlers.NewUserHandler(userService)
	return userHandler, productHandler
}

func Run() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)

	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("DB close error: %v", err)
		}
	}()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("failed to initialize logger")
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
	}(logger)

	userHandler, productHandler := InitApp(db, logger)

	r := routes.SetupRoutes(userHandler, productHandler)

	if err := r.Run(); err != nil {
		log.Fatal("failed to run server")
	}

}
