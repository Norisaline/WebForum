package controller

import (
	"github.com/Norisaline/WebForum.git/database"
	"github.com/Norisaline/WebForum.git/model"
	"github.com/gofiber/fiber/v2"
)

// Лист Блог
func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"message":    "blog List",
	}

	db := database.DBconn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

// Функция для добавление блога в базу данных
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Add a blog ",
	}

	c.Status(201)
	return c.JSON(context)
}

// Обновить блог в базе данных
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Update a  blog ",
	}
	c.Status(201)
	return c.JSON(context)
}

// Удалить блог
func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Delete a blog ",
	}
	c.Status(200)
	return c.JSON(context)
}
