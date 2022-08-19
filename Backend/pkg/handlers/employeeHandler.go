package handlers

import (
	"ESM-backend-app/pkg/helpers"
	"ESM-backend-app/pkg/models"
	"ESM-backend-app/pkg/models/out"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee

	log.Println("Trying to get employees")

	if result := h.DB.Find(&employees); result.Error != nil {
		fmt.Println(result.Error)
	}
	//log.Println(employees)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func (h Handler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var employee out.EmployeeOut

	h.DB.Model(&models.Employee{}).
		Select("employees.employee_id, employees.name, employees.last_name, employees.joining_date, employees.designation_id, employees.email, designations.name as Designation").
		Joins("left join designations on employees.designation_id = designations.designation_id").Where("employees.employee_id = ?", id).
		Scan(&employee)

	fmt.Println(employee)

	if employee.EmployeeId == 0 {
		errorMessage := "error : employee not found"
		fmt.Println(errorMessage)
		helpers.ApiError(w, http.StatusNoContent, errorMessage)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func (h Handler) AddEmployee(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var employee models.Employee
	json.Unmarshal(body, &employee)

	fmt.Println("Adding employee ...")
	fmt.Println(employee)

	if result := h.DB.Create(&employee); result.Error != nil {
		fmt.Println("error ", result.Error)
		helpers.ApiError(w, http.StatusForbidden, result.Error.Error())
		return
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedEmployee models.Employee
	json.Unmarshal(body, &updatedEmployee)

	var employee models.Employee

	if result := h.DB.First(&employee, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	employee.Name = updatedEmployee.Name
	employee.JoiningDate = updatedEmployee.JoiningDate
	employee.DesignationId = updatedEmployee.DesignationId
	employee.Email = updatedEmployee.Email

	h.DB.Save(&employee)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var employee models.Employee

	if result := h.DB.First(&employee, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&employee)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
