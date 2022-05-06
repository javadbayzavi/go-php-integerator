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
	Run(params map[string]string) (http.HandlerFunc, error)
	processMe([]string) (string, error)
}
