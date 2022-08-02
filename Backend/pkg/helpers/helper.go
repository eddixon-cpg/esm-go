package helpers

import (
	"ESM-backend-app/pkg/models/out"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomError struct{}

var JWT_SECRET string = "secretkeyjwt"

func GenerateJwtToken(payload out.Payload) (string, error) {

	var mySigningKey = []byte(JWT_SECRET)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = payload.Email
	claims["userName"] = payload.Username
	claims["id"] = payload.Id

	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		_ = fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

	/*
		if JWT_SECRET = os.Getenv("JWT_SECRET"); JWT_SECRET == "" {
			log.Fatal("[ ERROR ] JWT_SECRET environment variable not provided!\n")
		}

		key := []byte(JWT_SECRET)

		expirationTime := time.Now().Add(7 * 24 * 60 * time.Minute)

		claims := &Claims{
			Id:       payload.Id,
			Username: payload.Username,
			Email:    payload.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		UnsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		SignedToken, err := UnsignedToken.SignedString(key)
		if err != nil {
			return "", err
		}

		return SignedToken, nil
	*/
}

func VerifyJwtToken(strToken string) (*out.Claims, error) {
	key := []byte(JWT_SECRET)

	claims := &out.Claims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, fmt.Errorf("invalid token signature")
		}
	}

	if !token.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func RespondWithJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func ApiError(w http.ResponseWriter, status int, message string) {
	error := make(map[string]string)

	error["Message"] = message
	error["Status"] = strconv.Itoa(status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)

}
