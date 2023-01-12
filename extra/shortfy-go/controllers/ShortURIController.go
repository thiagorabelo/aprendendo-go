package controllers

import (
	"net/http"
	"shortfy/database"
	"shortfy/repository"

	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	db, err := database.GetDb()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	repo := repository.NewShortURIRepository(db)

	uris, err := repo.ListAll()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(uris)
}

func Get(c *fiber.Ctx) error {
	return nil
}

func Post(c *fiber.Ctx) error {
	return nil
}

func Patch(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
