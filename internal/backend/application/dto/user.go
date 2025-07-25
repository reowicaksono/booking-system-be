package dto

import "time"

type GetUserRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Search   string `json:"search"`
}

type GetUserResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Address     string  `json:"address"`
	IsActive    bool    `json:"is_active"`
	Role        string  `json:"role"`
	Balance     float64 `json:"balance,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name        string `json:"name" validate:"required|ascii"`
	Username    string `json:"username" validate:"required|alphaNum"`
	Email       string `json:"email" validate:"required|email"`
	PhoneNumber string `json:"phone_number" validate:"required|phone"`
	Role        string `json:"role" validate:"required|enum:SUPERADMIN,ADMIN,USER"`
}

type UpdateUserRequest struct {
	Name        string `json:"name" validate:"required|ascii"`
	Username    string `json:"username" validate:"required|alphaNum"`
	Email       string `json:"email" validate:"required|email"`
	PhoneNumber string `json:"phone_number" validate:"required|phone"`
	Role        string `json:"role" validate:"required|enum:SUPERADMIN,ADMIN,USER"`
}
