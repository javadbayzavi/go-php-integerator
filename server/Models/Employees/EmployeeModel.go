package Employees

import (
	Entities "go-microservices/Entities/Employees"
	"go-microservices/lib/Models"
)

type EmployeeModel struct {
	Models.Model
}

func (this EmployeeModel) List(params []string) []Entities.Employee {
	//Stub method. must return list from its provder
	return make([]Entities.Employee, 0)
}

func (this EmployeeModel) Insert() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeModel) Delete() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeModel) Update() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeModel) Find() Entities.Employee {
	//Stub method. must return list from its provder
	return Entities.Employee{}
}
