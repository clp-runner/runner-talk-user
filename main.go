package main

import (
	"github.com/clp-runner/runner-user/configs"
	"github.com/clp-runner/runner-user/database"
	"github.com/clp-runner/runner-user/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func main()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := routes.Router()
	database.ConnectToDB()
	configs.GoogleConfig()
	app.Listen(":3000")
}
