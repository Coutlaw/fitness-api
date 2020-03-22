package models

import (
	u "fitness-api/utils"
	"fmt"
)

// ProgramAssignment : mapping between programs and the user that created/is using them
type ProgramAssignment struct {
	UserID    uint `json:"user_id"`
	ProgramID uint `json:"program_id"`
}

// ValidateAssignment : make sure both the user and program exist before assignment
func (programAss *ProgramAssignment) ValidateAssignment() (string, bool) {

	if programAss.UserID == 0 {
		return "Destination user should be included in the route", false
	}

	if programAss.ProgramID == 0 {
		return "Program Id of assignment should be included in the request", false
	}

	//All the required parameters are present
	return "success", true
}

// AssignProgramToUser : handles assignment of a program to a user
func (programAss *ProgramAssignment) AssignProgramToUser() map[string]interface{} {

	if resp, ok := programAss.ValidateAssignment(); !ok {
		return u.Message(false, resp)
	}

	// get the base program that will be assigned
	baseProgram := Program{}

	err := db.QueryRow("SELECT * from base_programs WHERE base_program_id=$1", programAss.ProgramID).
		Scan(&baseProgram.ProgramID, &baseProgram.ProgramName, &baseProgram.ProgramCreator, &baseProgram.NumWeeks, &baseProgram.ProgramData)

	if err != nil {
		return u.Message(false, "program does not exist to assign")
	}

	// store the program in the programs table (to allow specific changes and comments from the user)
	_, err = db.
		Exec(
			"INSERT into programs (program_id, program_name, program_creator, number_of_weeks, program_data) VALUES ($1, $2, $3, $4, $5)",
			baseProgram.ProgramID,
			baseProgram.ProgramName,
			programAss.UserID,
			baseProgram.NumWeeks,
			baseProgram.ProgramData)

	if err != nil {
		return u.Message(false, "Error, unable to replicate base program")
	}

	// assign the program to the user
	_, err = db.Query("UPDATE users SET program=$1 WHERE user_id=$2", programAss.ProgramID, programAss.UserID)

	if err != nil {
		fmt.Println("err: ", err)
		return u.Message(false, "Failed to assign program, database error.")
	}

	resp := u.Message(true, "success")
	resp["Assignment"] = programAss
	return resp
}
