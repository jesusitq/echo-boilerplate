package graph

import "api-echo-template/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
// NewResolver :
func NewResolver(service *services.TodoService) *Resolver {
	return &Resolver{Service: service}
}

// Resolver :
type Resolver struct {
	Service services.ITodoService
}
