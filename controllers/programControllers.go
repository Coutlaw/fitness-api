package controllers

import (
	"encoding/json"
	"fitness-api/models"
	u "fitness-api/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateProgram = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole) //Grab the id of the user that send the request
	Program := &models.Program{}

	err := json.NewDecoder(r.Body).Decode(Program)
	if err != nil {
		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
		return
	}

	// Only admins can create things
	if tkRole.Role != "admin"{
		http.Error(w, "Only admins can create Programs", http.StatusUnauthorized)
		return
	}

	resp := Program.Create()
	if resp["success"].(bool) != true {
		http.Error(w, resp["message"].(string), http.StatusBadRequest)
		return
	}
	u.Respond(w, resp)
}

var GetUsersCurrentProgram = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole)
	data := models.GetUsersCurrentProgram(tkRole.UserId)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetProgramById = func(w http.ResponseWriter, r *http.Request) {

	// Fetch the inline params
	vars := mux.Vars(r)
	ProgramIdParam := vars["ProgramId"]

	// Convert inline param to uint
	ProgramId, err := strconv.ParseUint(ProgramIdParam, 10, 32)
	if err != nil {
		http.Error(w, "Error with ProgramId param, could not be converted to uint", http.StatusBadRequest)
		return
	}

	data := models.GetProgramById(uint(ProgramId))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteProgramById = func(w http.ResponseWriter, r *http.Request) {

	// Fetch the inline params
	vars := mux.Vars(r)
	ProgramIdParam := vars["ProgramId"]

	// Convert inline param to uint
	ProgramId, err := strconv.ParseUint(ProgramIdParam, 10, 32)
	if err != nil {
		http.Error(w, "Error with ProgramId param, could not be converted to uint", http.StatusBadRequest)
		return
	}

	_ = models.DeleteProgramById(uint(ProgramId))
	resp := u.Message(true, "success")
	u.Respond(w, resp)
}

var AssignProgramToUser = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole)

	// Only admins can create things
	if tkRole.Role != "admin"{
		http.Error(w, "Only admins can assign Programs", http.StatusUnauthorized)
		return
	}

	ProgramAss := &models.ProgramAssignment{}

	err := json.NewDecoder(r.Body).Decode(ProgramAss)
	if err != nil {
		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
		return
	}

	data := ProgramAss.AssignProgramToUser()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UnAssignProgram = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole)

	// Only admins can create things
	if tkRole.Role != "admin"{
		http.Error(w, "Only admins can assign Programs", http.StatusUnauthorized)
		return
	}

	ProgramAss := &models.ProgramAssignment{}

	err := json.NewDecoder(r.Body).Decode(ProgramAss)
	if err != nil {
		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
		return
	}

	data := ProgramAss.UnAssignProgramToUser()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
