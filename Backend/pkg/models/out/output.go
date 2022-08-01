package out

import (
	"ESM-backend-app/pkg/models"

	"github.com/golang-jwt/jwt"
)

type Payload struct {
	Username string
	Email    string
	Id       uint
}

type ResponseOutput struct {
	User  models.User
	Token string
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	jwt.StandardClaims
}
