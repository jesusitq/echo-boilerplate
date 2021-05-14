package logger

import (
	"api-echo-template/libs/logger"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Logger :
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {

		start := time.Now()
		next(c)
		defer logger.Request(c.Request().Method, c.Response().Status, c.Request().RequestURI, start)

		return
	}
}

// LoggingResponseWriter :
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader :
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
