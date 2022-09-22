package server

import (
	"fmt"
	"shrinkly/model"
	"shrinkly/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	deflen int = 8
)

func GetAllRedirects(c *fiber.Ctx) error {
	shrinklies, err := model.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find links" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(shrinklies)
}

func GetSingle(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not extract shrink id" + err.Error(),
		})
	}

	shrink, err := model.GetOneShrink(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find link" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(shrink)
}

func GetSingleFromUrl(c *fiber.Ctx) error {
	url := c.Params("url")
	redirection, err := model.FindShrinkFromUrl(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not extract url id" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(redirection)
}

func CreateShrink(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var shrink model.Shrinkly
	err := c.BodyParser(&shrink)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "JSON payload invalid" + err.Error(),
		})
	}

	if shrink.Random {
		shrink.Shrinkly = utils.RandomUrl(deflen)
	}

	err = model.CreateShrink(shrink)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error pushing new url to db" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(shrink)
}

func UpdateShrink(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var shrink model.Shrinkly
	err := c.BodyParser(&shrink)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "JSON payload invalid" + err.Error(),
		})
	}

	err = model.UpdateShrink(shrink)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating new url to db" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(shrink)
}

func DeleteShrink(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not extract shrink id" + err.Error(),
		})
	}

	err = model.DeleteShrink(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting url from database" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func RedirectUrl(c *fiber.Ctx) error {
	url := c.Params("url")
	redirection, err := model.FindShrinkFromUrl(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not find shrink from given url" + err.Error(),
		})
	}

	redirection.Clicked += 1
	err = model.UpdateShrink(redirection)
	if err != nil {
		fmt.Printf("Error updating clicked in ")
	}

	return c.Redirect(redirection.Redirect, fiber.StatusTemporaryRedirect)
}

func Setup() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/shrinkly", GetAllRedirects)
	router.Get("/shrinkly/:id", GetSingle)

	router.Post("/shrinkly", CreateShrink)

	router.Patch("/shrinkly", UpdateShrink)

	router.Delete("/shrinkly/:id", DeleteShrink)

	router.Get("/r/:url", RedirectUrl)

	router.Listen(":3001")
}
