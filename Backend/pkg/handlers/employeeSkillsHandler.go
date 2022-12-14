package handlers

import (
	"ESM-backend-app/pkg/helpers"
	"ESM-backend-app/pkg/models"
	"ESM-backend-app/pkg/models/in"
	"ESM-backend-app/pkg/models/out"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) AssignSkill(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var input in.EmployeeSkill
	json.Unmarshal(body, &input)

	fmt.Println("Assigning skill to employee ...")
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

func (h Handler) RemoveSkill(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	employeeid, _ := strconv.Atoi(vars["employeeid"])
	skillid, _ := strconv.Atoi(vars["skillid"])

	fmt.Println("deleting skill from employee ... ", employeeid, skillid)
	var employeeSkill models.EmployeeSkill
	h.DB.Where("employee_employee_id = ? AND skill_skill_id = ?", employeeid, skillid).First(&employeeSkill)

	message := ""
	if employeeSkill.Employee_employee_id == 0 || employeeSkill.Skill_skill_id == 0 {
		message = "Record not found"
	} else {
		message = "Deleted"
	}
	//h.DB.Where("employee_employee_id = ? AND skill_skill_id = ?", employeeid, skillid).Delete(&models.EmployeeSkill{}).Error
	err := h.DB.Delete(&models.EmployeeSkill{}, "employee_employee_id = ? AND skill_skill_id = ?", employeeid, skillid).Error
	if err != nil {
		fmt.Println(err)
		helpers.ApiError(w, http.StatusForbidden, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

func (h Handler) GetEmpployeeSkills(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeid, _ := strconv.Atoi(vars["employeeid"])
	var employeeSkills []out.EmployeSkillsOut

	h.DB.Table("employee_skills es").
		Select("s.name as Skill, l.name as Level, es.Experience as Experience").
		Joins("inner join levels l on l.level_id = es.level_id inner join skills s on s.skill_id = es.skill_skill_id").
		Where("es.Employee_employee_id = ?", employeeid).
		Scan(&employeeSkills)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employeeSkills)
}

func (h Handler) SkillLevel(w http.ResponseWriter, r *http.Request) {
	var levels []models.Level

	h.DB.Find(&levels)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(levels)
}
