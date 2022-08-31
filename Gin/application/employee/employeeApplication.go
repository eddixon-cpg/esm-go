package employee

import (
	"errors"
	"esm-backend/models/domain"
	"esm-backend/models/in"
	"esm-backend/models/out"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func AddEmployee(employeeIn in.EmployeeInput, db *gorm.DB) (err error) {
	var employee domain.Employee
	email := strings.ToLower(employeeIn.Email)

	db.Where("email = ?", email).First(&employee)
	if employee.EmployeeId != 0 {
		return errors.New("already exist a employee with that email")
	}

	//mapping input to entity
	employee.EmployeeId = employeeIn.EmployeeId
	employee.Name = employeeIn.Name
	employee.LastName = employeeIn.LastName
	employee.JoiningDate = employeeIn.JoiningDate
	employee.DesignationId = employeeIn.DesignationId
	employee.Email = email

	if result := db.Create(&employee); result.Error != nil {
		//fmt.Println("error ", result.Error)
		//helpers.ApiError(w, http.StatusForbidden, result.Error.Error())
		return result.Error
	}

	return nil
}

func GetAllEmployees(db *gorm.DB) (employees []out.EmployeeOut, err error) {
	fmt.Println("GetAllEmployees APP")
	var _error error

	var employeesOut []out.EmployeeOut

	scanResult := db.Table("employees E").
		Select("E.employee_id, E.name, E.last_name, E.joining_date, E.designation_id, E.email, D.name as Designation").
		Joins("left join designations D on E.designation_id = D.designation_id").
		Scan(&employeesOut)

	if scanResult.Error != nil {
		_error = scanResult.Error
		return employeesOut, _error
	}

	return employeesOut, _error
}
