package local

import (
	"errors"
	"github.com/BigGold1310/spend-smart-save/internal/domain"
	"sync"
	"time"
)

type userRepository struct {
	mu    sync.RWMutex
	users map[int]domain.User
	// Auto-incrementing ID counter
	nextID int
}

func NewUserRepository() *userRepository {
	return &userRepository{
		users:  make(map[int]domain.User),
		nextID: 1,
	}
}

func (r *userRepository) GetUsers() ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id int) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepository) CreateUser(create *domain.UserCreate) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check for unique username and email
	for _, existingUser := range r.users {
		if existingUser.Username == create.Username {
			return nil, errors.New("username already exists")
		}
		if existingUser.Email == create.Email {
			return nil, errors.New("email already exists")
		}
	}

	user := domain.User{
		ID:        r.nextID,
		Username:  create.Username,
		Email:     create.Email,
		Password:  create.Password, // Note: In real implementation, this should be hashed
		CreatedAt: time.Now(),
	}

	r.users[r.nextID] = user
	r.nextID++

	return &user, nil
}

func (r *userRepository) UpdateUser(id int, update *domain.UserUpdate) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	// Check uniqueness constraints if updating username or email
	if update.Username != nil {
		for _, existingUser := range r.users {
			if existingUser.ID != id && existingUser.Username == *update.Username {
				return nil, errors.New("username already exists")
			}
		}
		user.Username = *update.Username
	}

	if update.Email != nil {
		for _, existingUser := range r.users {
			if existingUser.ID != id && existingUser.Email == *update.Email {
				return nil, errors.New("email already exists")
			}
		}
		user.Email = *update.Email
	}

	if update.Password != nil {
		user.Password = *update.Password // Note: In real implementation, this should be hashed
	}

	r.users[id] = user
	return &user, nil
}
