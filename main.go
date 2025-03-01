package main

import (
	"go-auth/app/models"
	"go-auth/bootstrap"
	"go-auth/config"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	bootstrap.StartServer()
}
