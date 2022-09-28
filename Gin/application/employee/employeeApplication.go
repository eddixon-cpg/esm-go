package employee

import (
	"errors"
	"esm-backend/models/domain"
	"esm-backend/models/in"
	"esm-backend/models/out"
	"strings"

	"gorm.io/gorm"
)

func GetAllEmployees(db *gorm.DB) (employees []out.EmployeeOut, err error) {
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

func GetEmployee(id int, db *gorm.DB) (employee out.EmployeeOut, err error) {
	result := db.Model(&domain.Employee{}).
		Select("employees.employee_id, employees.name, employees.last_name, employees.joining_date, employees.designation_id, employees.email, designations.name as Designation").
		Joins("left join designations on employees.designation_id = designations.designation_id").
		Where("employees.employee_id = ?", id).
		Scan(&employee)

	if result.Error != nil {
		return out.EmployeeOut{}, result.Error
	}

	return employee, err
}

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

func UpdateEmployee(id int, updatedEmployee in.EmployeeInput, db *gorm.DB) error {
	var employee domain.Employee

	if result := db.First(&employee, id); result.Error != nil {
		return result.Error
	}

	employee.Name = updatedEmployee.Name
	employee.LastName = updatedEmployee.LastName
	employee.JoiningDate = updatedEmployee.JoiningDate
	employee.DesignationId = updatedEmployee.DesignationId
	employee.Email = strings.ToLower(updatedEmployee.Email)

	if result := db.Save(&employee); result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteEmployee(id int, db *gorm.DB) error {
	var employee domain.Employee

	if result := db.First(&employee, id); result.Error != nil {
		return result.Error
	}

	if result := db.Delete(&employee); result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllDesignations(db *gorm.DB) ([]out.DesignationOutput, error) {
	var _error error
	var designations []domain.Designation
	var designationsOutput []out.DesignationOutput

	result := db.Find(&designations)
	if result.Error != nil {
		return make([]out.DesignationOutput, 0), result.Error
	}

	for _, designation := range designations {
		var data out.DesignationOutput
		data.DesignationId = designation.DesignationId
		data.Name = designation.Name

		designationsOutput = append(designationsOutput, data)
	}

	return designationsOutput, _error
}
