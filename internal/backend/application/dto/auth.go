package dto

type WebLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type RenewTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SessionResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	Role        string          `json:"role"`
	PhoneNumber string          `json:"phone_number"`
	Bucket      *BucketResponse `json:"bucket"`
}

type BucketResponse struct {
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required|email"`
}

type ResetPasswordRequest struct {
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// register request and response structures
type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required|email"`
	Role     string `json:"role" validate:"required|enum:SUPERADMIN,ADMIN,USER"`
	Password string `json:"password" validate:"required|minLen:6"`
}

type RegisterResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
