package main

import (
	"go-etl/config"
	"go-etl/db"
	"go-etl/server"
	"go-etl/services"
	"go-etl/utils"
)

func main() {
	config.LoadConfig()
	db.ConnectDB()
	utils.InitLogger()
	go services.FetchData()
	server.StartServer()
}
