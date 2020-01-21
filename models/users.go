package models

import (
	u "fitness-api/utils"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

/*
JWT claims struct
*/
//a struct to rep user user
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// `Token` belongs to `User`, `UserID` is the foreign key
type Token struct {
	gorm.Model
	User User
	UserId uint
	SessionTK string
}

//var store = sessions.NewCookieStore([]byte("token_password"))

//Validate incoming user details...
func (user *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (user *User) Create() (map[string]interface{}, string) {

	if resp, ok := user.Validate(); !ok {
		return resp, ""
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error."), ""
	}

	user.Password = "" //delete password

	// Create a new random session token
	sessionToken:= uuid.NewV4().String()

	tk := &Token{
		User: *user,
		UserId: user.ID,
		SessionTK: sessionToken,
	}

	GetDB().Create(tk)

	response := u.Message(true, "user has been created")
	response["user"] = user
	return response, sessionToken
}

func Login(email, password string) (map[string]interface{}, string) {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found"), ""
		}
		return u.Message(false, "Connection error. Please retry"), ""
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again"), ""
	}
	//Worked! Logged In
	user.Password = ""

	// Create a new random session token
	sessionToken:= uuid.NewV4().String()

	tk := &Token{
		User: *user,
		UserId: user.ID,
		SessionTK: sessionToken,
	}

	err = GetDB().Table("tokens").Where("user_id = ?", user.ID).Update("session_tk", sessionToken).Error
	if err != nil {
		GetDB().Create(tk)
	}

	resp := u.Message(true, "Logged In")
	resp["user"] = user

	return resp, sessionToken
}
