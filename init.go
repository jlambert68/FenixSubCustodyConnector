package main

import (
	"github.com/jlambert68/FenixSyncShared/environmentVariables"
)

func init() {

	// Initiate object that extract Environment Variables or Injected Environment Variables
	environmentVariables.InitiateInjectedVariablesMap(&injectedVariablesMap)
}
