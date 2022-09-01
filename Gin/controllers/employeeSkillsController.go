package controllers

import (
	"esm-backend/application/employeeskills"

	"esm-backend/models/in"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h DbHandler) AssignSkill(c *gin.Context) {
	var input in.EmployeeSkillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := employeeskills.AssignSkill(input, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Skill assigned to employee")
}

func (h DbHandler) RemoveSkill(c *gin.Context) {
	employeeid, _ := strconv.Atoi(c.Param("employeeid"))
	skillid, _ := strconv.Atoi(c.Param("skillid"))
	err := employeeskills.RemoveSkill(employeeid, skillid, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Skill removed from employee")
}

func (h DbHandler) GetEmployeeSkills(c *gin.Context) {
	employeeid, _ := strconv.Atoi(c.Param("employeeid"))

	result, err := employeeskills.GetEmployeeSkills(employeeid, h.DB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h DbHandler) SkillLevel(c *gin.Context) {

	var result, err = employeeskills.SkillLevel(h.DB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}
