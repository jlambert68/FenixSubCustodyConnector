package executionOrchestrator

import (
	executeTestInstructionUsingTestApiEngine "FenixSubCustodyConnector/externalTestInstructionExecutionsViaTestApiEngine"
	"FenixSubCustodyConnector/sharedCode"
	"errors"
	"fmt"
	fenixConnectorAdminShared_sharedCode "github.com/jlambert68/FenixConnectorAdminShared/common_config"
	"github.com/jlambert68/FenixConnectorAdminShared/fenixConnectorAdminShared"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
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

var allowedUsers []byte

func InitiateExecutionOrchestratorEngine(tempAllowedUsers []byte) {

	allowedUsers = tempAllowedUsers

	connectorFunctionsToDoCallBackOn = &fenixConnectorAdminShared_sharedCode.ConnectorCallBackFunctionsStruct{
		GetMaxExpectedFinishedTimeStamp:        getMaxExpectedFinishedTimeStamp,
		ProcessTestInstructionExecutionRequest: processTestInstructionExecutionRequest,
		InitiateLogger:                         initiateLogger,
		GenerateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers: generateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers,
	}
	fenixConnectorAdminShared.InitiateFenixConnectorAdminShared(connectorFunctionsToDoCallBackOn)

}

func getMaxExpectedFinishedTimeStamp(testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.
	ProcessTestInstructionExecutionPubSubRequest) (
	maxExpectedFinishedTimeStamp time.Time,
	err error) {

	var expectedExecutionDuration time.Duration

	// Depending on TestInstruction calculate or set 'MaxExpectedFinishedTimeStamp'
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid) {

	case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:
		/*
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


		*/
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
			"id": "5e2fda4c-e5fe-4c6d-88db-0fadcae1d5ca",
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid": testInstructionExecutionPubSubRequest.
				TestInstruction.TestInstructionUuid,
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
	TestInstructionsAndTesInstructionContainersAndAllowedUsers.
		GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody(allowedUsers)

	// Get the full structure for supported TestInstructions, TestInstructionContainers and Allowed Users
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsers =
		TestInstructionsAndTesInstructionContainersAndAllowedUsers.
			TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody

	return supportedTestInstructionsAndTestInstructionContainersAndAllowedUsers

}

func processTestInstructionExecutionRequest(
	testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest) (
	testInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage,
	err error) {

	// Depending on TestInstruction then choose how to execution the TestInstruction
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid) {

	// Send a MT540 on MQ
	case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:
		fmt.Println("case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:")

		var version string
		version = string(testInstructionExecutionPubSubRequest.TestInstruction.GetMajorVersionNumber()) +
			"_" +
			string(testInstructionExecutionPubSubRequest.TestInstruction.GetMinorVersionNumber())

		// Convert message into message that can be used when sedning to TestApiEngine
		var testApiEngineRestApiMessageValues *executeTestInstructionUsingTestApiEngine.TestApiEngineRestApiMessageStruct
		testApiEngineRestApiMessageValues, err = executeTestInstructionUsingTestApiEngine.
			ConvertTestInstructionExecutionIntoTestApiEngineRestCallMessage(testInstructionExecutionPubSubRequest)
		if err != nil {
			// Something wrong when converting the 'TestInstructionExecutionPubSubRequest' into TestApiEngine-structure
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                                    "3380f600-ef95-477f-bc6d-34e0695c51da",
				"err":                                   err.Error(),
				"testInstructionExecutionPubSubRequest": testInstructionExecutionPubSubRequest,
			}).Error("Something wrong when converting the 'TestInstructionExecutionPubSubRequest' into " +
				"TestApiEngine-structure")

			testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
				ClientSystemIdentification: nil,
				TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
					TestInstructionExecutionUuid,
				TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
					TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION_CAN_BE_RERUN,
				TestInstructionExecutionEndTimeStamp: timestamppb.Now(),
			}

			break
		}

		// Get Json-schemas to use
		var finalTestInstructionExecutionResultJsonSchema *string
		var responseVariablesJsonSchema *string
		finalTestInstructionExecutionResultJsonSchema, responseVariablesJsonSchema =
			executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
				TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540,
				version)

		// Do Rest-call to 'TestApiEngine' for executing the TestInstruction
		var testApiEngineFinalTestInstructionExecutionResult executeTestInstructionUsingTestApiEngine.TestApiEngineFinalTestInstructionExecutionResultStruct
		testApiEngineFinalTestInstructionExecutionResult, err = executeTestInstructionUsingTestApiEngine.
			PostTestInstructionUsingRestCall(
				testApiEngineRestApiMessageValues,
				finalTestInstructionExecutionResultJsonSchema,
				responseVariablesJsonSchema)

		if err != nil {
			// Something went wrong when doing RestApi-call
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                                    "c7f0986d-b32c-4300-b096-ed8b4b773229",
				"err":                                   err.Error(),
				"testInstructionExecutionPubSubRequest": testInstructionExecutionPubSubRequest,
				"testApiEngineRestApiMessageValues":     testApiEngineRestApiMessageValues,
			}).Error("Something went wrong when doing RestApi-call to execute the TestInstruction")

			testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
				ClientSystemIdentification: nil,
				TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
					TestInstructionExecutionUuid,
				TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
					TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
				TestInstructionExecutionEndTimeStamp: timestamppb.Now(),
			}
		}

		testApiEngineFinalTestInstructionExecutionResult

		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification: nil,
			TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
				TestInstructionExecutionUuid,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_FINISHED_OK,
			TestInstructionExecutionEndTimeStamp: timestamppb.Now(),
		}

	// Send a MT542 on MQ
	case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:
		fmt.Println("case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:")

		// Do Rest-call to 'TestApiEngine' for executing the TestInstruction
		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification: nil,
			TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
				TestInstructionExecutionUuid,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_FINISHED_OK,
			TestInstructionExecutionEndTimeStamp: timestamppb.Now(),
		}

	// Validate the MT544 based on Related Reference received from MT54x
	case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:
		fmt.Println("case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:")

		// Do Rest-call to 'TestApiEngine' for executing the TestInstruction
		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification: nil,
			TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
				TestInstructionExecutionUuid,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_FINISHED_OK,
			TestInstructionExecutionEndTimeStamp: timestamppb.Now(),
		}

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "ba4e0810-a870-4ab0-b2b1-2f5fc02c2bf7",
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid,
		}).Fatal("Unknown TestInstruction Uuid")
	}

	return testInstructionExecutionResultMessage, err
}
