package employee

import (
	"ESM-backend-app/pkg/mocks"
	"ESM-backend-app/pkg/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock data
	for _, book := range mocks.Employees {
		if book.EmployeeId == id {
			// If ids are equal item  as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var employee models.Employee
	json.Unmarshal(body, &employee)

	// Append to the Book mocks
	employee.EmployeeId = rand.Intn(100)
	mocks.Employees = append(mocks.Employees, employee)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
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

	// Iterate over all the mock Books
	for index, employee := range mocks.Employees {
		if employee.EmployeeId == id {
			// Update and send response when book Id matches dynamic Id
			employee.Name = updatedEmployee.Name
			employee.JoiningData = updatedEmployee.JoiningData
			employee.DesignationId = updatedEmployee.DesignationId
			employee.Email = updatedEmployee.Email

			mocks.Employees[index] = employee

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock Books
	for index, book := range mocks.Employees {
		if book.EmployeeId == id {
			// Delete book and send response if the book Id matches dynamic Id
			mocks.Employees = append(mocks.Employees[:index], mocks.Employees[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
