package service

import (
	"github.com/google/uuid"
	"github.com/livingdolls/go-template/internal/core/dto"
	"github.com/livingdolls/go-template/internal/core/entity"
	"github.com/livingdolls/go-template/internal/core/model"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/pkg/hash"
)

type userService struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) port.UserService {
	return &userService{userRepo: userRepo}
}

// Register implements port.UserService.
func (u *userService) Register(req dto.RegisterUserRequest) (*model.User, error) {
	isUserExit, err := u.userRepo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if isUserExit != nil {
		return nil, entity.ErrConflictingData
	}

	hashedPassword, err := hash.HashString(req.Password)

	if err != nil {
		return nil, entity.ErrInternal
	}

	user := &model.User{
		ID:           uuid.New().String(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Provider:     "local",
		IsVerified:   false,
	}

	err = u.userRepo.CreateUser(user)

	if err != nil {
		return nil, entity.ErrInternal
	}

	return user, nil
}
