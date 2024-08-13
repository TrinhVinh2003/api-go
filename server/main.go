package main

import (
	"project/vnexpress/api"
	"project/vnexpress/config/driver/database"
	"project/vnexpress/internal/crawls"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.ConnectDB()
	crawls.Crawl()
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}

	defer sqlDb.Close()

	api.StartServer()
}
