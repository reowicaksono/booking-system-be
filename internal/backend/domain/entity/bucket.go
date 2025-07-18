package entity

import (
	"time"

	"gorm.io/gorm"
)

type Bucket struct {
	ID     uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId uint    `gorm:"not null" json:"user_id"`
	Name   string  `gorm:"type:varchar(22);not null" json:"name"`
	Amount float64 `gorm:"type:decimal(12,2);default:0" json:"amount"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	History []*HistoryBucket `gorm:"foreignKey:BucketId;references:ID" json:"history"`
}

type HistoryBucket struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	BucketId    uint      `gorm:"not null" json:"bucket_id"`
	ReferenceNo string    `gorm:"type:varchar(64);not null" json:"reference_no"`
	Type        string    `gorm:"type:varchar(10)" json:"type"`
	Amount      float64   `gorm:"type:decimal(12,2);default:0" json:"amount"`
	Before      float64   `gorm:"type:decimal(12,2);default:0" json:"before"`
	After       float64   `gorm:"type:decimal(12,2);default:0" json:"after"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
