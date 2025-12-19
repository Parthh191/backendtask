package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/Parthh191/backendtask/config"
	"github.com/Parthh191/backendtask/internal/handler"
	"github.com/Parthh191/backendtask/internal/logger"
	"github.com/Parthh191/backendtask/internal/middleware"
	"github.com/Parthh191/backendtask/internal/repository"
	"github.com/Parthh191/backendtask/internal/routes"
	"github.com/Parthh191/backendtask/internal/service"
)
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}
	appLogger:=logger.New()
	appLogger.Info("Starting server on port %s in %s environment", cfg.Port, cfg.Env)
	appLogger.Info("Connecting to database: %s", cfg.DatabaseURL)
	db, err:=sql.Open("postgres", cfg.DatabaseURL)
	if err!=nil {
		appLogger.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	defer db.Close()
	if err:=db.Ping(); err != nil {
		appLogger.Error("Failed to ping database: %v", err)
		os.Exit(1)
	}
	appLogger.Info("Successfully connected to database")
	userRepo:=repository.New(db)
	userService:=service.New(userRepo)
	userHandler:=handler.New(userService, appLogger)
	if cfg.Env=="production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router:=gin.New()
	router.Use(middleware.LoggingMiddleware(appLogger))
	router.Use(middleware.ErrorHandlingMiddleware())
	router.Use(gin.Recovery())
	routes.SetupRoutes(router, userHandler)
	appLogger.Info("Server is running on http://localhost:%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		appLogger.Error("Failed to start server: %v", err)
		os.Exit(1)
	}
}
