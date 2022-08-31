package utilities

import (
	"esm-backend/models/out"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET string = "secretkeyjwt"

func GenerateJwtToken(payload out.PayloadOut) (string, error) {

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

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyJwtToken(strToken string) (*out.ClaimsOutput, error) {
	key := []byte(JWT_SECRET)

	claims := &out.ClaimsOutput{}

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
