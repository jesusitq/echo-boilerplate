package logger

import (
	"fmt"
	"io"
	goLog "log"
	"os"
	"strconv"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/log"
)

var (
	logger   zerolog.Logger
	hostname string
)

const (
	filePathEnv      string = "logs/bootstrap_%s"
	logLevelEnv      string = "LOG_LEVEL"          // Default: DEBUG
	enableLogFileEnv string = "ENABLE_FILE_LOG"    // Default: false
	enableConsoleEnv string = "ENABLE_CONSOLE_LOG" // Default: true
)

func init() {

	filePath := fmt.Sprintf(filePathEnv, getHostname())
	logLevel, _ := strconv.Atoi(getEnvVar(logLevelEnv, "0"))
	enableLogFile, _ := strconv.ParseBool(getEnvVar(enableLogFileEnv, "false"))
	enableConsoleLogs, _ := strconv.ParseBool(getEnvVar(enableConsoleEnv, "true"))
	var outputs []io.Writer

	zerolog.SetGlobalLevel(zerolog.Level(logLevel))

	if enableConsoleLogs {
		wr := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Logger Dropped %d messages", missed)
		})
		consoleOutput := zerolog.ConsoleWriter{Out: wr, TimeFormat: "2006-01-02 15:04:05", NoColor: true}
		consoleOutput.FormatLevel = func(text interface{}) string {
			return ""
		}
		consoleOutput.FormatMessage = func(text interface{}) string {
			return text.(string)
		}
		outputs = append(outputs, consoleOutput)
	}

	if enableLogFile {
		logf, _ := rotatelogs.New(filePath+".%Y%m%d%H.log",
			rotatelogs.WithClock(rotatelogs.Local),
			rotatelogs.WithRotationTime(time.Hour*1),
		)
		fileOutput := zerolog.ConsoleWriter{Out: logf, NoColor: true}
		fileOutput.FormatLevel = func(text interface{}) string {
			return ""
		}
		fileOutput.FormatMessage = func(text interface{}) string {
			return text.(string)
		}
		outputs = append(outputs, fileOutput)
	}

	//https://github.com/rs/zerolog/pull/198/files
	multi := zerolog.MultiLevelWriter(outputs...)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	log.Info().Msgf("Log level set to '%d'", logLevel)
}

func getEnvVar(env string, defaultValue string) string {
	value, found := os.LookupEnv(env)

	if !found || value == "" {
		return defaultValue
	}

	return value
}

func getHostname() string {
	if hostname == "" {
		h, err0 := os.Hostname()
		if err0 != nil {
			hostname = "UNKNOWN"
		} else {
			hostname = h
		}
	}
	return hostname
}

// Request :
func Request(method string, statusCode int, uri string, start time.Time) {
	log.Log().Msgf("REQS [%s][%d] %s\t%.2f ms", method, statusCode, uri, float32(time.Since(start).Nanoseconds())/1000000.0)
}

// Performance :
func Performance(moduleName string, functionName string, start time.Time) {
	log.Debug().Msgf("DEBG [%s][%s][%.2f]", moduleName, functionName, float32(time.Since(start).Nanoseconds())/1000000.0)
}

// Debug :
func Debug(moduleName string, functionName string, text string) {
	log.Debug().Msgf("DEBG [%s][%s] %s", moduleName, functionName, text)
}

// Info :
func Info(moduleName string, functionName string, text string) {
	log.Info().Msgf("INFO [%s][%s] %s", moduleName, functionName, text)
}

// Warn :
func Warn(moduleName string, functionName string, text string) {
	log.Warn().Msgf("WARN [%s][%s] %s", moduleName, functionName, text)
}

// Error :
func Error(moduleName string, functionName string, text string) {
	log.Error().Msgf("ERRO [%s][%s] %s", moduleName, functionName, text)
}

// Fatal : It's a Critical error with an Exit statement.
func Fatal(moduleName string, functionName string, text string) {
	goLog.Printf("CRIT [%s][%s] %s\n", moduleName, functionName, text)
	log.Fatal().Msgf("CRIT [%s][%s] %s", moduleName, functionName, text)
}

// ConditionalFatal :
func ConditionalFatal(moduleName string, functionName string, errs ...error) {
	text := ""
	for _, err := range errs {
		if err != nil {
			text += fmt.Sprintf("%s \t", err.Error())
		}
	}
	if len(text) > 0 {
		Fatal(moduleName, functionName, text)
	}
}
