package handlers

import (
	"ESM-backend-app/pkg/models"
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
	log.Println(employees)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func (h Handler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var employee models.Employee

	if result := h.DB.First(&employee, id); result.Error != nil {
		fmt.Println(result.Error)
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

	if result := h.DB.Create(&employee); result.Error != nil {
		fmt.Println(result.Error)
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
	employee.JoiningData = updatedEmployee.JoiningData
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
