package main

import (
	"fmt"
)

func testAppEngine() {
	//TODO: here we have to ascertain about the object creational of the whole framework
	//var engin = Employees.EmployeeList{}
	//engin.Run(make([]string, 0))
	var engine = child{}
	//engine.executor = executor{context: engine}
	engine.Run(engine, "child class")
}

type parentinterface interface {
	Run(context parentinterface, name string)
	processMe(params []string)
}
type executorInterface interface {
	execute(params []string)
}
type executor struct {
	context parentinterface
}

func (this executor) execute(params []string) {

	this.context.processMe(params)
}

type parent struct {
	executor executorInterface
}

func (this parent) Run(context parentinterface, name string) {
	this.executor = executor{context: context}
	this.executor.execute(append(make([]string, 0), name))
}

type child struct {
	parent
	classname string
}

func (this child) processMe(params []string) {
	this.classname = params[0]
	//specific strategy for this child class
	fmt.Println(this.showMyInfo())
}

func (this child) showMyInfo() string {
	return this.classname
}
