package database

import (
	"fmt"

	"github.com/Norisaline/WebForum.git/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBconn *gorm.DB

// Подклчение к базе данных
func ConnectDB() {
	dsn := "host=localhost user=postgres password=2290805674 dbname=WebBlog port=5432 sslmode=disable TimeZone=Asia/Shanghai"

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
