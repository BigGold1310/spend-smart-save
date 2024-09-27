package postgres

import (
	"context"
	"github.com/BigGold1310/spend-smart-save/internal/domain"
	"github.com/jmoiron/sqlx"
)

type c struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *accountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Get(ctx context.Context, id int64) (*domain.Account, error) {
	var account domain.Account
	err := r.db.GetContext(ctx, &account, "SELECT * FROM accounts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) Update(ctx context.Context, account *domain.Account) error {
	_, err := r.db.NamedExecContext(ctx, `
        UPDATE accounts 
        SET balance = :balance, updated_at = NOW() 
        WHERE id = :id`, account)
	return err
}
