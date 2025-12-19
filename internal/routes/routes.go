package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Parthh191/backendtask/internal/handler"
	"github.com/Parthh191/backendtask/internal/middleware"
)

func SetupRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	// Apply global middleware
	router.Use(middleware.CORSMiddleware())

	// Health check endpoint
	router.GET("/health", userHandler.HealthCheck)

	// Users API routes
	usersGroup := router.Group("/api/v1")
	{
		// Create user
		usersGroup.POST("/users", userHandler.CreateUser)

		// Get all users
		usersGroup.GET("/users", userHandler.GetAllUsers)

		// Search users by name
		usersGroup.GET("/users/search", userHandler.SearchUsers)

		// Get user by ID
		usersGroup.GET("/users/:id", userHandler.GetUserByID)

		// Update user
		usersGroup.PUT("/users/:id", userHandler.UpdateUser)

		// Delete user
		usersGroup.DELETE("/users/:id", userHandler.DeleteUser)
	}
}
