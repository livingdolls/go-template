package port

import (
	"github.com/livingdolls/go-template/internal/core/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}
