package di

import (
	"api-echo-template/config"
	"api-echo-template/controllers"
	"api-echo-template/graph"
	"api-echo-template/libs/logger"
	"api-echo-template/repositories"
	"api-echo-template/services"
	"sync"

	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

// GetContainer :
func GetContainer() *dig.Container {
	once.Do(func() {
		container = buildContainer()
	})
	return container
}

// BuildContainer :
func buildContainer() *dig.Container {
	container := dig.New()
	handlerContainerErrors(
		container.Provide(config.NewConfiguration),
		container.Provide(repositories.NewTodoRepository),
		container.Provide(services.NewTodoService),
		container.Provide(graph.NewResolver),
		container.Provide(controllers.NewHealthController))

	return container
}
func handlerContainerErrors(errors ...error) {
	for _, err := range errors {
		if err != nil {
			logger.ConditionalFatal("container", "buildContainer ", err)
		}
	}
}
