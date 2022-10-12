package api

import (
	"net/http"

	"github.com/google/uuid"
)


var users = map[string]string {"user1" : "password1", "user2" : "password2"}


func requestIdHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := r.Header.Get("X-Request-Id")

		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}

		w.Header().Set("X-Request-Id", requestID)
		next.ServeHTTP(w, r)

	})

}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		user := r.Header.Get("Authoruzation")

		if users[user] == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w,r)
	})

}