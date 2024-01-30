package main

import (
	"database/sql"
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
	defer func(DBConn *sql.DB) {
		err = DBConn.Close()
		if err != nil {
			log.Fatal("error from close database connection")
		}
	}(database.DBConn)
	configs.GoogleConfig()
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("error from listen port")
		return
	}
}
