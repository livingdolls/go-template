package container

import (
	"github.com/livingdolls/go-template/internal/adapter/http/handler"
	"github.com/livingdolls/go-template/internal/adapter/repository"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/core/service"
)

type AuthContainer struct {
	UserRepo       port.UserRepository
	AuthService    port.AuthService
	AuthHanlder    handler.AuthHandler
	EventPublisher port.EventPublisher
}

func NewAuthContainer(db port.DatabasePort, publisher port.EventPublisher) *AuthContainer {
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, publisher)
	authHandler := handler.NewAuthHandler(authService)

	return &AuthContainer{
		UserRepo:       userRepo,
		AuthService:    authService,
		AuthHanlder:    *authHandler,
		EventPublisher: publisher,
	}
}
