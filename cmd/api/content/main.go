package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string `db:"DB_USER" env:"DB_USER"`
	DBPassword string `db:"DB_PASSWORD" env:"DB_PASSWORD"`
	DBHost     string `db:"DB_HOST" env:"DB_HOST"`
	DBPort     string `db:"DB_PORT" env:"DB_PORT"`
	DBName     string `db:"DB_NAME" env:"DB_NAME"`
}

var db *sqlx.DB

func main() {
	var err error
	// Load .env file if it exists
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Note: .env file not found, proceeding with environment variables, %s", err.Error())
	}
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	cfg.DBName = os.Getenv("DB_NAME")

	// Get a database handle.

	db, err = sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

}
