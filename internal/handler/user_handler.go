package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Parthh191/backendtask/internal/logger"
	"github.com/Parthh191/backendtask/internal/models"
	"github.com/Parthh191/backendtask/internal/service"
)
type UserHandler struct {
	service *service.UserService
	logger  *logger.Logger
}
// New creates a new UserHandler
func New(service *service.UserService, logger *logger.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}
// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err:=c.ShouldBindJSON(&req); err!=nil {
		h.logger.Error("Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.Name==""||req.DOB==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and dob are required"})
		return
	}

	user, err:=h.service.CreateUser(&req)
	if err!=nil {
		h.logger.Error("Failed to create user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("User created: %d", user.ID)
	c.JSON(http.StatusCreated, user)
}
// GetUserByID retrieves a user by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr:=c.Param("id")
	id, err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err:=h.service.GetUserByID(id)
	if err!=nil {
		h.logger.Error("Failed to get user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers retrieves all users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		h.logger.Error("Failed to get users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr:=c.Param("id")
	id, err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var req models.UpdateUserRequest
	if err:=c.ShouldBindJSON(&req); err!=nil {
		h.logger.Error("Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err:=h.service.UpdateUser(id, &req)
	if err!=nil {
		h.logger.Error("Failed to update user: %v", err)
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	h.logger.Info("User updated: %d", id)
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr:=c.Param("id")
	id, err:=strconv.Atoi(idStr)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err=h.service.DeleteUser(id)
	if err!=nil {
		h.logger.Error("Failed to delete user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	h.logger.Info("User deleted: %d", id)
	c.Status(http.StatusNoContent)
}

// SearchUsers searches users by name
func (h *UserHandler) SearchUsers(c *gin.Context) {
	name:=c.Query("name")
	if name=="" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name query parameter is required"})
		return
	}

	users, err:=h.service.GetUserByName(name)
	if err!=nil {
		h.logger.Error("Failed to search users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// HealthCheck health check endpoint
func (h *UserHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
