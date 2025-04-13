package http

import (
	"github.com/livingdolls/go-template/internal/adapter/http/container"
	"github.com/livingdolls/go-template/internal/core/port"
)

type AppContainer struct {
	AuthContainer *container.AuthContainer
}

func NewAppContainer(db port.DatabasePort, publisher port.EventPublisher) *AppContainer {
	return &AppContainer{
		AuthContainer: container.NewAuthContainer(db, publisher),
	}
}
