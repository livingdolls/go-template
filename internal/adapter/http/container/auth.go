package container

import (
	"github.com/livingdolls/go-template/internal/adapter/http/handler"
	"github.com/livingdolls/go-template/internal/adapter/repository"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/core/service"
)

type AuthContainer struct {
	UserRepo    port.UserRepository
	AuthService port.AuthService
	AuthHanlder handler.AuthHandler
}

func NewAuthContainer(db port.DatabasePort) *AuthContainer {
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	return &AuthContainer{
		UserRepo:    userRepo,
		AuthService: authService,
		AuthHanlder: *authHandler,
	}
}
