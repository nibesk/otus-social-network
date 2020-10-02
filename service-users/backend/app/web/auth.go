package web

import (
	"context"
	"fmt"
	"net/http"
	"service-users/app/globals"
	"service-users/app/storage"
	"service-users/app/utils"
)

var SessionAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// current request path
		requestPath := r.URL.Path

		// check if request does not need authentication, serve the request if it doesn't need it
		if globals.NonAuthorizedOnlyRoutes[requestPath] {
			next.ServeHTTP(w, r)
			return
		}

		session := sessionStorage.GetSession(r)

		// Check if user is authenticated
		userId, ok := session.Values[storage.SessionUserIdKey].(int)
		if !ok {
			if utils.IsJsonRequest(r) {
				utils.SendResponseJsonWithStatusCode(w, utils.ResponseMessage(false, "Unauthorized"), http.StatusUnauthorized)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}

			return
		}

		fmt.Sprintf("User %", userId) //Useful for monitoring

		// add userId to current request context
		ctx := context.WithValue(r.Context(), "userId", userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
