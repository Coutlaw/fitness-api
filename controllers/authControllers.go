package controllers

import (
	"encoding/json"
	"fitness-api/models"
	u "fitness-api/utils"
	"net/http"
	"time"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	resp, sessionToken := user.Create() //Create account

	// Cache that cookie son
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	u.Respond(w, resp)
}

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
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	u.Respond(w, resp)
}
