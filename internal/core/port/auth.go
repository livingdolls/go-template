package port

import (
	"github.com/livingdolls/go-template/internal/core/dto"
	"github.com/livingdolls/go-template/internal/core/model"
)

type AuthService interface {
	Register(req dto.RegisterUserRequest) (*model.User, error)
}
