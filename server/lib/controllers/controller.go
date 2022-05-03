package Controllers

import (
	"net/http"
)

type DoRun func([]string) string

type Controller struct {
	DoRun
}

//Create a generic http response helper for all incoming requests
func (this Controller) Run(params []string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		//Wrap the response into the http header
		this.DoRun(params)
	}
}
