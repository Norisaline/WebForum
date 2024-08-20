package database

import (
	"fmt"
	"os"

	"github.com/Norisaline/WebForum.git/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBconn *gorm.DB

// Подклчение к базе данных
func ConnectDB() {

	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbname := os.Getenv("db_name")

	dsn := "host=localhost" + user + ":" + password + ":" + dbname + "port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Подключение к БД => Error")
	}
	fmt.Println("Подключение к БД => ОК")

	//Автоматическая миграция в базу данных
	db.AutoMigrate(new(model.Blog))

	DBconn = db

}
