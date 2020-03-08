package models

import (
	"database/sql"
	u "fitness-api/utils"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

/*
JWT claims struct
*/
//a struct to rep user user
type User struct {
	UserId   uint   `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	// This is a nullable field
	Program sql.NullInt64 `json:"program_id"`
}

// `Token` belongs to `User`, `UserID` is the foreign key
type Token struct {
	UserId    uint
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

	var email string

	//check for errors and duplicate emails
	err := db.QueryRow("Select email FROM users WHERE email=$1", user.Email).Scan(&email)
	if err != nil && err != sql.ErrNoRows {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if email == user.Email {
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

	//// Prevent anyone but users from being created
	user.Role = "user"

	err := db.QueryRow("INSERT into users (email, password, role) VALUES ($1, $2, $3) RETURNING user_id", user.Email, user.Password, user.Role).Scan(&user.UserId)

	if user.UserId <= 0 || err != nil {
		return u.Message(false, "Failed to create user, connection error."), ""
	}

	user.Password = "" //delete password

	// Create a new random session token
	sessionToken := uuid.NewV4().String()

	_, err = db.Query("INSERT into tokens (session_tk, user_id) VALUES ($1, $2)", sessionToken, user.UserId)

	response := u.Message(true, "user has been created")
	response["user"] = user
	return response, sessionToken
}

func Login(email, password string) (map[string]interface{}, string) {

	user := User{}
	err := db.QueryRow("SELECT * from users WHERE email=$1", email).Scan(&user.UserId, &user.Email, &user.Password, &user.Role, &user.Program)

	if err != nil {
		if err == sql.ErrNoRows {
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
	sessionToken := uuid.NewV4().String()

	_, err = db.Query("UPDATE tokens SET session_tk=$1 WHERE user_id=$2", sessionToken, user.UserId)

	resp := u.Message(true, "Logged In")
	resp["user"] = user

	return resp, sessionToken
}
