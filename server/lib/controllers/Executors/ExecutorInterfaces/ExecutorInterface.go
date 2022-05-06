/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
 */

package ExecutorInterfaces

import (
	"go-microservices/lib/Controllers/ControllerInterfaces"
	"net/http"
)

type ExecutorInterface interface {
	execute(params map[string]string) (bool, error)
	executeAsHandler(params map[string]string) (http.HandlerFunc, error)
	executeAsync(params map[string]string) (bool, error)
	initiateEngine(context ControllerInterfaces.ControllerInterface)
}
