package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeId    int
	Name          string
	JoiningData   time.Time
	DesignationId int
	email         string
}
