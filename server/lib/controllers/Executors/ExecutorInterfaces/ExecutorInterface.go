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
	Execute(params map[string]string) string
	ExecuteAsHandler(params map[string]string) http.HandlerFunc
	ExecuteAsync(params map[string]string) (bool, error)
	initiateEngine(context *ControllerInterfaces.ControllerInterface)
	SetTarget(target *ControllerInterfaces.ControllerInterface)
	GetTarget() *ControllerInterfaces.ControllerInterface
}
