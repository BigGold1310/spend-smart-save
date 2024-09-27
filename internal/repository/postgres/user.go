package postgres

import (
	"github.com/BigGold1310/spend-smart-save/internal/domain"
	"github.com/jmoiron/sqlx"
)

type SQLUserRepository struct {
	db *sqlx.DB
}

// NewSQLUserRepository creates a new SQLUserRepository
func NewSQLUserRepository(db *sqlx.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

// GetByID retrieves a user by their ID
func (r *SQLUserRepository) GetByID(id int) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.Get(user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Create inserts a new user into the database
func (r *SQLUserRepository) Create(user *domain.User) error {
	_, err := r.db.NamedExec(`
		INSERT INTO users (name, email)
		VALUES (:name, :email)
	`, user)
	return err
}

// Update updates an existing user in the database
func (r *SQLUserRepository) Update(user *domain.User) error {
	_, err := r.db.NamedExec(`
		UPDATE users
		SET name=:name, email=:email
		WHERE id=:id
	`, user)
	return err
}

// Delete removes a user from the database
func (r *SQLUserRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
