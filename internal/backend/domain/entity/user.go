package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Username    string `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email       string `gorm:"type:varchar(50);unique" json:"email"`
	PhoneNumber string `gorm:"type:varchar(15)" json:"phone_number"`
	Password    string `gorm:"type:varchar(128)" json:"-"`
	Address     string `gorm:"type:varchar(200)" json:"address"`
	IsActive    bool   `gorm:"type:boolean;default:false" json:"is_active"`
	Role        string `gorm:"type:varchar(10)" json:"role"`
	WebToken    string `gorm:"type:text" json:"-"`
	DeviceToken string `gorm:"type:text" json:"-"`

	Bucket *Bucket `gorm:"foreignKey:UserId;references:ID;onDelete:RESTRICT;onUpdate:CASCADE" json:"-"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u User) isRole(role string) bool {
	return u.Role == role
}
