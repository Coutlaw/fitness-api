package controllers

import (
	"encoding/json"
	"fitness-api/models"
	u "fitness-api/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
