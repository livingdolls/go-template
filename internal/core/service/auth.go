package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/livingdolls/go-template/internal/core/dto"
	"github.com/livingdolls/go-template/internal/core/entity"
	"github.com/livingdolls/go-template/internal/core/model"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/pkg/hash"
)

type authService struct {
	userRepo port.UserRepository
}

func NewAuthService(userRepo port.UserRepository) port.AuthService {
	return &authService{userRepo: userRepo}
}

// Register implements port.UserService.
func (u *authService) Register(ctx context.Context, req dto.RegisterUserRequest) (*model.User, error) {
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

	if err := u.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	// Create token verification
	token := uuid.New().String()
	verificationToken := &model.VerificationToken{
		Token:     token,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := u.userRepo.CreateVerificationToken(ctx, verificationToken); err != nil {
		return nil, err
	}

	// TODO:: Create mail service to send verification

	return user, nil
}
