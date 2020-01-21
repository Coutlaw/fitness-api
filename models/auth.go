package models

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type TkRole struct {
	UserId uint
	SessionTK string
	Role string
}

var SessionAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/users/new", "/api/users/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                               //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		sessionToken, err := r.Cookie("session_token")

		if err != nil ||  time.Now().Unix() < sessionToken.Expires.Unix() {
			http.Error(w, "invalid or expired session_token, please log in again", http.StatusForbidden)
			return
		}

		// TODO remove this later
		fmt.Println("session expiration: ", sessionToken.Expires.Unix())
		fmt.Println("session expiration: ", sessionToken)

		tk := &TkRole{}
		// search for the token in the DB
		err = GetDB().Select("session_tk, role, user_id").Table("tokens").Joins("join users on users.id=tokens.user_id").Where("session_tk = ?", sessionToken.Value).Find(tk).Error
		if err != nil {
			http.Error(w, "session token does not match any users, please log in again", http.StatusForbidden)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		// Log the user
		fmt.Println(fmt.Sprintf("UserId: %v authenticated successfully", tk.UserId))
		ctx := context.WithValue(r.Context(), "TkRole", *tk)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	});
}
