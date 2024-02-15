package main

import (
	"FenixSubCustodyConnector/sharedCode"
	"fmt"
	"github.com/jlambert68/FenixSyncShared/environmentVariables"
	"os"
	"strconv"
)

// Extract environment variables used by this Connector-code
func init() {
	var err error

	// Initiate object that extract Environment Variables or Injected Environment Variables
	environmentVariables.InitiateInjectedVariablesMap(&injectedVariablesMap)

	// Extract environment variable for Port used by TestApiEngine web server
	sharedCode.UseInternalWebServerForTestInsteadOfCallingTestApiEngine, err = strconv.ParseBool(environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("UseInternalWebServerForTestInsteadOfCallingTestApiEngine"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseInternalWebServerForTestInsteadOfCallingTestApiEngine:' to a boolean, error: ", err)
		os.Exit(0)
	}

}
