package app

import (
	"booking-system/database/migratios"
	"booking-system/internal/backend/infrastructure/config"
	"booking-system/internal/pkg/drivers"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *gin.Engine
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

func (app *App) RouterInit() {
	router := gin.Default()

	app.Router = router
}

func (app *App) Run() {
	port := fmt.Sprintf(":%s", config.GetEnv("APP_PORT", "8080"))
	app.Router.Run(port)
}
