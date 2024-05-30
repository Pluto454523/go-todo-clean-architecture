package user_repository

import "gorm.io/gorm"

type userCollectionSchema struct {
	gorm.Model
	Name string
}

func (t userCollectionSchema) TableName() string {
	return "users"
}
