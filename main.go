package main

import (
	"fitness-api/controllers"
	"fitness-api/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
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

	router.Use(models.SessionAuthentication) //attach JWT auth middleware

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