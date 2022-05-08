/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
 */

package ControllerInterfaces

import (
	"net/http"
)

type ControllerInterface interface {
	//For CLI request process
	Run(params map[string]string) string

	//Handle http based requests
	HandleMe(context ControllerInterface, params map[string]string) http.HandlerFunc

	//Strategy engine handler
	ProcessMe(params map[string]string) string
}
