package Controllers

import (
	"net/http"
)

type DoRun func([]string) string

type Controller struct {
}

//Create a generic http response helper for all incoming requests
func (this Controller) Run(params []string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)

		//Wrap the response into the http header as sequence of bytes
		rw.Write([]byte(this.DoRun(params)))
	}
}

func (this Controller) DoRun(params []string) string
