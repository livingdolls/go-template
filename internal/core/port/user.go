package port

import (
	"context"

	"github.com/livingdolls/go-template/internal/core/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	CreateVerificationToken(ctx context.Context, token *model.VerificationToken) error
}
