package controllers

import (
	"encoding/json"
	"fitness-api/models"
	u "fitness-api/utils"
	"net/http"
	"time"
)

// CreateUser : creates a user
var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	resp, sessionToken := user.Create() //Create user

	// Cache that cookie son
	http.SetCookie(w, &http.Cookie{
		Path:    "/api",
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	u.Respond(w, resp)
}

// Authenticate : provides a new JWT for the user
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	resp, sessionToken := models.Login(user.Email, user.Password)

	// Cache that cookie son
	http.SetCookie(w, &http.Cookie{
		Path:    "/api",
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	u.Respond(w, resp)
}

// UpdateUser : updates a user
var UpdateUser = func(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user) //decode the request body into struct and failed if any error occur
	// check if request was empty
	if err != nil || user == (models.User{}) {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	resp := user.Update() //Create user

	u.Respond(w, resp)
}
