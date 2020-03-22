package controllers

import (
	"encoding/json"
	"fitness-api/models"
	u "fitness-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateProgram : creates a program
var CreateProgram = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole)
	program := &models.Program{}

	err := json.NewDecoder(r.Body).Decode(program)
	if err != nil {
		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
		return
	}

	// Only admins can create things
	if tkRole.Role != "trainer" {
		http.Error(w, "Only trainers can create programs", http.StatusUnauthorized)
		return
	}

	resp := program.Create(tkRole.UserId)
	if resp["success"].(bool) != true {
		http.Error(w, resp["message"].(string), http.StatusBadRequest)
		return
	}
	u.Respond(w, resp)
}

// GetProgramByID : gets a specific program by unique identifier
var GetProgramByID = func(w http.ResponseWriter, r *http.Request) {

	// Fetch the inline params
	vars := mux.Vars(r)
	programIDParam := vars["programID"]

	// Convert inline param to uint
	programID, err := strconv.ParseUint(programIDParam, 10, 32)
	if err != nil {
		http.Error(w, "Error with programID param, could not be converted to uint", http.StatusBadRequest)
		return
	}

	resp := models.GetProgramByID(uint(programID))
	u.Respond(w, resp)
}

// AssignProgramToUser : assigns an existing program to a user
var AssignProgramToUser = func(w http.ResponseWriter, r *http.Request) {

	tkRole := r.Context().Value("TkRole").(models.TkRole)

	// Only admins can create things
	if tkRole.Role != "trainer" {
		http.Error(w, "Only trainers can assign programs", http.StatusUnauthorized)
		return
	}

	// Fetch the inline params
	vars := mux.Vars(r)
	userIDParam := vars["userID"]

	// Convert inline param to uint
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		http.Error(w, "Error with userID param, could not be converted to uint", http.StatusBadRequest)
		return
	}

	programAss := &models.ProgramAssignment{}

	err = json.NewDecoder(r.Body).Decode(programAss)
	if err != nil {
		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
		return
	}

	programAss.UserID = uint(userID)

	resp := programAss.AssignProgramToUser()
	u.Respond(w, resp)
}

//
//var GetPrograms = func(w http.ResponseWriter, r *http.Request) {
//
//	tkRole := r.Context().Value("TkRole").(models.TkRole)
//	data := models.GetProgramById(tkRole.UserId)
//	resp := u.Message(true, "success")
//	resp["data"] = data
//	u.Respond(w, resp)
//}
//
//var DeleteProgramById = func(w http.ResponseWriter, r *http.Request) {
//
//	// Fetch the inline params
//	vars := mux.Vars(r)
//	programIDParam := vars["programID"]
//
//	// Convert inline param to uint
//	programID, err := strconv.ParseUint(programIDParam, 10, 32)
//	if err != nil {
//		http.Error(w, "Error with programID param, could not be converted to uint", http.StatusBadRequest)
//		return
//	}
//
//	// pull User Id from context
//	//tkRole := r.Context().Value("TkRole").(models.TkRole)
//
//	_ = models.DeleteProgramById(uint(programID))
//	resp := u.Message(true, "success")
//	u.Respond(w, resp)
//}
//
//var DeletePrograms = func(w http.ResponseWriter, r *http.Request) {
//
//	tkRole := r.Context().Value("TkRole").(models.TkRole)
//	// TODO these are wrong
//	_ = models.DeleteProgramById(tkRole.UserId)
//	resp := u.Message(true, "success")
//	u.Respond(w, resp)
//}
//

//var UnAssignProgram = func(w http.ResponseWriter, r *http.Request) {
//
//	tkRole := r.Context().Value("TkRole").(models.TkRole)
//
//	// Only admins can create things
//	if tkRole.Role != "admin"{
//		http.Error(w, "Only admins can assign programs", http.StatusUnauthorized)
//		return
//	}
//
//	data := models.GetProgramById(tkRole.UserId)
//	resp := u.Message(true, "success")
//	resp["data"] = data
//	u.Respond(w, resp)
//}
