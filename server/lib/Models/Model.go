package Models

import (
	"go-microservices/lib/Providers/Factory"
	"go-microservices/lib/Providers/Interface"
)

type Model struct {
	Provider Interface.ProviderInterface
}

func (this Model) Init() {
	this.Provider = Factory.CreateProvider()
}
