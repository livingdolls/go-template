package port

import (
	"github.com/livingdolls/go-template/internal/core/dto"
	"github.com/livingdolls/go-template/internal/core/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}

type UserService interface {
	Register(req dto.RegisterUserRequest) (*model.User, error)
}
