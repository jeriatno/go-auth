package services

import (
	"go-auth/app/models"
	"go-auth/config"
	"go-auth/utils"
)

func RegisterUser(username, password string) error {
	hashedPassword, _ := utils.HashPassword(password)
	user := models.User{Username: username, Password: hashedPassword}
	return config.DB.Create(&user).Error
}

func LoginUser(username, password string) (string, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", err
	}

	return utils.GenerateToken(username)
}
