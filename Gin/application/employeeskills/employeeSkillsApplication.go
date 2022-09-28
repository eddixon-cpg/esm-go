package employeeskills

import (
	"esm-backend/models/domain"
	"esm-backend/models/in"
	"esm-backend/models/out"

	"gorm.io/gorm"
)

func AssignSkill(input in.EmployeeSkillInput, db *gorm.DB) (err error) {
	var employeeSkill domain.EmployeeSkill
	employeeSkill.Employee_employee_id = input.EmployeId
	employeeSkill.Skill_skill_id = input.SkillId
	employeeSkill.LevelId = input.Level
	employeeSkill.Experience = input.Experience

	if result := db.Create(&employeeSkill); result.Error != nil {
		err = result.Error
		return err
	}
	err = nil
	return err
}

func RemoveSkill(employeeid int, skillid int, db *gorm.DB) (err error) {
	if result := db.Delete(&domain.EmployeeSkill{}, "employee_employee_id = ? AND skill_skill_id = ?", employeeid, skillid); result.Error != nil {
		err = result.Error
		return err
	}
	return nil
}

func GetEmployeeSkills(employeeid int, db *gorm.DB) (employeeSkills []out.EmployeSkillsOutput, err error) {

	result := db.Table("employee_skills es").
		Select("s.skill_id AS SkillId, es.Employee_employee_id AS EmployeeId, s.name AS Skill, l.name AS Level, es.Experience AS Experience").
		Joins("INNER JOIN levels l ON l.level_id = es.level_id INNER JOIN skills s ON s.skill_id = es.skill_skill_id").
		Where("es.Employee_employee_id = ?", employeeid).
		Scan(&employeeSkills)

	if result.Error != nil {
		err = result.Error
		return employeeSkills, err
	}

	return employeeSkills, nil
}

func GetSkillsByEmployeeId(id int, db *gorm.DB) {
	var skillByEmployee SkillsByEmployeeDTO
	db.Table("employees").
		Select("skills.id as skillId, skills.skill, levels.id as levelId, levels.name as level, employee_skills.experience as experience").
		Joins("JOIN employee_skills ON employee_skills.employee_id = employees.id JOIN skills ON skills.id = employee_skills.skill_id JOIN levels ON levels.id = employee_skills.level_id").
		Where("employees.id = ?", id).
		Scan(&skillByEmployee)
}

type SkillsByEmployeeDTO struct {
	Skillid    int64  `json:"skillID" `
	Skill      string `json:"skill"`
	Levelid    int64  `json:"levelId"`
	Level      string `json:"level"`
	Experience int64  `json:"experience"`
}

func SkillLevel(db *gorm.DB) ([]out.LevelOutput, error) {
	var levels []domain.Level
	var levelsOut []out.LevelOutput

	result := db.Order("levels.Order").Find(&levels)

	if result.Error != nil {
		return make([]out.LevelOutput, 0), result.Error
	}

	for _, level := range levels {
		var levelOut out.LevelOutput

		levelOut.LevelId = level.LevelId
		levelOut.Name = level.Name
		levelOut.Order = level.Order

		levelsOut = append(levelsOut, levelOut)
	}

	return levelsOut, nil
}
