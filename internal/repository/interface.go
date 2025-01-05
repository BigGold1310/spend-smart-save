package repository

import (
	"context"
	"github.com/BigGold1310/spend-smart-save/internal/domain"
)

type Account interface {
	Get(ctx context.Context, id int64) (*domain.Account, error)
	Update(ctx context.Context, account *domain.Account) error
}

type Transaction interface {
	Create(ctx context.Context, transaction *domain.Transaction) error
}

type UserRepository interface {
	GetUsers() ([]domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	CreateUser(user *domain.UserCreate) (*domain.User, error)
	UpdateUser(id int, update *domain.UserUpdate) (*domain.User, error)
}
