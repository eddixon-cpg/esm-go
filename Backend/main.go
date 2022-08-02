package main

import (
	"ESM-backend-app/pkg/db"
	"ESM-backend-app/pkg/handlers"
	"ESM-backend-app/pkg/middleware"
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

	router.HandleFunc("/employee", middleware.CheckAuth(h.GetAllEmployees)).Methods(http.MethodGet)
	router.HandleFunc("/employee/{id}", middleware.CheckAuth(h.GetEmployee)).Methods(http.MethodGet)
	router.HandleFunc("/employee", middleware.CheckAuth(h.AddEmployee)).Methods(http.MethodPost)
	router.HandleFunc("/employee/{id}", middleware.CheckAuth(h.UpdateEmployee)).Methods(http.MethodPut)
	router.HandleFunc("/employee/{id}", middleware.CheckAuth(h.DeleteEmployee)).Methods(http.MethodDelete)

	router.HandleFunc("/skill", middleware.CheckAuth(h.GetAllSkills)).Methods(http.MethodGet)
	router.HandleFunc("/skill/{id}", middleware.CheckAuth(h.GetSkill)).Methods(http.MethodGet)
	router.HandleFunc("/skill", middleware.CheckAuth(h.AddSkill)).Methods(http.MethodPost)
	router.HandleFunc("/skill/{id}", middleware.CheckAuth(h.DeleteSkill)).Methods(http.MethodDelete)

	router.HandleFunc("/login", h.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/signup", h.SignupUser).Methods(http.MethodPost)

	http.ListenAndServe(":4000", router)
	log.Println("API is running @ port 4000 !")
}
