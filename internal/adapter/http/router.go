package http

import (
	"github.com/gin-gonic/gin"
	"github.com/livingdolls/go-template/internal/core/port"
)

func SetupRouter(db port.DatabasePort, rmq port.EventPublisher) *gin.Engine {
	r := gin.Default()

	deps := NewAppContainer(db, rmq)

	apiV1 := r.Group("/api/v1")
	initV1Routes(apiV1, deps)

	return r
}

func initV1Routes(r *gin.RouterGroup, debs *AppContainer) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("register", debs.AuthContainer.AuthHanlder.Register)
	}
}
