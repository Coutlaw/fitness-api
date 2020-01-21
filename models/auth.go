package models


import (
	"context"
	"fmt"
	"net/http"
	"time"
)

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
		fmt.Println("session expiration: ", sessionToken)
		if err != nil {
			http.Error(w, "unable to find session token in cookies, please log in again", http.StatusForbidden)
			return
		}
		fmt.Println("session expiration: ", sessionToken.Expires.Unix())
		if time.Now().Unix() < sessionToken.Expires.Unix() {
			http.Error(w, "session is expired, pleas log in again", http.StatusForbidden)
			return
		}

		tk := &Token{}
		// search for the token in the DB
		err = GetDB().Table("tokens").Where("session_tk = ?", sessionToken.Value).Find(tk).Error
		if err != nil {
			http.Error(w, "session token does not match any users, please log in again", http.StatusForbidden)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		// Log the user
		fmt.Println(fmt.Sprintf("UserId: %v authenticated successfully", tk.UserId))
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	});
}
