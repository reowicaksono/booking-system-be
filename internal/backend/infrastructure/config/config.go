package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type MysqlDataConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	Secret     string
}

func GetMysqlDataConfig(dbhost string, dbuser string, dbpassword string, dbport string, dbname string) *MysqlDataConfig {
	if dbhost == "" {
		fmt.Println("DB_HOST is not set, using default 'localhost'")
	}
	return &MysqlDataConfig{
		DBUser:     dbuser,
		DBPassword: dbpassword,
		DBHost:     dbhost,
		DBPort:     dbport,
		DBName:     dbname,
	}
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
