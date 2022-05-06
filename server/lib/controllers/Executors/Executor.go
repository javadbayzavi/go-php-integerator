/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
 */

package Executors

import (
	"go-microservices/lib/Controllers/ControllerInterfaces"
	"net/http"
)

type Executor struct {
	target ControllerInterfaces.ControllerInterface
}

func (this Executor) execute(params map[string]string) (bool, error) {
	//TODO: exectue context gateway action
	//this.target.Run(params)
	return true, nil
}
func (this Executor) executeAsHandler(params map[string]string) (http.HandlerFunc, error) {
	//TODO: exectue context gateway action
	//this.target.Run(param)
	return nil, nil
}
func (this Executor) executeAsync(params map[string]string) (bool, error) {
	//TODO: exectue context gateway action
	//this.target.Run(params)
	return true, nil
}
func (this Executor) initiateEngine(context ControllerInterfaces.ControllerInterface) {
	this.target = context
}
