package handlers

import (
	"ESM-backend-app/pkg/helpers"
	"ESM-backend-app/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) GetAllSkills(w http.ResponseWriter, r *http.Request) {
	var skills []models.Skill

	log.Println("Trying to get employees")

	if result := h.DB.Find(&skills); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(skills)
}

func (h Handler) GetSkill(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var skill models.Skill

	if result := h.DB.First(&skill, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(skill)
}

func (h Handler) AddSkill(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var skill models.Skill
	json.Unmarshal(body, &skill)

	fmt.Println("Adding skiLL ...")
	fmt.Println(skill)

	if result := h.DB.Create(&skill); result.Error != nil {
		fmt.Println(result.Error)
		helpers.ApiError(w, http.StatusForbidden, result.Error.Error())
		return
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h Handler) DeleteSkill(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var skill models.Skill

	if result := h.DB.First(&skill, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&skill)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
