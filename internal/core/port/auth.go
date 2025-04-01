package port

import (
	"context"

	"github.com/livingdolls/go-template/internal/core/dto"
	"github.com/livingdolls/go-template/internal/core/model"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterUserRequest) (*model.User, error)
}
