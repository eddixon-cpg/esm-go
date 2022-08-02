package middleware

import (
	"net/http"
	"strings"

	"ESM-backend-app/pkg/helpers"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) < 2 {
			helpers.ApiError(w, http.StatusForbidden, "Token not provided!")
			return
		}

		token := bearerToken[1]

		_, err := helpers.VerifyJwtToken(token)
		if err != nil {
			helpers.ApiError(w, http.StatusForbidden, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
