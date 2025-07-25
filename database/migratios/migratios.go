package migratios

import (
	"booking-system/internal/backend/domain/entity"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.User{}, &entity.Bucket{}, &entity.Admin{}, &entity.Profile{}, &entity.AdminUser{})
	if err != nil {
		return err
	}
	log.Println("Database migration completed successfully")
	return nil

}
