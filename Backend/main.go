package main

import (
	"ESM-backend-app/pkg/db"
	"ESM-backend-app/pkg/handlers"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting API!")
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
	})

	router.HandleFunc("/employee", h.GetAllEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employee/{id}", h.GetEmployee).Methods(http.MethodGet)
	router.HandleFunc("/employee", h.AddEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employee/{id}", h.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/employee/{id}", h.DeleteEmployee).Methods(http.MethodDelete)

	router.HandleFunc("/skill", h.GetAllSkills).Methods(http.MethodGet)
	router.HandleFunc("/skill/{id}", h.GetSkill).Methods(http.MethodGet)
	router.HandleFunc("/skill", h.AddSkill).Methods(http.MethodPost)
	router.HandleFunc("/skill/{id}", h.DeleteSkill).Methods(http.MethodDelete)

	router.HandleFunc("/login", h.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/signup", h.SignupUser).Methods(http.MethodPost)

	http.ListenAndServe(":4000", router)
	log.Println("API is running @ port 4000 !")
}
