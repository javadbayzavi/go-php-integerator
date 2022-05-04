package Interface

import (
	"go-microservices/lib/Entities"
)

type ProviderInterface interface {
	Find() Entities.Entity
	List() []Entities.Entity
	Insert() bool
	Update() bool
	Delete() bool
	Connect() bool
}
