/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
*/

package Interface

import (
	"go-microservices/lib/Entities"
)

type ProviderInterface interface {
	Find(params map[string]string) *Entities.Entity
	List(params map[string]string) []*Entities.Entity
	Insert() bool
	Update() bool
	Delete() bool
	Connect() bool
}
