package main

import (
	"fitness-api/controllers"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// Package level redis cache
var cache redis.Conn

func main() {
	// Redis all the things
	initCache()

	router := mux.NewRouter()

	// ROUTES

	// Create a user (gets a jwt)
	router.HandleFunc("/api/users/new", controllers.CreateAccount).Methods("POST")

	// Get a new JWT if users is expired
	router.HandleFunc("/api/users/login", controllers.Authenticate).Methods("POST")

	// Create contacts for a user
	router.HandleFunc("/api/users/contacts", controllers.CreateContact).Methods("POST")

	// Get all contacts for a user
	router.HandleFunc("/api/users/contacts", controllers.GetContacts).Methods("GET")

	// Delete all contacts for a user
	router.HandleFunc("/api/users/contacts", controllers.DeleteContacts).Methods("DELETE")

	// Get a contact by ID that belongs to a User
	router.HandleFunc("/api/users/contacts/{contactId}", controllers.GetContactById).Methods("GET")

	// Delete a contact by ID that belongs to a User
	router.HandleFunc("/api/users/contacts/{contactId}", controllers.DeleteContactById).Methods("DELETE")

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("Listening on port: " + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

func initCache() {
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
	cache = conn
}
