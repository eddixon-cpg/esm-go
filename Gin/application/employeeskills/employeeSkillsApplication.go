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
		Select("s.name as Skill, l.name as Level, es.Experience as Experience").
		Joins("inner join levels l on l.level_id = es.level_id inner join skills s on s.skill_id = es.skill_skill_id").
		Where("es.Employee_employee_id = ?", employeeid).
		Scan(&employeeSkills)

	if result.Error != nil {
		err = result.Error
		return employeeSkills, err
	}

	return employeeSkills, nil
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
