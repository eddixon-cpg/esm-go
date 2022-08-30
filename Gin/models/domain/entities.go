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
	//Skills        []Skill `gorm:"many2many:employee_skills;"`
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
