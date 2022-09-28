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

type SkillOutput struct {
	SkillId int    `json:"skillId"`
	Name    string `json:"name"`
}

type EmployeSkillsOutput struct {
	Skillid    int64  `json:"skillId" `
	Employeeid int64  `json:"employeeId"`
	Skill      string `json:"skill"`
	Level      string `json:"level"`
	Experience int    `json:"experience"`
}

type LevelOutput struct {
	LevelId int    `gorm:"primaryKey;autoIncrement:true"`
	Name    string `gorm:"type:varchar(20)"`
	Order   int
}

type DesignationOutput struct {
	DesignationId int
	Name          string
}
