package controller

import (
	"log"

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

	record := new(model.Blog)

	if err := c.BodyParser(&record); err != nil {
		log.Println("Ошибка озора запроса")
		context["statuText"] = ""
		context["msg"] = "Что то пошло не так"
	}
	result := database.DBconn.Create(record)

	if result.Error != nil {
		log.Println("Ошибка при сохранение данных")
		context["statuText"] = ""
		context["msg"] = "Что то пошло не так"
	}

	context["msg"] = "Запись сохраненна"
	context["data"] = record
	c.Status(201)
	return c.JSON(context)
}

// Обновить блог в базе данных
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Update a  blog ",
	}

	//http://localhost:8080/1
	id := c.Params("id")

	var record model.Blog

	database.DBconn.First(&record, id)

	if record.ID == 0 {
		log.Println("Запись не найдена")

		context["statusText"] = ""
		context["msg"] = "Запись не найдена"
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Ошибка в Записи")
	}

	result := database.DBconn.Save(record)

	if result.Error != nil {
		log.Println("Ошибка в сохранение данных")
	}

	context["msg"] = "Запись успешно обновлнена"

	context["data"] = record

	c.Status(200)
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
