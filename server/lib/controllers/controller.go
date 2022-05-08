package Controllers

import (
	"go-microservices/lib/Controllers/ControllerInterfaces"
	"go-microservices/lib/Controllers/Executors"
	"net/http"
)

type Controller struct {
	ControllerInterfaces.ControllerInterface
	executor *Executors.Executor
}

//Create a generic http response helper for all incoming requests
func (this *Controller) Run(params map[string]string) string {

	//Checking for the configuration of the execution engine
	if this.executor == nil {
		//TODO: check for the context of execution engine {Web, CLI, Microservice}
		this.executor = &Executors.Executor{}
		this.executor.SetTarget(this)
	}
	this.executor.Execute(params)
	return ""
}

//Create a generic http response helper for all incoming requests
func (this *Controller) HandleMe(me ControllerInterfaces.ControllerInterface, params map[string]string) http.HandlerFunc {

	//Checking for the configuration of the execution engine
	if this.executor == nil {
		//TODO: check for the context of execution engine {Web, CLI, Microservice}
		this.executor = &Executors.Executor{}
		this.executor.SetTarget(*&me)
	}
	return this.executor.ExecuteAsHandler(params)
}
