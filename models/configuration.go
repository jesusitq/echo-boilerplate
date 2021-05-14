package models

// Configuration :
type Configuration struct {
	Port        string
	Environment string `envconfig:"GO_ENV"`
	LogLevel    string `envconfig:"LOG_LEVEL"`
	Version     string `envconfig:"VERSION"`
}
