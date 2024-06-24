package main

import (
	"FenixSubCustodyConnector/executionOrchestrator"
	executeTestInstructionUsingTestApiEngine "FenixSubCustodyConnector/externalTestInstructionExecutionsViaTestApiEngine"
	"FenixSubCustodyConnector/sharedCode"
	"fmt"
	uuidGenerator "github.com/google/uuid"
)

func main() {

	//for a, b := range injectedVariablesMap {
	//	fmt.Println(a, *b)
	//}

	// Create Unique Uuid for run time instance used as identification when communication with GuiExecutionServer
	sharedCode.ApplicationRunTimeUuid = uuidGenerator.New().String()
	fmt.Println("sharedCode.ApplicationRunTimeUuid: " + sharedCode.ApplicationRunTimeUuid)

	// Initiate TestApiEngine
	executeTestInstructionUsingTestApiEngine.InitiateTestApiEngine()

	// Initiate ExecutionOrchestratorEngine
	executionOrchestrator.InitiateExecutionOrchestratorEngine(
		allowedUsers,
		templateUrlParameters)

	// Keep program running
	var waitChannel chan bool
	waitChannel = make(chan bool)
	<-waitChannel

}
