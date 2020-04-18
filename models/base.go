package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("TEST_DB_USER")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")
	dbHost := os.Getenv("TEST_DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := sql.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDB() *sql.DB {
	return db
}
