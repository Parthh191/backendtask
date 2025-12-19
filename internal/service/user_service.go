package service

import (
	"fmt"
	"time"

	"github.com/Parthh191/backendtask/internal/models"
	"github.com/Parthh191/backendtask/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func New(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CalculateAge calculates age from date of birth
func (s *UserService) CalculateAge(dob time.Time) int {
	today:=time.Now()
	age:=today.Year() - dob.Year()

	if today.Month() < dob.Month() || (today.Month()==dob.Month() && today.Day() < dob.Day()) {
		age--
	}

	return age
}

// CreateUser creates a new user with validation
func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.UserResponse, error) {
	dob, err:=time.Parse("2006-01-02", req.DOB)
	if err!=nil {
		return nil, fmt.Errorf("invalid date format, use YYYY-MM-DD: %w", err)
	}
	if dob.After(time.Now()) {
		return nil, fmt.Errorf("date of birth cannot be in the future")
	}
	user:=&models.User{
		Name: req.Name,
		DOB:  dob,
	}
	createdUser, err:=s.repo.CreateUser(user)
	if err!=nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return s.userToResponse(createdUser), nil
}

// GetUserByID retrieves a user and includes calculated age
func (s *UserService) GetUserByID(id int) (*models.UserResponse, error) {
	user, err:=s.repo.GetUserByID(id)
	if err!=nil {
		return nil, err
	}

	if user==nil {
		return nil, fmt.Errorf("user not found")
	}

	return s.userToResponse(user), nil
}

// GetAllUsers retrieves all users with calculated ages
func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {
	users, err:=s.repo.GetAllUsers()
	if err!=nil {
		return nil, err
	}

	var responses []models.UserResponse
	for _, user:=range users {
		responses=append(responses, *s.userToResponse(&user))
	}

	return responses, nil
}

// UpdateUser updates a user with validation
func (s *UserService) UpdateUser(id int, req *models.UpdateUserRequest) (*models.UserResponse, error) {
	existingUser, err:=s.repo.GetUserByID(id)
	if err!=nil {
		return nil, err
	}

	if existingUser==nil {
		return nil, fmt.Errorf("user not found")
	}
	if req.Name != "" {
		existingUser.Name = req.Name
	}

	if req.DOB!="" {
		dob, err:=time.Parse("2006-01-02", req.DOB)
		if err!=nil {
			return nil, fmt.Errorf("invalid date format, use YYYY-MM-DD: %w", err)
		}
		if dob.After(time.Now()) {
			return nil, fmt.Errorf("date of birth cannot be in the future")
		}
		existingUser.DOB=dob
	}

	updatedUser, err:=s.repo.UpdateUser(id, existingUser)
	if err!=nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return s.userToResponse(updatedUser), nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int) error {
	user, err:=s.repo.GetUserByID(id)
	if err!=nil {
		return err
	}

	if user==nil {
		return fmt.Errorf("user not found")
	}

	return s.repo.DeleteUser(id)
}

// GetUserByName retrieves users by name search
func (s *UserService) GetUserByName(name string) ([]models.UserResponse, error) {
	users, err:=s.repo.GetUserByName(name)
	if err!=nil {
		return nil, err
	}

	var responses []models.UserResponse
	for _, user:=range users {
		responses=append(responses, *s.userToResponse(&user))
	}

	return responses, nil
}

func (s *UserService) userToResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		DOB:       user.DOB.Format("2006-01-02"),
		Age:       s.CalculateAge(user.DOB),
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
