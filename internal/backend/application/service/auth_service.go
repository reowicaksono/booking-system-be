package service

import (
	"booking-system/internal/backend/application/dto"
	"booking-system/internal/backend/domain/entity"
	"booking-system/internal/backend/domain/enum"
	"booking-system/internal/backend/infrastructure/config"
	"booking-system/internal/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// auth implement
type AuthService struct {
	db   *gorm.DB
	conf *config.MysqlDataConfig
}

// auth service contract
func NewAuthService(db *gorm.DB, conf *config.MysqlDataConfig) *AuthService {
	return &AuthService{
		db:   db,
		conf: conf,
	}
}

// method auth
func (s *AuthService) Login(c *gin.Context, body dto.WebLoginRequest) (*dto.LoginResponse, *dto.ApiStatus) {
	var user entity.User

	err := s.db.Model(&entity.User{}).Preload("Profile").Preload("AdminUser").Where("email = ?", body.Email).Take(&user).Error
	if err != nil {
		return nil, dto.ApiStatusBadRequest.WithReason("Email atau password salah")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, dto.ApiStatusBadRequest.WithReason("Email atau password salah")
	}

	refreshToken := utils.Hash256(user.ID, time.Now())

	claims := dto.JWTClaims{
		UserId:  user.ID,
		Client:  enum.ClientWeb,
		Role:    user.Role,
		IsLogin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "booking-system",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: &jwt.NumericDate{},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        utils.Hash256(user.ID, time.Now()),
		},
	}

	if user.AdminUser != nil {
		claims.AdminId = user.AdminUser.AdminId
	}

	if user.Profile != nil {
		claims.AdminId = user.Profile.AdminId
	}

	token, err2 := utils.MarshalJwt(s.conf.Secret, &claims)

	if err2 != nil {
		return nil, dto.ApiStatusError.WithReason("Failed to create token: " + err2.Error())
	}

	return &dto.LoginResponse{
		RefreshToken: refreshToken,
		AccessToken:  token.TokenString,
	}, nil
}

// Register handles user registration
func (s *AuthService) Register(c *gin.Context, body dto.RegisterRequest) (*dto.RegisterResponse, *dto.ApiStatus) {
	var user entity.User

	if err := s.db.Where("email = ?", body.Email).First(&user).Error; err == nil {
		return nil, dto.ApiStatusConflict.WithReason("Email already registered")
	}
	if err := s.db.Where("username = ?", body.Username).First(&user).Error; err == nil {
		return nil, dto.ApiStatusConflict.WithReason("Username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, dto.ApiStatusError.WithReason("Failed to hash password: " + err.Error())
	}

	user = entity.User{
		Username: body.Username,
		Name:     body.Name,
		Email:    body.Email,
		Role:     body.Role,
		Password: string(hashedPassword),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, dto.ApiStatusError.WithReason("Failed to create user: " + err.Error())
	}

	return &dto.RegisterResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
