package service

import (
	"errors"
	"github.com/BigGold1310/spend-smart-save/internal/domain"
	"github.com/BigGold1310/spend-smart-save/internal/repository"
)

type UserService interface {
	ListUsers() ([]domain.User, error)
	GetUser(id int) (*domain.User, error)
	CreateUser(create *domain.UserCreate) (*domain.User, error)
	UpdateUser(id int, update *domain.UserUpdate) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) ListUsers() ([]domain.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) GetUser(id int) (*domain.User, error) {
	// Business logic for getting a user
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(create *domain.UserCreate) (*domain.User, error) {
	// Business validation
	if err := validateNewUser(create); err != nil {
		return nil, err
	}

	// Create the user
	return s.repo.CreateUser(create)
}

func (s *userService) UpdateUser(id int, update *domain.UserUpdate) (*domain.User, error) {
	// Business validation
	if err := validateUserUpdate(update); err != nil {
		return nil, err
	}

	// Check if user exists and is active
	_, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Perform the update
	return s.repo.UpdateUser(id, update)
}

// Business logic helper functions
func validateNewUser(user *domain.UserCreate) error {
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	// Add more validation rules
	return nil
}

func validateUserUpdate(update *domain.UserUpdate) error {
	if update.Password != nil && len(*update.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	// Add more validation rules
	return nil
}

func isValidUserRole(role string) bool {
	validRoles := map[string]bool{
		"admin": true,
		"user":  true,
	}
	return validRoles[role]
}
