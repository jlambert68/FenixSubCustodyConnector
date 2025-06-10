package executionOrchestrator

import (
	fenixConnectorAdminShared_sharedCode "github.com/jlambert68/FenixConnectorAdminShared/common_config"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
)

// SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers
// Holds the structure of TestInstructions, TestInstructionContainers and Users which are published to
// TestCaseBuilder-Server at start
var SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers *TestInstructionAndTestInstuctionContainerTypes.
	TestInstructionsAndTestInstructionsContainersStruct

// Initiate Call-Back-struct and initiate
var connectorFunctionsToDoCallBackOn *fenixConnectorAdminShared_sharedCode.ConnectorCallBackFunctionsStruct

var allowedUsers []byte
var templateUrlParameters []byte
var simpleTestDataFiles [][]byte
var supportedTestCaseMetaData []byte
var supportedTestSuiteMetaData []byte
