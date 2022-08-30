package out

import (
	"time"
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
