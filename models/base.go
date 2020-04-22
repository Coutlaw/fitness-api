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

	// Production configs as env arguments
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(dbName)
	fmt.Println(dbHost)

	if username == "" || password == "" || dbName == "" || dbHost == "" {
		e := godotenv.Load()
		if e != nil {
			fmt.Print(e)
		}

		username = os.Getenv("TEST_DB_USER")
		password = os.Getenv("TEST_DB_PASSWORD")
		dbName = os.Getenv("TEST_DB_NAME")
		dbHost = os.Getenv("TEST_DB_HOST")
	}

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	fmt.Println(dbURI)

	conn, err := sql.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	db = conn
}

// GetDB : initialization of db connection
func GetDB() *sql.DB {
	return db
}
