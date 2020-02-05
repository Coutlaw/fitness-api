package models

import (
	u "fitness-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Program struct {
	gorm.Model
	ProgramName   string `json:"name"`
	DurationWeeks uint   `json:"duration"`
}

type ProgramAssignment struct {
	UserId    uint
	ProgramId uint
}


// create Validation
func (program *Program) Validate() (string, bool) {

	if program.ProgramName == "" {
		return "ProgramName should be included in request", false
	}

	if program.DurationWeeks == 0 {
		return "Duration should be included, to represent the number of weeks", false
	}

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

func (program *Program) Create() map[string]interface{} {

	if resp, ok := program.Validate(); !ok {
		return u.Message(false, resp)
	}

	GetDB().Create(program)

	resp := u.Message(true, "success")
	resp["program"] = program
	return resp
}

func GetProgramById(programId uint) *Program {

	program := &Program{}
	err := GetDB().Table("workouts").Where("id = ?", programId).First(program).Error
	if err != nil {
		return nil
	}
	return program
}

func DeleteProgramById(programId uint) *Program {

	program := &Program{}
	err := GetDB().Table("workouts").Where("id = ?", programId).Delete(program).Error
	if err != nil {
		return nil
	}
	return program
}

func GetUsersCurrentProgram(userId uint) *Program {

	program := &Program{}
	err := GetDB().Table("workouts").Where("user_id = ?", userId).Find(&program).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return program
}

func (programAss *ProgramAssignment) AssignProgramToUser() map[string]interface{} {

	if resp, ok := programAss.ValidateAssignment(); !ok {
		return u.Message(false, resp)
	}

	GetDB().Table("users").Where("user_id = ?", programAss.UserId).Update("workout_id", programAss.ProgramId)

	resp := u.Message(true, "success")
	resp["Assignment"] = programAss
	return resp
}

func (programAss *ProgramAssignment) UnAssignProgramToUser() map[string]interface{} {

	if resp, ok := programAss.ValidateAssignment(); !ok {
		return u.Message(false, resp)
	}

	GetDB().Table("users").Where("user_id = ?", programAss.UserId).Update("workout_id", programAss.ProgramId)

	resp := u.Message(true, "success")
	resp["Assignment"] = programAss
	return resp
}