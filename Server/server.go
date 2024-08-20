package main

import (
	"log"

	"github.com/Norisaline/WebForum.git/database"
	"github.com/Norisaline/WebForum.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// отдельная функция для подключение к базе данных
func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Ошиба в загрузке файла .env")
	}

	database.ConnectDB()
}

func main() {
	sqlDb, err := database.DBconn.DB()

	if err != nil {
		panic("Ошибка в подключение к SQL")
	}

	//закрыть подключение к базе данных после отработки функции
	defer sqlDb.Close()

	app := fiber.New()

	app.Use(logger.New())

	router.SetupRoutes(app)
	//Порт 8080 для локалхоста
	app.Listen(":8080")
}
