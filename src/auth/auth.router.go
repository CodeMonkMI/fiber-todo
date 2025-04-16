package auth

import (
	"github.com/CodeMonkMI/fiber-todo/src/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func AuthHandlers(api fiber.Router) {

	db = database.GetDB()
	db.AutoMigrate(&UserModel{})

	api.Post("/login", login)
	api.Post("/register", register)

}
