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

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("db_host"),
		os.Getenv("db_user"),
		os.Getenv("db_password"),
		os.Getenv("db_name"),
		os.Getenv("db_port"),
		os.Getenv("db_sslmodel"),
		os.Getenv("db_timezone"),
	)

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
