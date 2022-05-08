/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
 */

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

func (this EmployeeList) ProcessMe(params map[string]string) string {
	//Fetch employees from DAO through data provider
	EmployeeDAO := Employees.EmployeeModel{}

	//Initiate DAO for employee
	EmployeeDAO.Init()

	//Create a JSON encoded result
	empjson, _ := json.MarshalIndent(EmployeeDAO.List(params), "", "  ")

	//convert encoded data into a string format
	return string(empjson)
}
