package skill

import (
	"errors"
	"esm-backend/models/domain"
	"esm-backend/models/in"
	"esm-backend/models/out"
	"fmt"

	"gorm.io/gorm"
)

func GetAllSkills(db *gorm.DB) ([]out.SkillOutput, error) {
	var skills []out.SkillOutput

	if result := db.Find(&skills); result.Error != nil {
		return skills, result.Error
	}

	return skills, nil
}

func GetSkill(id int, db *gorm.DB) (out.SkillOutput, error) {
	var skill out.SkillOutput

	result := db.Where("skillId =? ", id).First(skill)

	if result.Error != nil {
		return skill, result.Error
	}

	if skill.SkillId == 0 {
		return skill, errors.New("skill not found")
	}

	return skill, nil
}

func AddSkill(skillIn in.SkillInput, db *gorm.DB) error {
	var skill domain.Skill

	skill.Name = skillIn.Name

	if result := db.Create(&skill); result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func DeleteSkill(id int, db *gorm.DB) error {
	var skill domain.Skill

	if result := db.First(&skill, id); result.Error != nil {
		return result.Error
	}

	if result := db.Delete(&skill); result.Error != nil {
		return result.Error
	}

	return nil
}
