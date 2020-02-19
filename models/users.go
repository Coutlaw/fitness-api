package models

import (
	"database/sql"
	u "fitness-api/utils"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
/*
JWT claims struct
*/
//a struct to rep user user
type User struct {
	UserId uint `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Workout  uint   `json:"workout_id"`
}

// `Token` belongs to `User`, `UserID` is the foreign key
type Token struct {
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
	//check for errors and duplicate emails
	result, err := psql.Select("email").From("users").Where(sq.Eq{"email": user.Email}).RunWith(db).Query()
	fmt.Println("result ", result)
	fmt.Println("err ", err)


	if err != nil && err != sql.ErrNoRows {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if result != nil {
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
	//user.Role = "user"

	query := psql.Insert("users").
		Columns( "email", "password", "role", "program").
		Values(user.Email, user.Password, user.Role, user.Workout).
		Suffix("RETURNING \"userid\"").
		RunWith(db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRow().Scan(user.UserId)

	if user.UserId <= 0 || err != nil{
		return u.Message(false, "Failed to create user, connection error."), ""
	}

	user.Password = "" //delete password

	// Create a new random session token
	sessionToken:= uuid.NewV4().String()

	sqls, args, err := psql.Insert("").
		Into("tokens").
		Columns("userid", "sessiontk").
		Values(user.UserId, sessionToken).
		ToSql()

	_, _ = db.Exec(sqls, args...)

	response := u.Message(true, "user has been created")
	response["user"] = user
	return response, sessionToken
}

func Login(email, password string) (map[string]interface{}, string) {

	//user := &User{}
	//err := "" //GetDB().Table("users").Where("email = ?", email).First(user).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return u.Message(false, "Email address not found"), ""
	//	}
	//	return u.Message(false, "Connection error. Please retry"), ""
	//}
	//
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
	//	return u.Message(false, "Invalid login credentials. Please try again"), ""
	//}
	////Worked! Logged In
	//user.Password = ""
	//
	//// Create a new random session token
	sessionToken:= uuid.NewV4().String()
	//
	//tk := &Token{
	//	UserId: user.UserId,
	//	SessionTK: sessionToken,
	//}
	//
	//err = nil //GetDB().Table("tokens").Where("userId = ?", user.UserId).Update("session_tk", sessionToken).Error
	//if err != nil {
	//	GetDB().Create(tk)
	//}

	resp := u.Message(true, "Logged In")
	//resp["user"] = user

	return resp, sessionToken
}
