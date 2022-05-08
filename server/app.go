/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
 */

package main

import (
	_ "encoding/json"
	"go-microservices/Controllers/Employees"
	Address "go-microservices/Lib/Core/Addresses"

	_ "io/ioutil"
	_ "log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Just some basic test
	//testAppEngine()

	// Create new Router
	router := mux.NewRouter()
	//TODO: we need a Abstract Factory pattern to generate all type of microservice requests
	var engin = Employees.EmployeeList{}
	var params map[string]string = map[string]string{"context": "service"}

	// route properly to respective handlers
	router.Handle("/employees", engin.HandleMe(&engin, params)).Methods("GET")

	// Create new server and assign the router
	server := http.Server{
		Addr:    Address.HOST_DEFAULT_PORT,
		Handler: router,
	}

	// Start Server on defined port/host.
	server.ListenAndServe()
}
