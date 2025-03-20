package dto

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/livingdolls/go-template/internal/core/entity"
)

// ErrorStatusMap maps specific errors to their HTTP status codes.
var ErrorStatusMap = map[error]int{
	entity.ErrInternal:                   http.StatusInternalServerError,
	entity.ErrDataNotFound:               http.StatusNotFound,
	entity.ErrConflictingData:            http.StatusConflict,
	entity.ErrInvalidCredentials:         http.StatusUnauthorized,
	entity.ErrUnauthorized:               http.StatusUnauthorized,
	entity.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	entity.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	entity.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	entity.ErrInvalidToken:               http.StatusUnauthorized,
	entity.ErrExpiredToken:               http.StatusUnauthorized,
	entity.ErrForbidden:                  http.StatusForbidden,
	entity.ErrNoUpdatedData:              http.StatusBadRequest,
	entity.ErrNoMatchPassword:            http.StatusUnauthorized,
	entity.ErrSessionBlocked:             http.StatusUnauthorized,
	entity.ErrMissmatchSession:           http.StatusUnauthorized,
	entity.ErrCaptchaInvalid:             http.StatusBadRequest,
}

// ErrorResponse represents the error response structure.
type ErrorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Code     int      `json:"code" example:"500"`
	Messages []string `json:"messages"`
}

// SuccessResponse represents the success response structure.
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewErrorResponse creates a structured error response.
func NewErrorResponse(code int, errMsgs []string) ErrorResponse {
	return ErrorResponse{
		Success:  false,
		Code:     code,
		Messages: errMsgs,
	}
}

// NewSuccessResponse creates a structured success response.
func NewSuccessResponse(message string, data any) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// HandleErrorResponse sends an error response in JSON format.
func HandleErrorResponse(c *gin.Context, err error) {
	statusCode, ok := ErrorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := ParseError(err)
	errResponse := NewErrorResponse(statusCode, errMsg)
	c.AbortWithStatusJSON(statusCode, errResponse)
}

// HandleValidationError sends a validation error response.
func HandleValidationError(c *gin.Context, err error) {
	errMsg := ParseValidationError(err)
	errResponse := NewErrorResponse(http.StatusBadRequest, errMsg)
	c.JSON(http.StatusBadRequest, errResponse)
}

// HandleSuccessResponse sends a success response with data.
func HandleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	resp := NewSuccessResponse(message, data)
	c.JSON(code, resp)
}

// ParseError parses general errors into a slice of strings.
func ParseError(err error) []string {
	var errMsgs []string
	if errors.As(err, &validator.ValidationErrors{}) {
		errMsgs = ParseValidationError(err)
	} else {
		errMsgs = append(errMsgs, err.Error())
	}
	return errMsgs
}

// ParseValidationError formats validation errors into readable messages.
func ParseValidationError(err error) []string {
	var errMsgs []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			errMsgs = append(errMsgs, formatValidationError(err))
		}
	}
	return errMsgs
}

// formatValidationError provides a clear error message for each validation error.
func formatValidationError(err validator.FieldError) string {
	return "Field validation for '" + err.Field() + "' failed on the '" + err.Tag() + "' tag"
}
