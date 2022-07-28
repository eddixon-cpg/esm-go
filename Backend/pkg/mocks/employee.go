package mocks

import (
	"ESM-backend-app/pkg/models"
	"time"
)

var Employees = []models.Employee{
	{
		EmployeeId:    1,
		Name:          "Martivs Graz",
		JoiningData:   time.Date(2022, 7, 17, 0, 0, 0, 0, time.Local),
		DesignationId: 6,
		Email:         "margra@mail.com",
	},
	{
		EmployeeId:    2,
		Name:          "Marcia Moretto",
		JoiningData:   time.Date(2022, 4, 27, 0, 0, 0, 0, time.Local),
		DesignationId: 6,
		Email:         "marmor@mail.com",
	},
}
