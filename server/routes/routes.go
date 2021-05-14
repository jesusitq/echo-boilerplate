package routes

import (
	"api-echo-template/controllers"
	"api-echo-template/graph"
	"api-echo-template/graph/generated"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

// Route :
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

// ServiceRoutes :
type ServiceRoutes []Route

// Routes :
var Routes ServiceRoutes

const serviceName = "todo"

// RegisterRoutes :
func RegisterRoutes(controllersProvider *dig.Container) {

	controllersProvider.Invoke(func(resolver *graph.Resolver, healthController *controllers.HealthController) {
		handlerQL := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
		Routes = ServiceRoutes{

			Route{Method: "GET", Name: "HealthGet", Pattern: "/health", HandlerFunc: healthController.Get},
			Route{Method: "POST", Name: "GraphQL-GET", Pattern: fmt.Sprintf("/%s/graphql", serviceName), HandlerFunc: echo.WrapHandler(handlerQL)},
		}
	})
}
