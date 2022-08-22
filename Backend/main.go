package main

import (
	"ESM-backend-app/pkg/db"
	"ESM-backend-app/pkg/handlers"
	"ESM-backend-app/pkg/middleware"
	"encoding/json"
	"log"
	"net/http"

	hndlr "github.com/gorilla/handlers"
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

	router.HandleFunc("/assign-skill", middleware.CheckAuth(h.AssignSkill)).Methods(http.MethodPost)
	router.HandleFunc("/remove-skill/{employeeid}/{skillid}", middleware.CheckAuth(h.RemoveSkill)).Methods(http.MethodDelete)

	router.HandleFunc("/login", h.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/signup", h.SignupUser).Methods(http.MethodPost)
	router.HandleFunc("/verify", h.Verify).Methods(http.MethodGet)

	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})

	log.Println("API listen at port 4000 !")
	err := http.ListenAndServe(":4000", hndlr.CORS(
		hndlr.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}),
		hndlr.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		hndlr.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		log.Fatal(err)
	}
}
