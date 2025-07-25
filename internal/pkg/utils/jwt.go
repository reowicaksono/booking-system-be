package utils

import (
	"booking-system/internal/backend/application/dto"

	"github.com/golang-jwt/jwt/v5"
)

func UnMarshalJwt(secret, tokenString string) (*dto.JWTClaims, error) {
	var claims dto.JWTClaims

	opt := func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, opt)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenNotValidYet
	}

	return &claims, nil

}

func MarshalJwt(secret string, claims *dto.JWTClaims) (*dto.JWTToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &dto.JWTToken{
		CreatedAt:   claims.IssuedAt.Time,
		ExpiredAt:   claims.ExpiresAt.Time,
		TokenString: tokenString,
	}, nil
}
