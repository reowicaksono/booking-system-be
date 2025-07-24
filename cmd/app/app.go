package app

import (
	"booking-system/database/migratios"
	"booking-system/internal/backend/domain/enum"
	"booking-system/internal/backend/infrastructure/config"
	"booking-system/internal/backend/interface/router"
	"booking-system/internal/pkg/drivers"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *gin.Engine
	ctx    context.Context
}

// DB initializes the database connection
func (app *App) DBInit() {
	mysqConn := drivers.MysqlConnection()
	app.DB = mysqConn

	err := migratios.Migrate(app.DB)

	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

func (app *App) RouterInit(ctx context.Context) {
	routers := gin.Default()
	app.ctx = ctx
	app.Router = routers

	ctxWithDb := context.WithValue(ctx, enum.GormCtxKey, app.DB)

	mysqlConfig := &config.MysqlDataConfig{
		DBHost:     config.GetEnv("DB_HOST", "localhost"),
		DBPort:     config.GetEnv("DB_PORT", "3306"),
		DBUser:     config.GetEnv("DB_USER", "root"),
		DBPassword: config.GetEnv("DB_PASSWORD", ""),
		DBName:     config.GetEnv("DB_NAME", "booking_system"),
	}

	ctxWithConfig := context.WithValue(ctxWithDb, enum.ConfigCtxKey, mysqlConfig)

	router.SetupRoutes(ctxWithConfig, routers, app.DB)
}

func (app *App) Run() {
	port := fmt.Sprintf(":%s", config.GetEnv("APP_PORT", "8080"))
	app.Router.Run(port)
}
