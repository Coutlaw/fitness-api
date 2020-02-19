package controllers
//
//import (
//	"encoding/json"
//	"fitness-api/models"
//	u "fitness-api/utils"
//	"github.com/gorilla/mux"
//	"net/http"
//	"strconv"
//)
//
//var CreateWorkout = func(w http.ResponseWriter, r *http.Request) {
//
//	tkRole := r.Context().Value("TkRole").(models.TkRole) //Grab the id of the user that send the request
//	workout := &models.Workout{}
//
//	err := json.NewDecoder(r.Body).Decode(workout)
//	if err != nil {
//		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
//		return
//	}
//
//	// Only admins can create things
//	if tkRole.Role != "admin"{
//		http.Error(w, "Only admins can create workouts", http.StatusUnauthorized)
//		return
//	}
//
//	resp := workout.Create()
//	if resp["success"].(bool) != true {
//		http.Error(w, resp["message"].(string), http.StatusBadRequest)
//		return
//	}
//	u.Respond(w, resp)
//}
//
//var GetWorkouts = func(w http.ResponseWriter, r *http.Request) {
//
//	tkRole := r.Context().Value("TkRole").(models.TkRole)
//	data := models.GetWorkoutById(tkRole.UserId)
//	resp := u.Message(true, "success")
//	resp["data"] = data
//	u.Respond(w, resp)
//}
//
//var GetWorkoutById = func(w http.ResponseWriter, r *http.Request) {
//
//	// Fetch the inline params
//	vars := mux.Vars(r)
//	workoutIdParam := vars["workoutId"]
//
//	// Convert inline param to uint
//	workoutId, err := strconv.ParseUint(workoutIdParam, 10, 32)
//	if err != nil {
//		http.Error(w, "Error with workoutId param, could not be converted to uint", http.StatusBadRequest)
//		return
//	}
//
//	data := models.GetWorkoutById(uint(workoutId))
//	resp := u.Message(true, "success")
//	resp["data"] = data
//	u.Respond(w, resp)
//}
//
//var DeleteWorkoutById = func(w http.ResponseWriter, r *http.Request) {
//
//	// Fetch the inline params
//	vars := mux.Vars(r)
//	workoutIdParam := vars["workoutId"]
//
//	// Convert inline param to uint
//	workoutId, err := strconv.ParseUint(workoutIdParam, 10, 32)
//	if err != nil {
//		http.Error(w, "Error with workoutId param, could not be converted to uint", http.StatusBadRequest)
//		return
//	}
//
//	// pull User Id from context
//	//tkRole := r.Context().Value("TkRole").(models.TkRole)
//
//	_ = models.DeleteWorkoutById(uint(workoutId))
//	resp := u.Message(true, "success")
//	u.Respond(w, resp)
//}
//
//var DeleteWorkouts = func(w http.ResponseWriter, r *http.Request) {
//
//	tkRole := r.Context().Value("TkRole").(models.TkRole)
//	// TODO these are wrong
//	_ = models.DeleteWorkoutById(tkRole.UserId)
//	resp := u.Message(true, "success")
//	u.Respond(w, resp)
//}
//
//var AssignWorkoutToUser = func(w http.ResponseWriter, r *http.Request) {
//
//	tkRole := r.Context().Value("TkRole").(models.TkRole)
//
//	// Only admins can create things
//	if tkRole.Role != "admin"{
//		http.Error(w, "Only admins can assign workouts", http.StatusUnauthorized)
//		return
//	}
//
//	workoutAss := &models.WorkoutAssignment{}
//
//	err := json.NewDecoder(r.Body).Decode(workoutAss)
//	if err != nil {
//		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
//		return
//	}
//
//	data := workoutAss.AssignWorkoutToUser()
//	resp := u.Message(true, "success")
//	resp["data"] = data
//	u.Respond(w, resp)
//}
//
//var UnAssignWorkout = func(w http.ResponseWriter, r *http.Request) {
//
//	tkRole := r.Context().Value("TkRole").(models.TkRole)
//
//	// Only admins can create things
//	if tkRole.Role != "admin"{
//		http.Error(w, "Only admins can assign workouts", http.StatusUnauthorized)
//		return
//	}
//
//	data := models.GetWorkoutById(tkRole.UserId)
//	resp := u.Message(true, "success")
//	resp["data"] = data
//	u.Respond(w, resp)
//}
