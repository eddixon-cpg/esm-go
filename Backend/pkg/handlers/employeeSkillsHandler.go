package handlers

import (
	"ESM-backend-app/pkg/helpers"
	"ESM-backend-app/pkg/models"
	"ESM-backend-app/pkg/models/in"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

func (h Handler) AssingSkill(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var input in.EmployeeSkill
	json.Unmarshal(body, &input)

	fmt.Println("Assignin skill to employee ...")
	fmt.Println(input)

	var employeeSkill models.EmployeeSkill
	employeeSkill.Employee_employee_id = input.EmployeId
	employeeSkill.Skill_skill_id = input.SkillId
	employeeSkill.LevelId = input.Level
	employeeSkill.Experience = input.Experience

	if result := h.DB.Create(&employeeSkill); result.Error != nil {
		fmt.Println("error ", result.Error)
		helpers.ApiError(w, http.StatusForbidden, result.Error.Error())
		return
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
