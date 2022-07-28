package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeId    int       `json:"employeeId"`
	Name          string    `json:"name"`
	JoiningData   time.Time `json:"joiningData"`
	DesignationId int       `json:"designationId"`
	Email         string    `json:"email"`
}
