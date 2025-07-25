package router

import (
	"booking-system/internal/backend/domain/enum"
	"booking-system/internal/backend/infrastructure/config"
	"booking-system/internal/backend/interface/controller"
	"booking-system/internal/backend/interface/handler"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UseAuthRouter(ctx context.Context, r *gin.RouterGroup, db *gorm.DB) {
	conf := ctx.Value(enum.ConfigCtxKey).(*config.MysqlDataConfig)
	contexDB := ctx.Value(enum.GormCtxKey).(*gorm.DB)
	responseHandler := handler.NewResponseHandler()

	ctrl := controller.NewAuthController(conf, contexDB)

	auth := r.Group("/auth")

	auth.POST("/login", responseHandler.WrapController(ctrl.WebLogin))
	auth.POST("/register", responseHandler.WrapController(ctrl.Register))
}
