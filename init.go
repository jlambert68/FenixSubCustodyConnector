package main

import (
	"FenixSubCustodyConnector/sharedCode"
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

// Extract environment variables used by this Connector-code
func init() {
	var err error

	// Extract environment variable for Port used by TestApiEngine web server
	sharedCode.UseInternalWebServerForTestInsteadOfCallingTestApiEngine, err = strconv.ParseBool(mustGetenv("UseInternalWebServerForTestInsteadOfCallingTestApiEngine"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseInternalWebServerForTestInsteadOfCallingTestApiEngine:' to a boolean, error: ", err)
		os.Exit(0)
	}

}
