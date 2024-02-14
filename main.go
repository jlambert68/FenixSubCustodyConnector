package main

import (
	"FenixSubCustodyConnector/executionOrchestrator"
	executeTestInstructionUsingTestApiEngine "FenixSubCustodyConnector/externalTestInstructionExecutionsViaTestApiEngine"
)

func main() {

	// Initiate TestApiEngine
	executeTestInstructionUsingTestApiEngine.InitiateTestApiEngine()

	// Initiate ExecutionOrchestratorEngine
	executionOrchestrator.InitiateExecutionOrchestratorEngine(allowedUsers)

	// Keep program running
	var waitChannel chan bool
	waitChannel = make(chan bool)
	<-waitChannel

}
