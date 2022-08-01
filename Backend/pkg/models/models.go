package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	//gorm.Model
	EmployeeId    int       `json:"employeeId" gorm:"primaryKey;autoIncrement:true"`
	Name          string    `json:"name"`
	JoiningData   time.Time `json:"joiningData"`
	DesignationId int       `json:"designationId"`
	Email         string    `json:"email"`
	//CreatedAt     time.Time `gorm:"autoCreateTime:true"`
	//UpdatedAt     time.Time `gorm:"autoUpdateTime:false"`
}

type Skill struct {
	//gorm.Model
	SkillId int    `json:"skillId" gorm:"primaryKey"`
	Name    string `json:"name"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}
