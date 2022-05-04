package main

import (
	_ "encoding/json"
	"go-microservices/Controllers/Employees"
	_ "go-microservices/Controllers/Employees"
	Address "go-microservices/Lib/Core/Addresses"

	_ "io/ioutil"
	_ "log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create new Router
	router := mux.NewRouter()
	engine := Employees.EmployeeList{}

	// route properly to respective handlers
	router.Handle("/employees", engine.Run(make([]string, 0))).Methods("GET")

	// Create new server and assign the router
	server := http.Server{
		Addr:    Address.HOST_DEFAULT_PORT,
		Handler: router,
	}

	// Start Server on defined port/host.
	server.ListenAndServe()
}
