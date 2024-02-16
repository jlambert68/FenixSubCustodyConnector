package main

import (
	"FenixSubCustodyConnector/executionOrchestrator"
	executeTestInstructionUsingTestApiEngine "FenixSubCustodyConnector/externalTestInstructionExecutionsViaTestApiEngine"
	"fmt"
)

func main() {

	for a, b := range injectedVariablesMap {
		fmt.Println(a, *b)
	}

	// Initiate TestApiEngine
	executeTestInstructionUsingTestApiEngine.InitiateTestApiEngine()

	// Initiate ExecutionOrchestratorEngine
	executionOrchestrator.InitiateExecutionOrchestratorEngine(allowedUsers)

	// Keep program running
	var waitChannel chan bool
	waitChannel = make(chan bool)
	<-waitChannel

}
