package controllers

import (
	"encoding/json"
	"fitness-api/models"
	u "fitness-api/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
		return
	}

	contact.UserId = tkRole.UserId

	// Only admins can create things
	if tkRole.Role != "admin"{
		http.Error(w, "Only admins can create contacts", http.StatusUnauthorized)
		return
	}

	resp := contact.Create()
	if resp["success"].(bool) != true {
		http.Error(w, resp["message"].(string), http.StatusBadRequest)
		return
	}
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole)
	data := models.GetContacts(tkRole.UserId)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetContactById = func(w http.ResponseWriter, r *http.Request) {

	// Fetch the inline params
	vars := mux.Vars(r)
	contactIdParam := vars["contactId"]

	// Convert inline param to uint
	contactId, err := strconv.ParseUint(contactIdParam, 10, 32)
	if err != nil {
		http.Error(w, "Error with contactId param, could not be converted to uint", http.StatusBadRequest)
		return
	}

	// pull User Id from context
	tkRole := r.Context().Value("TkRole").(models.TkRole)

	data := models.GetContact(uint(contactId), tkRole.UserId)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteContactById = func(w http.ResponseWriter, r *http.Request) {

	// Fetch the inline params
	vars := mux.Vars(r)
	contactIdParam := vars["contactId"]

	// Convert inline param to uint
	contactId, err := strconv.ParseUint(contactIdParam, 10, 32)
	if err != nil {
		http.Error(w, "Error with contactId param, could not be converted to uint", http.StatusBadRequest)
		return
	}

	// pull User Id from context
	tkRole := r.Context().Value("TkRole").(models.TkRole)

	_ = models.DeleteContact(uint(contactId), tkRole.UserId)
	resp := u.Message(true, "success")
	u.Respond(w, resp)
}

var DeleteContacts = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole)
	_ = models.DeleteContacts(tkRole.UserId)
	resp := u.Message(true, "success")
	u.Respond(w, resp)
}
