package Employees

import (
	"encoding/json"
	_ "encoding/json"
	"go-microservices/Models/Employees"
	"go-microservices/lib/Controllers"
)

type EmployeeList struct {
	Controllers.Controller
}

func (this EmployeeList) DoRun(params []string) string {
	//Fetch employees from DAO through data provider
	EmployeeDAO := Employees.EmployeeModel{}

	//Initiate DAO for employee
	EmployeeDAO.Init()

	//Create a JSON encoded result
	empjson, _ := json.MarshalIndent(EmployeeDAO.List(params), "", "  ")

	//convert encoded data into a string format
	return string(empjson)
}
