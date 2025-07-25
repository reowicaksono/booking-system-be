package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string         `gorm:"type:varchar(50);unique" json:"username"`
	Name        string         `gorm:"type:varchar(50)" json:"name"`
	Email       string         `gorm:"type:varchar(50);unique" json:"email"`
	PhoneNumber string         `gorm:"type:varchar(15)" json:"phone_number"`
	Password    string         `gorm:"type:varchar(255);default:null"`
	Role        string         `gorm:"type:varchar(10)" json:"role"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Bucket    *Bucket    `gorm:"foreignKey:UserId;references:ID;onDelete:RESTRICT;onUpdate:CASCADE" json:"-"`
	Profile   *Profile   `gorm:"foreignKey:UserId;references:ID;onDelete:RESTRICT;onUpdate:CASCADE" json:"-"`
	AdminUser *AdminUser `gorm:"foreignKey:UserId;references:ID;onDelete:RESTRICT;onUpdate:CASCADE" json:"-"`
}

func (u User) isRole(role string) bool {
	return u.Role == role
}

type Profile struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId      uint   `gorm:"not null" json:"user_id"`
	AdminId     uint   `gorm:"not null" json:"admin_id"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	PhoneNumber string `gorm:"type:varchar(15)" json:"phone_number"`
	Address     string `gorm:"type:varchar(255)" json:"address"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type AdminUser struct {
	UserId  uint `gorm:"not null" json:"user_id"`
	AdminId uint `gorm:"not null" json:"admin_id"`

	Admin *Admin `gorm:"foreignKey:AdminId;references:ID;onDelete:RESTRICT;onUpdate:CASCADE" json:"-"`
}

func (u User) TableName() string {
	return "users"
}

func (p Profile) TableName() string {
	return "user_profiles"
}

func (au AdminUser) TableName() string {
	return "admin_users"
}
