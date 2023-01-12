package main

import (
	"fmt"
	"shortfy/controllers"
	"shortfy/database"
	"shortfy/models"
	"shortfy/settings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// https://medium.com/@gophertuts/vendor-directory-in-go-723de6cab46a

func init() {
	settings.Config()
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/uris", controllers.List)
	app.Listen(":3000")
}

func InitExample() {
	db, err := database.GetDb()
	if err != nil {
		panic("Database fail")
	}

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.ShortURI{})

	user := models.User{
		Username: "trts",
		FullName: "Thiago Rabelo Torres Sales",
		Email:    "thiagorabelo1@gmail.com",
	}
	user.SetPassword("sadan")

	result := db.Create(&user)
	fmt.Println(result.Error)

	result = db.Create(&models.ShortURI{
		FullURI:     "https://www.linkedin.com/in/thiagorabelosales",
		Description: "Meu perfil no LikedIn",
		User:        user,
	})
	fmt.Println(result.Error)
}
