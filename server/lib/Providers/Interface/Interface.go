package Interface

import (
	"go-microservices/Lib/Entities"
)

type ProviderInterface interface {
	Find() Entities.Entity
	List() []Entities.Entity
	Insert() bool
	Update() bool
	Delete() bool
	Connect() bool
}
