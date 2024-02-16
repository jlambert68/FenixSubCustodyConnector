package executeTestInstructionUsingTestApiEngine

import (
	"fmt"
	"github.com/jlambert68/FenixSyncShared/environmentVariables"
	"os"
	"strconv"
)

func Init() {
	var err error

	// Extract environment variable for Port used by TestApiEngine web server
	UseInternalWebServerForTestInsteadOfCallingTestApiEngine, err = strconv.ParseBool(environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("UseInternalWebServerForTestInsteadOfCallingTestApiEngine"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseInternalWebServerForTestInsteadOfCallingTestApiEngine:' to a boolean, error: ", err)
		os.Exit(0)
	}

	// Extract environment variable for Address used by local web server
	LocalWebServerAddress = environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("LocalWebServerAddress")

	// Extract environment variable for Port used by local web server
	_, err = strconv.ParseInt(environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("LocalWebServerPort"), 10, 64)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'LocalWebServerPort:' to an integer, error: ", err)
		os.Exit(0)
	}
	LocalWebServerPort = environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("LocalWebServerPort")
}
