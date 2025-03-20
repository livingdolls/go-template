package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/livingdolls/go-template/internal/core/model"
	"github.com/livingdolls/go-template/internal/core/port"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) port.UserRepository {
	return &userRepository{db: db}
}

// CreateUser implements port.UserRepository.
func (u *userRepository) CreateUser(user *model.User) error {
	query := `INSERT INTO users (id, name, email, password_hash, provider, is_verified, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())`

	_, err := u.db.Exec(query, user.ID, user.Name, user.Email, user.PasswordHash, user.Provider, user.IsVerified)
	return err
}

// GetUserByEmail implements port.UserRepository.
func (u *userRepository) GetUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, name, email, password_hash, provider, is_verified, created_at, updated_at FROM users WHERE email = ?`
	row := u.db.QueryRow(query, email)

	var user model.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Provider, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
