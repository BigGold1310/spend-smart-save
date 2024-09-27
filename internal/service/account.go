package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/BigGold1310/spend-smart-save/internal/domain"
	"github.com/BigGold1310/spend-smart-save/internal/repository"
)

type AccountService struct {
	accountRepo     repository.Account
	transactionRepo repository.Transaction
	db              *sql.DB
}

func NewAccountService(accountRepo repository.Account, transactionRepo repository.Transaction, db *sql.DB) *AccountService {
	return &AccountService{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
		db:              db,
	}
}

func (s *AccountService) TransferMoney(ctx context.Context, fromAccountID, toAccountID int64, amount float64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get the 'from' account
	fromAccount, err := s.accountRepo.Get(ctx, fromAccountID)
	if err != nil {
		return err
	}

	// Check if 'from' account has sufficient balance
	if fromAccount.Balance < amount {
		return errors.New("insufficient funds")
	}

	// Get the 'to' account
	toAccount, err := s.accountRepo.Get(ctx, toAccountID)
	if err != nil {
		return err
	}

	// Update balances
	fromAccount.Balance -= amount
	toAccount.Balance += amount

	// Update 'from' account
	err = s.accountRepo.Update(ctx, fromAccount)
	if err != nil {
		return err
	}

	// Update 'to' account
	err = s.accountRepo.Update(ctx, toAccount)
	if err != nil {
		return err
	}

	// Create transaction record
	transaction := &domain.Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
	err = s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
