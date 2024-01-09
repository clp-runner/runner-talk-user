package database

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	DBConn *sql.DB
)

const (
	host     = "localhost"
	port     = 5432
	username = "root"
	password = "password"
	database = "postgres"
)

type DBConfig struct {
	host     string
	port     int
	username string
	password string
	database string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}
}

func (c *DBConfig) GetHost() string {
	return c.host
}

func (c *DBConfig) GetPort() int {
	return c.port
}

func (c *DBConfig) GetUsername() string {
	return c.username
}

func (c *DBConfig) GetPassword() string {
	return c.password
}

func (c *DBConfig) GetDatabase() string {
	return c.database
}

func ConnectToDB() {
	var err error

	dbConfig := NewDBConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.GetHost(), dbConfig.GetPort(), dbConfig.GetUsername(), dbConfig.GetPassword(), dbConfig.GetDatabase())
	DBConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect database")
	}
}