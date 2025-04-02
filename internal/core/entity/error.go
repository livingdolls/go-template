package entity

import "errors"

// General Errors
var (
	ErrInternal        = errors.New("internal error")
	ErrDataNotFound    = errors.New("data not found")
	ErrNoUpdatedData   = errors.New("no data to update")
	ErrConflictingData = errors.New("data conflicts with existing unique column")
)

// Authentication & Authorization Errors
var (
	ErrInvalidTokenSymmetricKey   = errors.New("invalid token key size")
	ErrTokenCreation              = errors.New("error creating token")
	ErrExpiredToken               = errors.New("access token has expired")
	ErrInvalidToken               = errors.New("access token is invalid")
	ErrInvalidCredentials         = errors.New("invalid email or password")
	ErrEmptyAuthorizationHeader   = errors.New("authorization header is missing")
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")
	ErrInvalidAuthorizationType   = errors.New("authorization type is not supported")
	ErrUnauthorized               = errors.New("user is unauthorized to access this resource")
	ErrForbidden                  = errors.New("user is forbidden from accessing this resource")
	ErrEmailAlreadyExits          = errors.New("email address is exits")
)

// Session & Security Errors
var (
	ErrNoMatchPassword  = errors.New("passwords do not match")
	ErrSessionBlocked   = errors.New("session is blocked")
	ErrMissmatchSession = errors.New("session mismatch")
	ErrCaptchaInvalid   = errors.New("invalid captcha")
)
