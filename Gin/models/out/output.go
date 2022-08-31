package out

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type EmployeeOut struct {
	EmployeeId    int
	Name          string
	LastName      string
	JoiningDate   time.Time
	DesignationId int
	Email         string
	Designation   string
}

type PayloadOut struct {
	Username string
	Email    string
	Id       uint
}

type ClaimsOutput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	jwt.StandardClaims
}

type UserOut struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseOutput struct {
	User  UserOut
	Token string
}