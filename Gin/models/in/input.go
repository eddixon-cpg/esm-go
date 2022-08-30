package in

import (
	"time"
)

type EmployeeIn struct {
	EmployeeId    int       `json:"employeeId"`
	Name          string    `json:"name"`
	LastName      string    `json:"lastName"`
	JoiningDate   time.Time `json:"joiningDate"`
	DesignationId int       `json:"designationId"`
	Email         string    `json:"email"`
	//Designation   Designation
	//Skills        []Skill `gorm:"many2many:employee_skills;"`

}
