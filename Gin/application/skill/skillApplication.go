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
	var skillsOut []out.SkillOutput

	var skills []domain.Skill

	if result := db.Find(&skills); result.Error != nil {
		return skillsOut, result.Error
	}

	for _, s := range skills {

		skill := out.SkillOutput{SkillId: s.SkillId, Name: s.Name}
		skillsOut = append(skillsOut, skill)
	}

	return skillsOut, nil
}

func GetSkill(id int, db *gorm.DB) (out.SkillOutput, error) {
	var skillOut out.SkillOutput
	var skill domain.Skill

	result := db.Where("skillId =? ", id).First(&skill)

	if result.Error != nil {
		return skillOut, result.Error
	}

	if skillOut.SkillId == 0 {
		return skillOut, errors.New("skill not found")
	}

	skillOut.SkillId = skill.SkillId
	skillOut.Name = skill.Name

	return skillOut, nil
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
