package main

import (
	"github.com/joho/godotenv"
	"os"
	"webservice-pattern/Config"
	"webservice-pattern/Router"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	Config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := Config.Init()
	app := Router.Init(init)

	err := app.Run(":" + port)
	if err != nil {
		return
	}
}
