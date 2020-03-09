package main

import (
	"fitness-api/controllers"
	"fitness-api/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// ROUTES

	// Create a user (gets a jwt)
	router.HandleFunc("/api/users/new", controllers.CreateAccount).Methods("POST")

	// Get a new JWT if users is expired
	router.HandleFunc("/api/users/login", controllers.Authenticate).Methods("POST")

	//Create a program (trainers only)
	router.HandleFunc("/api/programs", controllers.CreateProgram).Methods("POST")

	// Get a workout by workoutId
	router.HandleFunc("/api/programs/{programId}", controllers.GetProgramByID).Methods("GET")

	//// Get all Workouts (needed for trainers
	//router.HandleFunc("/api/users/workouts", controllers.GetWorkouts).Methods("GET")
	//

	//
	//// Delete a workout by workoutId
	//router.HandleFunc("/api/users/workouts/{workoutId}", controllers.DeleteWorkoutById).Methods("DELETE")

	router.Use(models.SessionAuthentication) //attach JWT auth middleware

	port := os.Args[1]

	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("Listening on port: " + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	defer models.GetDB()
}
