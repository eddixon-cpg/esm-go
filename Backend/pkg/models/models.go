package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeId int    `json:"employeeId" gorm:"primaryKey"`
	Name       string `json:"name"`
	//JoiningData   time.Time `json:"joiningData"`
	DesignationId int    `json:"designationId"`
	Email         string `json:"email"`
}

type Skill struct {
	gorm.Model
	SkillId int    `json:"skillId" gorm:"primaryKey"`
	Name    string `json:"nameId"`
}
