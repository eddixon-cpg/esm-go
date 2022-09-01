package controllers

import (
	"esm-backend/application/skill"
	"esm-backend/models/in"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h DbHandler) GetAllSkills(c *gin.Context) {
	result, err := skill.GetAllSkills(h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNoContent, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h DbHandler) GetSkill(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	skill, err := skill.GetSkill(id, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if skill.SkillId == 0 {
		c.JSON(http.StatusNoContent, skill)
		return
	}

	c.JSON(http.StatusOK, skill)
}

func (h DbHandler) AddSkill(c *gin.Context) {
	var skillIn in.SkillInput

	if err := c.ShouldBindJSON(&skillIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := skill.AddSkill(skillIn, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Skill created")
}

func (h DbHandler) DeleteSkill(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := skill.DeleteSkill(id, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Skill deleted")
}
