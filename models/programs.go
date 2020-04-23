package models

import (
	"database/sql"
	"encoding/json"
	u "fitness-api/utils"
	"fmt"
)

// Program : base structure returned from db
type Program struct {
	ProgramID      uint            `json:"program_id"`
	ProgramName    string          `json:"program_name"`
	ProgramCreator uint            `json:"program_creator"`
	NumWeeks       int             `json:"number_of_weeks"`
	ProgramData    json.RawMessage `json:"program_data"`
}

// Week : represents the json week of a program
type Week struct {
	Ordinal        uint   `json:"ordinal"`
	TypeOfWeek     string `json:"type_of_week"`
	WeekName       string `json:"week_name"`
	NumWorkoutDays uint   `json:"number_of_workout_days"`
	Days           []Day  `json:"days"`
}

// Day : represents the json day of a program
type Day struct {
	Ordinal     uint      `json:"ordinal"`
	NumWorkouts uint      `json:"number_of_workouts"`
	Name        string    `json:"name"`
	Workouts    []Workout `json:"workouts"`
}

// Workout : represents a workout within the day of a program, there can be multiple per day
type Workout struct {
	Ordinal   uint       `json:"ordinal"`
	Durration string     `json:"durration"`
	Exercises []Exercise `json:"exercises"`
}

// Exercise : represents the exercise component of a workout
type Exercise struct {
	Ordinal   uint     `json:"ordinal"`
	Name      string   `json:"name"`
	Equipment []string `json:"equipment"`
	Weight    string   `json:"weight"`
	Reps      uint     `json:"reps"`
	Sets      uint     `json:"sets"`
	RestTime  string   `json:"rest_between_sets"`
}

// Validate : Validate the program structure from the request body
func (program *Program) Validate() (string, bool) {

	if program.ProgramName == "" {
		return "Program Name should be included in request", false
	}

	//if program.DurationWeeks == 0 {
	//	return "Duration should be included, to represent the number of weeks", false
	//}

	//All the required parameters are present
	return "success", true
}

// Create : create a program function
func (program *Program) Create(userID uint) map[string]interface{} {

	if resp, ok := program.Validate(); !ok {
		return u.Message(false, resp)
	}

	// Using QueryRow over exec because I need the id that the DB generated
	err := db.
		QueryRow(
			"INSERT into fitness.base_programs (program_name, program_creator, number_of_weeks, program_data) VALUES ($1, $2, $3, $4) RETURNING base_program_id",
			program.ProgramName,
			userID,
			program.NumWeeks,
			program.ProgramData).
		Scan(&program.ProgramID)

	if program.ProgramID <= 0 || err != nil {
		return u.Message(false, "Failed to create program, connection error.")
	}
	program.ProgramCreator = userID
	resp := u.Message(true, "success")
	resp["program"] = program
	return resp
}

// GetProgramByID : gets a program based on its unique identifier
func GetProgramByID(programID uint) map[string]interface{} {

	program := Program{}

	err := db.QueryRow("SELECT * from fitness.base_programs WHERE base_program_id=$1", programID).
		Scan(&program.ProgramID, &program.ProgramName, &program.ProgramCreator, &program.NumWeeks, &program.ProgramData)

	if err != nil {
		if err == sql.ErrNoRows {
			return u.Message(false, "Program not found")
		}
		fmt.Println("error: ", err)
		return u.Message(false, "Connection error. Please retry")
	}

	resp := u.Message(true, "Program found")
	resp["program"] = program
	return resp
}

//
//
//func DeleteProgramById(programID uint) *Program {
//
//	program := &Program{}
//	err := GetDB().Table("programs").Where("id = ?", programID).Delete(program).Error
//	if err != nil {
//		return nil
//	}
//	return program
//}
//
//func GetUsersCurrentPrograms(userID uint) []*Program {
//
//	programs := make([]*Program, 0)
//	err := GetDB().Table("programs").Where("userID = ?", userID).Find(&programs).Error
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	return programs
//}
//
