package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// Переменная для подключения к базе данных
var DB *sql.DB

// Функция инициализации базы данных
func InitDB() {
	connStr := "postgres://postgres:12345678@localhost/stats-collection?sslmode=disable"
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Database connection established")
}
