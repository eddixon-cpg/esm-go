package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeId    int       `json:"employeeId" gorm:"primaryKey;autoIncrement:true"`
	Name          string    `json:"name" gorm:"type:varchar(30)"`
	LastName      string    `json:"lastName" gorm:"type:varchar(30)"`
	JoiningDate   time.Time `json:"joiningDate"`
	DesignationId int       `json:"designationId"`
	Email         string    `json:"email" gorm:"type:varchar(30)"`
	Designation   Designation
	Skills        []Skill `gorm:"many2many:employee_skills;"`
}

type Designation struct {
	DesignationId int    `json:"designationId" gorm:"primaryKey;autoIncrement:true"`
	Name          string `json:"name"  gorm:"type:varchar(20)"`
}

type Skill struct {
	SkillId int    `json:"skillId" gorm:"primaryKey;autoIncrement:true"`
	Name    string `json:"name" gorm:"type:varchar(20)"`
	//Employees []*Employee `gorm:"many2many:employee_skills;"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type EmployeeSkill struct {
	Employee_employee_id int `gorm:"primaryKey;joinForeignKey"`
	Skill_skill_id       int `gorm:"primaryKey;joinForeignKey"`
	LevelId              int
	Experience           int
}

type Level struct {
	LevelId int    `json:"levelId" gorm:"primaryKey;autoIncrement:true"`
	Name    string `gorm:"type:varchar(20)"`
	Order   int
}
