package executionOrchestrator

import (
	"FenixSubCustodyConnector/sharedCode"
	"errors"
	"fmt"
	fenixConnectorAdminShared_sharedCode "github.com/jlambert68/FenixConnectorAdminShared/common_config"
	"github.com/jlambert68/FenixConnectorAdminShared/fenixConnectorAdminShared"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	TestInstruction_Standard_IsServerAlive "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsServerAlive"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers"
	TestInstruction_SendOnMQTypeMT_SendMT540 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540/version_1_0"
	TestInstruction_SendOnMQTypeMT_SendMT542 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT544 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// Initiate Call-Back-struct and initiate
var connectorFunctionsToDoCallBackOn *fenixConnectorAdminShared_sharedCode.ConnectorCallBackFunctionsStruct

func InitiateExecutionOrchestratorEngine() {

	connectorFunctionsToDoCallBackOn = &fenixConnectorAdminShared_sharedCode.ConnectorCallBackFunctionsStruct{
		GetMaxExpectedFinishedTimeStamp:        getMaxExpectedFinishedTimeStamp,
		ProcessTestInstructionExecutionRequest: processTestInstructionExecutionRequest,
		InitiateLogger:                         initiateLogger,
		GenerateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers: generateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers,
	}
	fenixConnectorAdminShared.InitiateFenixConnectorAdminShared(connectorFunctionsToDoCallBackOn)

}

func getMaxExpectedFinishedTimeStamp(testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest) (
	maxExpectedFinishedTimeStamp time.Time,
	err error) {

	var expectedExecutionDuration time.Duration

	// Depending on TestInstruction calculate or set 'MaxExpectedFinishedTimeStamp'
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid) {

	case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:

		var version string
		version = string(testInstructionExecutionPubSubRequest.TestInstruction.GetMajorVersionNumber()) +
			"_" +
			string(testInstructionExecutionPubSubRequest.TestInstruction.GetMinorVersionNumber())

		// Extract execution duration depending on version
		switch version {
		case "1_0":
			TestInstruction_GeneralSetupTearDown_TestCaseSetUp_version_1_0..LocalExecutionMethods
		case "1_1":

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "ff8e9a06-cdca-45bc-bb19-24eb290a8502",
				"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid,
				"version": version,
			}).Error("Unhandled version")

			expectedExecutionDuration = 2 * time.Minute
			maxExpectedFinishedTimeStamp = time.Now().Add(expectedExecutionDuration)
		}

		expectedExecutionDuration = 2 * time.Minute
		maxExpectedFinishedTimeStamp = time.Now().Add(expectedExecutionDuration)

	case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:
		expectedExecutionDuration = 2 * time.Minute
		maxExpectedFinishedTimeStamp = time.Now().Add(expectedExecutionDuration)

	case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:
		expectedExecutionDuration = 2 * time.Minute
		maxExpectedFinishedTimeStamp = time.Now().Add(expectedExecutionDuration)

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "4192bcf6-09f7-4ee7-ad3d-4640bca4b2ba",
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid,
		}).Error("Unknown TestInstruction Uuid")

		err = errors.New(fmt.Sprintf("Unknown TestInstruction Uuid: %s",
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid))

		expectedExecutionDuration = 0 * time.Minute
		maxExpectedFinishedTimeStamp = time.Now().Add(expectedExecutionDuration)
	}

	return maxExpectedFinishedTimeStamp, err
}

// Initiate logger with same logger as Shared Connector code uses
func initiateLogger(logger *logrus.Logger) {
	sharedCode.Logger = logger
}

// Generates the 'SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers' that will be sent via gRPC to Worker
func generateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers() (
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsers *TestInstructionAndTestInstuctionContainerTypes.
		TestInstructionsAndTestInstructionsContainersStruct) {

	// Generate the full structure for supported TestInstructions, TestInstructionContainers and Allowed Users
	TestInstructionsAndTesInstructionContainersAndAllowedUsers.GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody()

	// Get the full structure for supported TestInstructions, TestInstructionContainers and Allowed Users
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsers =
		TestInstructionsAndTesInstructionContainersAndAllowedUsers.
			TestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo

	return supportedTestInstructionsAndTestInstructionContainersAndAllowedUsers

}

func processTestInstructionExecutionRequest(
	testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest) (
	testInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage,
	err error) {

	// Depending on TestInstruction then choose how to execution the TestInstruction
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid) {

	// Each TestCase is always started with this TestInstruction to set up everything
	case TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp:
		fmt.Println("case TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp:")

		// Do Rest-call to 'TestApiEngine' for executing the TestInstruction
		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification:           nil,
			TestInstructionExecutionUuid:         testInstructionExecutionPubSubRequest.GetTestInstruction().TestInstructionExecutionUuid,
			TestInstructionExecutionStatus:       fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum_TIE_FINISHED_OK,
			TestInstructionExecutionEndTimeStamp: timestamppb.Now(),
		}

	// Each TestCase is always ended with this TestInstruction to close down everything
	case TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown:
		fmt.Println("case TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown:")

		// Do Rest-call to 'TestApiEngine' for executing the TestInstruction
		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification:           nil,
			TestInstructionExecutionUuid:         testInstructionExecutionPubSubRequest.GetTestInstruction().TestInstructionExecutionUuid,
			TestInstructionExecutionStatus:       fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum_TIE_FINISHED_OK,
			TestInstructionExecutionEndTimeStamp: timestamppb.Now(),
		}

	// Checks if a certain date is a 'public holiday' for a country
	// Not executed by this Connector
	//case TestInstruction_Standard_IsDateAPublicHoliday.TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday:

	// Checks if a Server is 'Alive'
	case TestInstruction_Standard_IsServerAlive.TestInstructionUUID_OnPremDemo_IsServerAlive:

		// Call 'local' code for executing the TestInstruction
		testInstructionExecutionResultMessage, err = internalTestInstructionExecutions.TestInstruction_Standard_IsServerAlive(
			testInstructionExecutionPubSubRequest)

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "83c4d742-812e-4598-8a39-feabff216e11",
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid,
		}).Fatal("Unknown TestInstruction Uuid")
	}

	return testInstructionExecutionResultMessage, err
}
