package dto

type RegisterUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterUserResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}
