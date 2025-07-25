package router

import (
	"booking-system/internal/backend/interface/middleware"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(ctx context.Context, app *gin.Engine, db *gorm.DB) {
	app.Use(middleware.CORSMiddleware())

	api := app.Group("/api/v1")

	UseAuthRouter(ctx, api, db)
}
