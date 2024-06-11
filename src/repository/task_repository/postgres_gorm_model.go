package task_repository

import (
	"gorm.io/gorm"
	"time"
)

type taskCollectionSchema struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;"`
	CreatedAt time.Time      `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time      `gorm:"not null;autoCreateTime;"`
	DeleteAt  gorm.DeletedAt `gorm:"index;"`
	IsDeleted bool           `gorm:"not null;default:false;"`

	Title       string
	Description string
	DueDate     time.Time
	Status      string
	UserID      uint64
}

func (t taskCollectionSchema) TableName() string {
	return "tasks"
}
