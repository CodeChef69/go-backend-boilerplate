package services

import (
	"go-simple-backend/internal/models"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}
