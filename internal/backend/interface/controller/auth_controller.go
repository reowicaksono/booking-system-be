package controller

import (
	"booking-system/internal/backend/application/dto"
	"booking-system/internal/backend/application/service"
	"booking-system/internal/backend/infrastructure/config"
	"booking-system/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	auth service.AuthService

	config config.MysqlDataConfig
}

func NewAuthController(conf *config.MysqlDataConfig, db *gorm.DB) *AuthController {
	return &AuthController{
		auth:   *service.NewAuthService(db, conf),
		config: *conf,
	}
}

func (ctrl AuthController) WebLogin(c *gin.Context) error {
	var body dto.WebLoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.ResponseAPI{
			Status: dto.ApiStatusUnprocessableEntity,
		})
	}

	if valid := utils.ValidateStruct(body); !valid.Empty() {
		c.JSON(http.StatusBadRequest, dto.ResponseAPI{
			Status: dto.ApiStatusBadRequest.WithReason(valid.One()),
		})
		return nil
	}

	response, status := ctrl.auth.Login(c, body)

	if status != nil {
		c.JSON(status.Code, dto.ResponseAPI{
			Status: status,
		})
		return nil
	}

	c.JSON(http.StatusOK, dto.ResponseAPI{
		Status: dto.ApiStatusOK,
		Data:   response,
	})
	return nil
}

// Register handles user registration
func (ctrl AuthController) Register(c *gin.Context) error {
	var body dto.RegisterRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.ResponseAPI{
			Status: dto.ApiStatusUnprocessableEntity,
		})
		return nil
	}

	if valid := utils.ValidateStruct(body); !valid.Empty() {
		c.JSON(http.StatusBadRequest, dto.ResponseAPI{
			Status: dto.ApiStatusBadRequest.WithReason(valid.One()),
		})
		return nil
	}

	response, status := ctrl.auth.Register(c, body)

	if status != nil {
		c.JSON(status.Code, dto.ResponseAPI{
			Status: status,
		})
		return nil
	}

	c.JSON(http.StatusCreated, dto.ResponseAPI{
		Status: dto.ApiStatusOK,
		Data:   response,
	})
	return nil
}
