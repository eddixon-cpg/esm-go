package in

import (
	"time"
)

type EmployeeInput struct {
	EmployeeId    int       `json:"employeeId"`
	Name          string    `json:"name"`
	LastName      string    `json:"lastName"`
	JoiningDate   time.Time `json:"joiningDate"`
	DesignationId int       `json:"designationId"`
	Email         string    `json:"email"`
	//Designation   Designation
	//Skills        []Skill `gorm:"many2many:employee_skills;"`
}

type CredentialsInput struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserInput struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Username string `json:"userName"`
	Password string `json:"password"`
}

type SkillInput struct {
	SkillId int    `json:"skillId"`
	Name    string `json:"name"`
}

type EmployeeSkillInput struct {
	EmployeId  int `json:"employeId"`
	SkillId    int `json:"skillId"`
	Level      int `json:"level"`
	Experience int `json:"experience"`
}
