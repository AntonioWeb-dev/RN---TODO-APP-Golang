package middlewares

import (
	"net/http"
	"strings"

	"api/helpers/jwt"
	"api/helpers/response"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if len(strings.Split(token, " ")) == 2 {
			token = strings.Split(token, " ")[1]
		}
		userId, err := jwt.VerifyToken(token, "secretLogin")
		if err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}
		r.Header.Set("userId", userId)
		next(w, r)
	}
}
