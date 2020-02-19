package models
//
//import (
//	u "fitness-api/utils"
//	"fmt"
//	"github.com/jinzhu/gorm"
//)
//
//type Workout struct {
//	gorm.Model
//	WorkoutName string `json:"name"`
//	DurationWeeks  uint `json:"duration"`
//}
//
//type WorkoutAssignment struct {
//	UserId uint
//	WorkoutId uint
//}
//
//
//// create Validation
//func (workout *Workout) Validate() (string, bool) {
//
//	if workout.WorkoutName == "" {
//		return "WorkoutName should be included in request", false
//	}
//
//	if workout.DurationWeeks == 0 {
//		return "Duration should be included, to represent the number of weeks", false
//	}
//
//	//All the required parameters are present
//	return "success", true
//}
//
//// create Validation
//func (workoutAss *WorkoutAssignment) ValidateAssignment() (string, bool) {
//
//	if workoutAss.UserId == 0 {
//		return "Destination user should be included in the request", false
//	}
//
//	if workoutAss.WorkoutId == 0 {
//		return "Workout Id of assignment should be included in the request", false
//	}
//
//	//All the required parameters are present
//	return "success", true
//}
//
//func (workout *Workout) Create() map[string]interface{} {
//
//	if resp, ok := workout.Validate(); !ok {
//		return u.Message(false, resp)
//	}
//
//	GetDB().Create(workout)
//
//	resp := u.Message(true, "success")
//	resp["workout"] = workout
//	return resp
//}
//
//func GetWorkoutById(workoutId uint) *Workout {
//
//	workout := &Workout{}
//	err := GetDB().Table("workouts").Where("id = ?", workoutId).First(workout).Error
//	if err != nil {
//		return nil
//	}
//	return workout
//}
//
//func DeleteWorkoutById(workoutId uint) *Workout {
//
//	workout := &Workout{}
//	err := GetDB().Table("workouts").Where("id = ?", workoutId).Delete(workout).Error
//	if err != nil {
//		return nil
//	}
//	return workout
//}
//
//func GetUsersCurrentWorkouts(userId uint) []*Workout {
//
//	workouts := make([]*Workout, 0)
//	err := GetDB().Table("workouts").Where("userId = ?", userId).Find(&workouts).Error
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	return workouts
//}
//
//func (workoutAss *WorkoutAssignment) AssignWorkoutToUser() map[string]interface{} {
//
//	if resp, ok := workoutAss.ValidateAssignment(); !ok {
//		return u.Message(false, resp)
//	}
//
//	GetDB().Table("users").Where("userId = ?", workoutAss.UserId).Update("workout_id", workoutAss.WorkoutId)
//
//	resp := u.Message(true, "success")
//	resp["Assignment"] = workoutAss
//	return resp
//}