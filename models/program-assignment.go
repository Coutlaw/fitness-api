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

	err := db.
		QueryRow("SELECT base_program_id, program_name, program_creator, number_of_weeks, program_data FROM fitness.base_programs WHERE base_program_id=$1",
			programAss.ProgramID).
		Scan(&baseProgram.ProgramID, &baseProgram.ProgramName, &baseProgram.ProgramCreator, &baseProgram.NumWeeks, &baseProgram.ProgramData)

	if err != nil {
		return u.Message(false, "program does not exist to assign")
	}

	// store the program in the programs table (to allow specific changes and comments from the user)
	err = db.
		QueryRow(
			"INSERT into fitness.programs (program_name, program_creator, number_of_weeks, program_data, base_program) VALUES ($1, $2, $3, $4, $5) RETURNING program_id",
			baseProgram.ProgramName,
			programAss.UserID,
			baseProgram.NumWeeks,
			baseProgram.ProgramData,
			baseProgram.ProgramID).
		Scan(&programAss.ProgramID)

	if err != nil {
		return u.Message(false, "Error, unable to replicate base program")
	}

	// assign the program to the user
	_, err = db.Query("UPDATE fitness.users SET program=$1 WHERE user_id=$2", programAss.ProgramID, programAss.UserID)

	if err != nil {
		fmt.Println("err: ", err)
		return u.Message(false, "Failed to assign program, database error.")
	}

	resp := u.Message(true, "success")
	resp["Assignment"] = programAss
	return resp
}

// UnAssignProgramToUser : handles assignment of a program to a user
func UnAssignProgramToUser(userID uint) map[string]interface{} {

	// query to update the users table, returning the old program_id
	updateReturningOriginalQuery := ` UPDATE fitness.users x
																		SET    program = null
																		FROM  (SELECT * FROM fitness.users WHERE user_id = $1 FOR UPDATE) y
																		WHERE  x.user_id = y.user_id
																		RETURNING y.program`
	var originalProgramID uint
	// remove assignment FK
	err := db.QueryRow(updateReturningOriginalQuery, userID).Scan(&originalProgramID)

	// delete the program from the program table
	_, err = db.
		Exec(
			"DELETE from fitness.programs where program_id=$1", originalProgramID)

	if err != nil {
		return u.Message(false, "Error, unable to un assign program")
	}

	resp := u.Message(true, "success")
	resp["Assignment"] = "null"
	return resp
}
