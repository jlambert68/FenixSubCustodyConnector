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
	tempSimpleTestDataFiles [][]byte) {

	allowedUsers = tempAllowedUsers
	templateUrlParameters = tempTemplateUrlParameters
	simpleTestDataFiles = tempSimpleTestDataFiles

	connectorFunctionsToDoCallBackOn = &fenixConnectorAdminShared_sharedCode.ConnectorCallBackFunctionsStruct{
		GetMaxExpectedFinishedTimeStamp:        getMaxExpectedFinishedTimeStamp,
		ProcessTestInstructionExecutionRequest: processTestInstructionExecutionRequest,
		InitiateLogger:                         initiateLogger,
		GenerateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers: generateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers,
		GenerateTemplateRepositoryConnectionParameters:                               generateTemplateRepositoryConnectionParameters,
		GenerateSimpleTestData: generateSimpleTestData,
	}
	fenixConnectorAdminShared.InitiateFenixConnectorAdminShared(connectorFunctionsToDoCallBackOn)

}

// Initiate logger with same logger as Shared Connector code uses
func initiateLogger(logger *logrus.Logger) {
	sharedCode.Logger = logger
}
