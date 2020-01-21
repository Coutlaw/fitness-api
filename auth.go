package main

import (
	"context"
	"fitness-api/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

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

		// Grab the token from the header
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			// Token is missing, return error code 403 Unauthorized
			http.Error(w, "Missing auth token", http.StatusForbidden)
			return
		}

		// The token normally comes in format `Bearer {token-body}`
		// we check if the retrieved token matched this requirement
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			http.Error(w, "Invalid/Malformed auth token, insure it comes in as Bearer {token-body}", http.StatusForbidden)
			return
		}

		//Grab the token part, what we are truly interested in
		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		//Malformed token, returns with http code 403 as usual
		if err != nil {
			http.Error(w, "Malformed authentication token, unable to correctly parse token", http.StatusForbidden)
			return
		}

		//Token is invalid, maybe not signed on this server
		if !token.Valid {
			http.Error(w, "Invalid Token, maybe not signed on this server", http.StatusForbidden)
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
