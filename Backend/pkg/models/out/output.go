package out

import (
	"ESM-backend-app/pkg/models"
	"time"

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

type EmployeeOut struct {
	EmployeeId    int
	Name          string
	LastName      string
	JoiningDate   time.Time
	DesignationId int
	Email         string
	Designation   string
}

type EmployeSkillsOut struct {
	Skill      string
	Level      string
	Experience int
}
