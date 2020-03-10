package models

import (
	"database/sql"
	u "fitness-api/utils"
)

type Program struct {
	ProgramId      uint                   `json:"program_id"`
	ProgramName    string                 `json:"program_name"`
	ProgramCreator uint                   `json:"program_creator"`
	NumWeeks       map[string]interface{} `json:"number_of_weeks"`
}

type ProgramAssignment struct {
	UserId    uint
	ProgramId uint
}

// create Validation
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

// create Validation
func (programAss *ProgramAssignment) ValidateAssignment() (string, bool) {

	if programAss.UserId == 0 {
		return "Destination user should be included in the request", false
	}

	if programAss.ProgramId == 0 {
		return "Program Id of assignment should be included in the request", false
	}

	//All the required parameters are present
	return "success", true
}

func (program *Program) Create(userId uint) map[string]interface{} {

	if resp, ok := program.Validate(); !ok {
		return u.Message(false, resp)
	}

	err := db.QueryRow("INSERT into programs (program_name, program_creator)VALUES ($1, $2) RETURNING program_id", program.ProgramName, userId).Scan(&program.ProgramId)

	if program.ProgramId <= 0 || err != nil {
		return u.Message(false, "Failed to create program, connection error.")
	}
	program.ProgramCreator = userId
	resp := u.Message(true, "success")
	resp["program"] = program
	return resp
}

func GetProgramById(programId uint) map[string]interface{} {

	program := &Program{}

	err := db.QueryRow("SELECT * from programs WHERE program_id=$1", programId).Scan(&program.ProgramId, &program.ProgramName, &program.ProgramCreator)

	if err != nil {
		if err == sql.ErrNoRows {
			return u.Message(false, "Program not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	resp := u.Message(true, "Program found")
	resp["program"] = program
	return resp
}

//
//
//func DeleteProgramById(programId uint) *Program {
//
//	program := &Program{}
//	err := GetDB().Table("programs").Where("id = ?", programId).Delete(program).Error
//	if err != nil {
//		return nil
//	}
//	return program
//}
//
//func GetUsersCurrentPrograms(userId uint) []*Program {
//
//	programs := make([]*Program, 0)
//	err := GetDB().Table("programs").Where("userId = ?", userId).Find(&programs).Error
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	return programs
//}
//
//func (programAss *ProgramAssignment) AssignProgramToUser() map[string]interface{} {
//
//	if resp, ok := programAss.ValidateAssignment(); !ok {
//		return u.Message(false, resp)
//	}
//
//	GetDB().Table("users").Where("userId = ?", programAss.UserId).Update("program_id", programAss.ProgramId)
//
//	resp := u.Message(true, "success")
//	resp["Assignment"] = programAss
//	return resp
//}
