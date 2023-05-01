package main

import (
	"gin-boilerplate/config"
	"gin-boilerplate/db"
	"gin-boilerplate/routes"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	router := routes.SetupRouter()
	router.Run(config.Cfg.SERVER.Port)
}
