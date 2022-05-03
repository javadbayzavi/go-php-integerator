package Interface

import (
	"net/http"
)

type ControllerInterface interface {
	Run(params []string) http.HandlerFunc
}
