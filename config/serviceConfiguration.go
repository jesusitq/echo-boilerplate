package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"api-echo-template/models"

	"github.com/kelseyhightower/envconfig"
)

// https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152

var (
	configuration = models.Configuration{}
)

// NewConfiguration :
func NewConfiguration() *models.Configuration {
	loadConfig()
	return &configuration
}

// loadConfig :
func loadConfig() {
	readConfigFile()
	readConfigEnv()
}

func readConfigFile() {
	var _, b, _, _ = runtime.Caller(0)
	var basepath = filepath.Dir(b)
	fmt.Println("Environment: ", os.Getenv("GO_ENV"))
	file, err := os.Open(basepath + "/config." + os.Getenv("GO_ENV") + ".json")

	if err != nil {
		log.Fatalf("[Error]: %s", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalf("[Error]: %s", err)
	}
}

func readConfigEnv() {
	err := envconfig.Process("", &configuration)
	if err != nil {
		log.Fatalf("[Error]: %s", err)
	}
}
