package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/livingdolls/go-template/internal/core/dto"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/pkg/response"
)

type AuthHandler struct {
	userService port.AuthService
}

func NewAuthHandler(userService port.AuthService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleValidationError(c, err)
		return
	}

	user, err := a.userService.Register(req)

	if err != nil {
		response.HandleErrorResponse(c, err)
		return
	}

	userResponse := &dto.RegisterUserResponse{
		Id:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}

	response.HandleSuccessResponse(c, http.StatusCreated, "Register successfully", userResponse)
}
