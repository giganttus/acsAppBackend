package handlers

import (
	"acsApp/services"
)

type Handlers struct {
	Services services.IServices
}

func NewHandlers(services services.IServices) *Handlers {
	return &Handlers{
		Services: services,
	}
}
