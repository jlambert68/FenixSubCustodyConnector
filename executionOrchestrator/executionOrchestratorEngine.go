package executionOrchestrator

import (
	"FenixSubCustodyConnector/sharedCode"
	fenixConnectorAdminShared_sharedCode "github.com/jlambert68/FenixConnectorAdminShared/common_config"
	"github.com/jlambert68/FenixConnectorAdminShared/fenixConnectorAdminShared"
	"github.com/sirupsen/logrus"
)

func InitiateExecutionOrchestratorEngine(
	tempAllowedUsers []byte,
	tempTemplateUrlParameters []byte,
	tempSimpleTestDataFiles [][]byte,
	tempSupportedTestCaseMetaData []byte,
	tempSupportedTestSuiteMetaData []byte,
	tempSupportedSubInstructions []byte,
	tempSupportedSubInstructionsPerTestInstructionSlice [][]byte) {

	allowedUsers = tempAllowedUsers
	templateUrlParameters = tempTemplateUrlParameters
	simpleTestDataFiles = tempSimpleTestDataFiles
	supportedTestCaseMetaData = tempSupportedTestCaseMetaData
	supportedTestSuiteMetaData = tempSupportedTestSuiteMetaData
	supportedSubInstructions = tempSupportedSubInstructions
	supportedSubInstructionsPerTestInstructionSlice = tempSupportedSubInstructionsPerTestInstructionSlice

	connectorFunctionsToDoCallBackOn = &fenixConnectorAdminShared_sharedCode.ConnectorCallBackFunctionsStruct{
		GetMaxExpectedFinishedTimeStamp:        getMaxExpectedFinishedTimeStamp,
		ProcessTestInstructionExecutionRequest: processTestInstructionExecutionRequest,
		InitiateLogger:                         initiateLogger,
		GenerateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers: generateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers,
		GenerateTemplateRepositoryConnectionParameters:                               generateTemplateRepositoryConnectionParameters,
		GenerateSimpleTestData:                             generateSimpleTestData,
		GenerateSupportedTestCaseMetaData:                  generateSupportedTestCaseMetaData,
		GenerateSupportedTestSuiteMetaData:                 generateSupportedTestSuiteMetaData,
		GenerateSupportedSubInstructions:                   generateSupportedSubInstructions(),
		GenerateSupportedSubInstructionsPerTestInstruction: generateSupportedSubInstructionsPerTestInstruction(),
	}
	fenixConnectorAdminShared.InitiateFenixConnectorAdminShared(connectorFunctionsToDoCallBackOn)

}

// Initiate logger with same logger as Shared Connector code uses
func initiateLogger(logger *logrus.Logger) {
	sharedCode.Logger = logger
}
