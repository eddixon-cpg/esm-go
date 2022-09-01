package domain

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeId    int       `gorm:"primaryKey;autoIncrement:true"`
	Name          string    `gorm:"type:varchar(30)"`
	LastName      string    `gorm:"type:varchar(30)"`
	JoiningDate   time.Time ``
	DesignationId int       `gorm:"foreignKey"`
	Email         string    `gorm:"unique;type:varchar(30)"`
	Designation   Designation
	Skills        []Skill `gorm:"many2many:employee_skills;"`
}

type Designation struct {
	DesignationId int    `gorm:"primaryKey;autoIncrement:true"`
	Name          string `gorm:"type:varchar(20)"`
}

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(30)"`
	LastName string `gorm:"type:varchar(30)"`
	Email    string `gorm:"unique;type:varchar(30)"`
	Username string `gorm:"unique;type:varchar(30)"`
	Password string ``
}

type Skill struct {
	SkillId int    `gorm:"primaryKey;autoIncrement:true"`
	Name    string `gorm:"type:varchar(20)"`
}

type EmployeeSkill struct {
	Employee_employee_id int `gorm:"primaryKey;joinForeignKey"`
	Skill_skill_id       int `gorm:"primaryKey;joinForeignKey"`
	LevelId              int `gorm:"foreignKey"`
	Experience           int
	Level                Level
}

type Level struct {
	LevelId int    ` gorm:"primaryKey;autoIncrement:true"`
	Name    string `gorm:"type:varchar(20)"`
	Order   int
}
