package local

import (
	"context"
	"github.com/BigGold1310/spend-smart-save/internal/domain"
	"github.com/BigGold1310/spend-smart-save/internal/repository"
	"math/rand"
)

func NewUserRepository() repository.User {
	ur := userRepository{users: []domain.User{
		{ID: rand.Int(), Name: "Petra", Email: "petra@example.com"},
		{ID: rand.Int(), Name: "Peter", Email: "peter@example.com"},
	}}
	return ur
}

type userRepository struct {
	users []domain.User
}

func (u userRepository) GetByID(id int) (*domain.User, error) {
	for i, user := range u.users {
		if user.ID == id {
			return &u.users[i], nil
		}
	}
	return nil, nil
}

func (u userRepository) GetAll() ([]domain.User, error) {
	return u.users, nil
}

func (u userRepository) Create(ctx context.Context, user *domain.User) error {
	user.ID = rand.Int()
	u.users = append(u.users, *user)
	return nil
}
