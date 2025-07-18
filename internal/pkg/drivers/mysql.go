package drivers

import (
	"booking-system/internal/backend/infrastructure/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnection() *gorm.DB {
	val := config.GetMysqlDataConfig(
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_USER", "root"),
		config.GetEnv("DB_PASSWORD", ""),
		config.GetEnv("DB_PORT", "3306"),
		config.GetEnv("DB_NAME", "db_booking_system"),
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", val.DBUser, val.DBPassword, val.DBHost, val.DBPort, val.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	fmt.Println("Database connection established successfully")
	return db
}
