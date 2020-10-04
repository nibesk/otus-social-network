package web

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"service-users/app/config"
	"service-users/app/globals"
	"service-users/app/utils"
	"strings"
)

var SessionAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r, err := validateJwt(w, r)

		if globals.NonAuthorizedOnlyRoutes[r.URL.Path] {
			next.ServeHTTP(w, r)
			return
		}

		if err != nil {
			utils.SendResponseJsonWithStatusCode(w, utils.ResponseMessage(false, err.Error()), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateJwt(w http.ResponseWriter, r *http.Request) (*http.Request, error) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return r, errors.New("Missing auth token")
	}

	// because inside should be "Bearer JWT"
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return r, errors.New("Invalid/Malformed auth token")
	}

	tk := &utils.Token{}
	token, err := jwt.ParseWithClaims(splitted[1], tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.Server.SessionKey), nil
	})

	if err != nil {
		return r, errors.New("Malformed authentication token")
	}

	if !token.Valid {
		return r, errors.New("Token is not valid")
	}

	log.Printf("User is logged in %", tk.UserId) //Useful for monitoring

	// add userId to current request context
	ctx := context.WithValue(r.Context(), globals.AuthUserIdKey, tk.UserId)
	r = r.WithContext(ctx)

	return r, nil
}
