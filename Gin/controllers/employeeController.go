package controllers

import (
	app "esm-backend/application/employee"

	"esm-backend/models/in"
	"fmt"
	"net/http"

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

func (h DbHandler) AddEmployee(c *gin.Context) {
	fmt.Println("AddEmployee")
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
