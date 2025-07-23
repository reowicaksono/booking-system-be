package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserId  uint   `json:"user_id"`
	Client  string `json:"client"`
	Role    string `json:"role"`
	IsLogin bool   `json:"is_login"`
	jwt.RegisteredClaims
}

func (jw JWTClaims) IsRole(role string) bool {
	return jw.Role == role
}

func (jw JWTClaims) IsCliient(client string) bool {
	return jw.Client == client
}

type JWTToken struct {
	ExpiredAt   time.Time `json:"expired_at"`
	CreatedAt   time.Time `json:"created_at"`
	TokenString string    `json:"token_string"`
}
