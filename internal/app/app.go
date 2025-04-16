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

func InitApp(db *sql.DB, logger *zap.Logger) *handlers.UserHandler {
	userRepo := repository.NewUserRepository(db, logger)
	userService := service.NewUserService(userRepo, logger)
	userHandler := handlers.NewUserHandler(userService)
	return userHandler
}

func Run() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)

	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("failed to close  database: %v", err)

		}
	}(db)

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("failed to initialize logger")
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal("failed to sync logger")
		}
	}(logger)

	userHandler := InitApp(db, logger)

	r := routes.SetupRoutes(userHandler)

	if err := r.Run(); err != nil {
		log.Fatal("failed to run server")
	}

}
