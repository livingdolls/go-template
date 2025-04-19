package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/livingdolls/go-template/internal/core/model"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type userRepository struct {
	db port.DatabasePort
}

func NewUserRepository(db port.DatabasePort) port.UserRepository {
	return &userRepository{db: db}
}

// CreateUser implements port.UserRepository.
func (u *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	// Mulai transaction
	tx, err := u.db.GetDatabase().BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Error("failed to begin transaction", zap.Error(err))
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Gunakan defer dengan pengecekan error
	var txErr error
	defer func() {
		if txErr != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				txErr = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, txErr)
				logger.Log.Error("rollback error", zap.Error(rbErr), zap.Error(txErr))
			}
		}
	}()

	// 1. Insert user
	userQuery := `INSERT INTO users (id, name, email, password_hash, provider, is_verified, created_at, updated_at)
                 VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())`

	if _, err = tx.ExecContext(ctx, userQuery,
		user.ID, user.Name, user.Email, user.PasswordHash, user.Provider, user.IsVerified); err != nil {

		txErr = fmt.Errorf("failed to insert user: %w", err)
		return txErr
	}

	// 2. Assign default role
	roleQuery := `INSERT INTO user_roles (user_id, role_id)
                 SELECT ?, id FROM roles WHERE name = 'user'`

	res, err := tx.ExecContext(ctx, roleQuery, user.ID)
	if err != nil {
		txErr = fmt.Errorf("failed to assign user role: %w", err)
		return txErr
	}

	// Verifikasi role benar-benar terassign
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		txErr = fmt.Errorf("failed to check role assignment: %w", err)
		return txErr
	}
	if rowsAffected == 0 {
		txErr = errors.New("default 'user' role not found in database")
		return txErr
	}

	// Commit transaction jika semua sukses
	if err := tx.Commit(); err != nil {
		txErr = fmt.Errorf("failed to commit transaction: %w", err)
		return txErr
	}

	return nil
}

// GetUserByEmail implements port.UserRepository.
func (u *userRepository) GetUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, name, email, password_hash, provider, is_verified, created_at, updated_at FROM users WHERE email = ?`
	row := u.db.GetDatabase().QueryRow(query, email)

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

// CreateVerificationToken implements port.UserRepository.
func (u *userRepository) CreateVerificationToken(ctx context.Context, token *model.VerificationToken) error {
	_, err := u.db.GetDatabase().ExecContext(ctx, `
        INSERT INTO verification_tokens 
        (token, user_id, token_type, expires_at)
        VALUES (?, ?, 'email_verification', ?)`,
		token.Token, token.UserID, token.ExpiresAt,
	)

	return err
}
