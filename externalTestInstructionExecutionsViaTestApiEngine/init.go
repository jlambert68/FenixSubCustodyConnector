package executeTestInstructionUsingTestApiEngine

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func Init() {
	var err error

	// Extract environment variable for to be able to turn of all communication to and from Worker is turned off.
	TurnAllCommunicationWithWorker, err = strconv.ParseBool(mustGetenv("TurnAllCommunicationWithWorker"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'TurnAllCommunicationWithWorker:' to an boolean, error: ", err)
		os.Exit(0)
	}

	// Extract environment variable for to redirect calls from TestApiEngine to a local web server
	UseInternalWebServerForTestInsteadOfCallingTestApiEngine, err = strconv.ParseBool(mustGetenv("UseInternalWebServerForTestInsteadOfCallingTestApiEngine"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseInternalWebServerForTestInsteadOfCallingTestApiEngine:' to an boolean, error: ", err)
		os.Exit(0)
	}

	// Extract environment variable for Address used by local web server
	LocalWebServerAddress = mustGetenv("LocalWebServerAddress")

	// Extract environment variable for Port used by local web server
	_, err = strconv.ParseInt(mustGetenv("LocalWebServerPort"), 10, 64)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'LocalWebServerPort:' to an integer, error: ", err)
		os.Exit(0)
	}
	LocalWebServerPort = mustGetenv("LocalWebServerPort")
}
