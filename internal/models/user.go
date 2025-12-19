package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	DOB       time.Time `json:"dob"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse includes the calculated age
type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DOB       string `json:"dob"` 
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateUserRequest for API input
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	DOB  string `json:"dob" binding:"required"` 
}

// UpdateUserRequest for API input
type UpdateUserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"` 
}
