package postgres

import (
	"context"
	"github.com/BigGold1310/spend-smart-save/internal/domain"
	"github.com/jmoiron/sqlx"
)

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	_, err := r.db.NamedExecContext(ctx, `
        INSERT INTO transactions (from_account_id, to_account_id, amount, created_at)
        VALUES (:from_account_id, :to_account_id, :amount, NOW())`, transaction)
	return err
}
