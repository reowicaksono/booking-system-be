package entity

import "time"

type Admin struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	PhoneNumber string `gorm:"type:varchar(15)" json:"phone_number"`
	Address     string `gorm:"type:varchar(255)" json:"address"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (u Admin) TableName() string {
	return "admins"
}
