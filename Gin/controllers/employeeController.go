package controllers

import (
	app "esm-backend/application/employee"

	"esm-backend/models/in"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h DbHandler) GetAllEmployees(c *gin.Context) {

	result, err := app.GetAllEmployees(h.DB)
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

func (h DbHandler) GetEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := app.GetEmployee(id, h.DB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if employee.EmployeeId == 0 {
		c.JSON(http.StatusNoContent, employee)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h DbHandler) AddEmployee(c *gin.Context) {
	var employee in.EmployeeInput
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := app.AddEmployee(employee, h.DB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Employee created")
}

func (h DbHandler) UpdateEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var employee in.EmployeeInput
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.UpdateEmployee(id, employee, h.DB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Employee updated")
}

func (h DbHandler) DeleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := app.DeleteEmployee(id, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Employee deleted")
}

func (h DbHandler) GetAllDesignations(c *gin.Context) {

	result, err := app.GetAllDesignations(h.DB)
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
