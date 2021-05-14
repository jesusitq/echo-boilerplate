package server

import (
	"api-echo-template/libs/logger"
	"api-echo-template/models"
	"api-echo-template/server/routes"
	"fmt"
	"net/http"

	"api-echo-template/server/di"

	"github.com/labstack/echo/v4"

	"go.uber.org/dig"
)

var server *echo.Echo

func init() {
	server = echo.New()
	server.HideBanner = true
}

// InitServer :
func InitServer() {
	container := di.GetContainer()
	setupRoutes(container)
	setupNotFoundHandler()
	startServer(container)
}

func setupRoutes(container *dig.Container) {
	routes.RegisterRoutes(container)
	for _, r := range routes.Routes {
		server.Add(r.Method, r.Pattern, r.HandlerFunc).Name = r.Name
	}
}

func setupNotFoundHandler() {
	echo.NotFoundHandler = func(c echo.Context) error {

		msg := fmt.Sprintf("Resource not found :%s", c.Request().URL)
		return c.String(http.StatusNotFound, msg)
	}
}

func startServer(container *dig.Container) {
	err := container.Invoke(func(config *models.Configuration) {
		server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", config.Port)))
	})
	logger.ConditionalFatal("Server", "startServer", err)
}
