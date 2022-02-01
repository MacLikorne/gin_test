package pkg

import (
	"log"
	"os"
)

func GetEnvironmentVariable(environmentVariableName string) string {
	value, isOk := os.LookupEnv(environmentVariableName)
	if !isOk {
		log.Fatalf("Environment variable %s not found. Closing app.", environmentVariableName)
	}

	return value
}