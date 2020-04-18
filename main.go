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

	// USER ROUTES
	// Create a user (gets a jwt)
	router.HandleFunc("/api/users/new", controllers.CreateUser).Methods("POST")

	// Get a new JWT if users is expired
	router.HandleFunc("/api/users/login", controllers.Authenticate).Methods("POST")

	// update a users info (only specific user)
	router.HandleFunc("/api/users", controllers.UpdateUser).Methods("PUT")

	// TODO:

	// set a users account status {paid, behind, paused, cancelled}
	// delete a user

	// BASE PROGRAM ROUTES (Trainers Only)
	//Create a base level program (trainers only)
	router.HandleFunc("/api/programs", controllers.CreateProgram).Methods("POST")

	// Get a base program by its unique id
	router.HandleFunc("/api/programs/{programID}", controllers.GetProgramByID).Methods("GET")

	// TODO:
	// update base program fields (trainer)
	// delete program (trainer)

	// PROGRAM ROUTES
	// Assign a program to a user (trainer)
	router.HandleFunc("/api/users/{userID}/program", controllers.AssignProgramToUser).Methods("POST")

	// this should be reworked to a DELETE for the above route ^
	// Un-assign a program to a user (trainer)
	router.HandleFunc("/api/users/{userID}/program/unassign", controllers.UnAssignProgram).Methods("POST")

	// TODO:
	// get a users program (specifc user only)
	// get specific parts of a program (specific user only) Ex week, day, exercise
	// leave a comment about a day
	// input data about what the user completed

	//old dead code
	//router.HandleFunc("/api/users/workouts", controllers.GetWorkouts).Methods("GET")
	//router.HandleFunc("/api/users/workouts/{workoutId}", controllers.DeleteWorkoutById).Methods("DELETE")

	router.Use(models.SessionAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println("Listening on port: " + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app
	if err != nil {
		fmt.Print(err)
	}

	defer models.GetDB()
}
