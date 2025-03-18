package http

import (
	"github.com/gin-gonic/gin"
	"github.com/livingdolls/go-template/internal/core/port"
)

func NewRouter(db port.DatabasePort) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())

	return router
}
