package main

import (
	"FenixSubCustodyConnector/executionOrchestrator"
)

func main() {

	// Initiate ExecutionOrchestratorEngine
	executionOrchestrator.InitiateExecutionOrchestratorEngine(allowedUsers)

	// Keep program running
	var waitChannel chan bool
	waitChannel = make(chan bool)
	<-waitChannel

}
