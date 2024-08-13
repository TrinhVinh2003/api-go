package database

import (
	"fmt"
	"log"
	"project/vnexpress/api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := "root:534473@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(db)

	log.Println("Connection successful.")

	db.AutoMigrate(new(models.Article))
	db.AutoMigrate(new(models.User))
	DBConn = db
}
