package main

import (
	"ESM-backend-app/pkg/handlers/employee"
	"ESM-backend-app/pkg/handlers/skill"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hola Data data data ")
	router := mux.NewRouter()

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
	})

	router.HandleFunc("/employee", employee.GetAllEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employee/{id}", employee.GetEmployee).Methods(http.MethodGet)
	router.HandleFunc("/employee", employee.AddEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employee/{id}", employee.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/employee/{id}", employee.DeleteEmployee).Methods(http.MethodDelete)

	router.HandleFunc("/skill", skill.GetAllSkills).Methods(http.MethodGet)
	router.HandleFunc("/skill/{id}", skill.GetSkill).Methods(http.MethodGet)
	router.HandleFunc("/skill", skill.AddSkill).Methods(http.MethodPost)
	router.HandleFunc("/skill/{id}", skill.DeleteSkill).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
