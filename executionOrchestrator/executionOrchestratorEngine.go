package executionOrchestrator

import (
	executeTestInstructionUsingTestApiEngine "FenixSubCustodyConnector/externalTestInstructionExecutionsViaTestApiEngine"
	"FenixSubCustodyConnector/sharedCode"
	"errors"
	"fmt"
	"github.com/google/uuid"
	fenixConnectorAdminShared_sharedCode "github.com/jlambert68/FenixConnectorAdminShared/common_config"
	"github.com/jlambert68/FenixConnectorAdminShared/fenixConnectorAdminShared"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers"
	TestInstruction_SendOnMQTypeMT_SendMT540 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540/version_1_0"
	TestInstruction_SendOnMQTypeMT_SendMT542 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT544 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544/version_1_0"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT546 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT546/version_1_0"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT548 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT548/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
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
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {

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
					"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid,
					"version": version,
				}).Error("Unhandled version")

				expectedExecutionDuration = 2 * time.Minute
				maxExpectedFinishedTimeStamp = time.Now().Add(expectedExecutionDuration)
			}


		*/
		var methodsForLocalExecutions *LocalExecutionMethods.MethodsForLocalExecutionsStruct

		var methodsForLocalExecutionsAsInterface interface{}
		var ok bool
		methodsForLocalExecutionsAsInterface = TestInstruction_SendOnMQTypeMT_SendMT540.TestInstruction_SubCustody_SendMT540.LocalExecutionMethods.Value
		methodsForLocalExecutions, ok = methodsForLocalExecutionsAsInterface.(*LocalExecutionMethods.MethodsForLocalExecutionsStruct)
		if ok != true {
			log.Fatalln("NEJ")
		}
		fmt.Println(methodsForLocalExecutions)

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
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.
				TestInstruction.TestInstructionOriginalUuid,
		}).Error("Unknown TestInstruction Uuid")

		err = errors.New(fmt.Sprintf("Unknown TestInstruction Uuid: %s",
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid))

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

	var errLogPostsToAdd []*fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage

	// Create a temporary Start-TimeStamp to be used for when something goes wrong
	var tempTestInstructionExecutionStartTimeStamp *timestamppb.Timestamp
	tempTestInstructionExecutionStartTimeStamp = timestamppb.Now()

	// Depending on TestInstruction then choose how to execution the TestInstruction
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {

	// Send a MT54x on MQ or Validate MT54x
	case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540,
		TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542,
		TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544,
		TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546,
		TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstructionUUID_SubCustody_ValidateMT548:

		// Extract maximum timeout time from TestInstruction
		var methodsForLocalExecutionsAsInterface interface{}
		switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {
		case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:
			methodsForLocalExecutionsAsInterface = TestInstruction_SendOnMQTypeMT_SendMT540.TestInstruction_SubCustody_SendMT540.LocalExecutionMethods.Value

		case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:
			methodsForLocalExecutionsAsInterface = TestInstruction_SendOnMQTypeMT_SendMT542.TestInstruction_SubCustody_SendMT542.LocalExecutionMethods.Value

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:
			methodsForLocalExecutionsAsInterface = TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstruction_SubCustody_ValidateMT544.LocalExecutionMethods.Value

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546:
			methodsForLocalExecutionsAsInterface = TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstruction_SubCustody_ValidateMT546.LocalExecutionMethods.Value

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstructionUUID_SubCustody_ValidateMT548:
			methodsForLocalExecutionsAsInterface = TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstruction_SubCustody_ValidateMT548.LocalExecutionMethods.Value

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                                    "cb71e52e-d27c-4c59-a12f-cebcf577ba0e",
				"testInstructionExecutionPubSubRequest": testInstructionExecutionPubSubRequest,
				"TestInstructionUuid":                   TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid),
			}).Fatalln("Unhandled 'TestInstructionUuid' when extracting Timeout-time to be used towards TestApiEngine")
		}

		// Convert from interface onto correct typ
		var methodsForLocalExecutions *LocalExecutionMethods.MethodsForLocalExecutionsStruct
		var couldBeConverted bool

		methodsForLocalExecutions, couldBeConverted = methodsForLocalExecutionsAsInterface.(*LocalExecutionMethods.MethodsForLocalExecutionsStruct)
		if couldBeConverted != true {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                                    "5a4fd568-7858-4e3b-825e-339663c7ac02",
				"testInstructionExecutionPubSubRequest": testInstructionExecutionPubSubRequest,
				"TestInstructionUuid":                   TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid),
			}).Fatalln("Couldn't convert 'interface-type into correct type")
		}

		// Extract the maximum allowed time before timeout occurs
		var maximumExecutionDurationInSeconds int64
		maximumExecutionDurationInSeconds = methodsForLocalExecutions.LocalParametersUsedInRunTime.ExpectedTestInstructionExecutionDurationInSeconds

		// Create version number to be used in attributes request
		// Also use version number when getting correct json-schemas
		var testInstructionVersion string
		testInstructionVersion = fmt.Sprintf("v%s.%s",
			strconv.Itoa(int(testInstructionExecutionPubSubRequest.TestInstruction.GetMajorVersionNumber())),
			strconv.Itoa(int(testInstructionExecutionPubSubRequest.TestInstruction.GetMinorVersionNumber())))

		// Convert message into message that can be used when sending to TestApiEngine
		var testApiEngineRestApiMessageValues *executeTestInstructionUsingTestApiEngine.TestApiEngineRestApiMessageStruct
		testApiEngineRestApiMessageValues, err = executeTestInstructionUsingTestApiEngine.
			ConvertTestInstructionExecutionIntoTestApiEngineRestCallMessage(
				testInstructionExecutionPubSubRequest,
				maximumExecutionDurationInSeconds)

		if err != nil {
			// Something wrong when converting the 'TestInstructionExecutionPubSubRequest' into TestApiEngine-structure
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                                    "3380f600-ef95-477f-bc6d-34e0695c51da",
				"err":                                   err.Error(),
				"testInstructionExecutionPubSubRequest": testInstructionExecutionPubSubRequest,
			}).Error("Something wrong when converting the 'TestInstructionExecutionPubSubRequest' into " +
				"TestApiEngine-structure")

			// Add a log post
			var logPostText string
			logPostText = fmt.Sprintf("Something wrong when converting the 'TestInstructionExecutionPubSubRequest' into "+
				"TestApiEngine-structure in Connector. "+
				"TestCaseExecutionUuid='%s', "+
				"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s', "+
				"TestInstructionExecutionVersion='%d'. "+
				"Errror='%s'",
				testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
				err.Error())

			// Generate new LogPostUuid
			var logPostUuid uuid.UUID
			logPostUuid, err = uuid.NewRandom()
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id": "94378951-2bff-4c45-906c-207bcc530951",
				}).Error("Failed to generate UUID")
			}

			var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
			errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
				LogPostUuid:                         logPostUuid.String(),
				LogPostTimeStamp:                    timestamppb.Now(),
				LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
				LogPostText:                         logPostText,
				FoundVersusExpectedValueForVariable: nil,
			}

			errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

			testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
				ClientSystemIdentification: nil,
				TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
					TestInstructionExecutionUuid,
				TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
					TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION_CAN_BE_RERUN,
				TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
				TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
				ResponseVariables:                      nil,
				LogPosts:                               errLogPostsToAdd,
			}

			break
		}

		// Get Json-schemas to use
		var requestMessageToTestApiEngineJsonSchema *string
		var testApiEngineResponseMessageJsonSchema *string
		var finalTestInstructionExecutionResultJsonSchema *string
		var responseVariablesJsonSchema *string

		// Get correct Response Schema depending on message type
		switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {

		// Send a MT540 on MQ
		case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:
			requestMessageToTestApiEngineJsonSchema, testApiEngineResponseMessageJsonSchema,
				finalTestInstructionExecutionResultJsonSchema, responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540,
					testInstructionVersion)

		case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:
			requestMessageToTestApiEngineJsonSchema, testApiEngineResponseMessageJsonSchema,
				finalTestInstructionExecutionResultJsonSchema, responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542,
					testInstructionVersion)

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:
			requestMessageToTestApiEngineJsonSchema, testApiEngineResponseMessageJsonSchema,
				finalTestInstructionExecutionResultJsonSchema, responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544,
					testInstructionVersion)

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546:
			requestMessageToTestApiEngineJsonSchema, testApiEngineResponseMessageJsonSchema,
				finalTestInstructionExecutionResultJsonSchema, responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546,
					testInstructionVersion)

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstructionUUID_SubCustody_ValidateMT548:
			requestMessageToTestApiEngineJsonSchema, testApiEngineResponseMessageJsonSchema,
				finalTestInstructionExecutionResultJsonSchema, responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstructionUUID_SubCustody_ValidateMT548,
					testInstructionVersion)

		default:
			testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
				ClientSystemIdentification: nil,
				TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
					TestInstructionExecutionUuid,
				TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
					TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
				TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
				TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
				ResponseVariables:                      nil,
				LogPosts:                               nil,
			}

			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "6f559867-9061-4985-8b01-38b01e5aacd6",
				"TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid)": TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid),
			}).Fatalln("Unhandled message when getting json-schema for Response message. Hard exit")

			break

		}

		// Do Rest-call to 'TestApiEngine' for executing the TestInstruction
		var testApiEngineFinalTestInstructionExecutionResult executeTestInstructionUsingTestApiEngine.TestApiEngineFinalTestInstructionExecutionResultStruct
		testApiEngineFinalTestInstructionExecutionResult, err = executeTestInstructionUsingTestApiEngine.
			PostTestInstructionUsingRestCall(
				testApiEngineRestApiMessageValues,
				requestMessageToTestApiEngineJsonSchema,
				testApiEngineResponseMessageJsonSchema,
				finalTestInstructionExecutionResultJsonSchema,
				responseVariablesJsonSchema,
				testInstructionVersion)

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
				TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
				TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
				ResponseVariables:                      nil,
				LogPosts:                               nil,
			}

			break
		}

		// Validate the TestApiEngine-response
		testInstructionExecutionResultMessage = validateAndConvertTestApiEngineResponse(
			tempTestInstructionExecutionStartTimeStamp,
			&testApiEngineFinalTestInstructionExecutionResult,
			testInstructionExecutionPubSubRequest)

		break

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "ba4e0810-a870-4ab0-b2b1-2f5fc02c2bf7",
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid,
		}).Fatal("Unknown TestInstruction Uuid")
	}

	return testInstructionExecutionResultMessage, err
}

// Validates the TestApiResponse
func validateAndConvertTestApiEngineResponse(
	tempTestInstructionExecutionStartTimeStamp *timestamppb.Timestamp,
	testApiEngineFinalTestInstructionExecutionResult *executeTestInstructionUsingTestApiEngine.
		TestApiEngineFinalTestInstructionExecutionResultStruct,
	testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest) (
	testInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage) {

	var err error
	var foundError bool
	var errLogPostsToAdd []*fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage

	// Validate that outgoing in incoming TestInstructionExecution is the same
	if testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid !=
		testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionUUID ||
		strconv.Itoa(int(testInstructionExecutionPubSubRequest.TestInstruction.
			TestInstructionExecutionVersion)) != testApiEngineFinalTestInstructionExecutionResult.
			TestInstructionExecutionVersion {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                                    "c033b03e-88e1-4d0a-8b92-e1603ccc13c8",
			"testInstructionExecutionPubSubRequest": testInstructionExecutionPubSubRequest,
			"testApiEngineFinalTestInstructionExecutionResult": testApiEngineFinalTestInstructionExecutionResult,
		}).Error("Incoming TestInstructionExecution is not the same as outgoing")

		// Add a log post
		var logPostText string
		logPostText = fmt.Sprintf("Incoming TestInstructionExecution is not the same as outgoing. "+
			"TestCaseExecutionUuid='%s', "+
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s' <> "+
			"testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionUUID='%s', "+
			"TestInstructionExecutionVersion='%d' <> testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionVersion='%s'",
			testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionUUID,
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionVersion)

		// Generate new LogPostUuid
		var logPostUuid uuid.UUID
		logPostUuid, err = uuid.NewRandom()
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "9105c341-ed86-4b61-81a7-cb0d35916cfe",
			}).Error("Failed to generate UUID")
		}

		var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
		errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
			LogPostUuid:                         logPostUuid.String(),
			LogPostTimeStamp:                    timestamppb.Now(),
			LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
			LogPostText:                         logPostText,
			FoundVersusExpectedValueForVariable: nil,
		}

		errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification: nil,
			TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
				TestInstructionExecutionUuid,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
			TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
			ResponseVariables:                      nil,
			LogPosts:                               errLogPostsToAdd,
		}

		// At least one error found
		foundError = true
	}

	// Convert TestInstructionExecutionStartTimeStamp into time-variable
	var testInstructionExecutionStartTimeStamp time.Time
	var timeStampLayoutForParser string //:= "2006-01-02 15:04:05.999999999 -0700 MST"
	timeStampLayoutForParser, err = sharedCode.GenerateTimeStampParserLayout(
		testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStartTimeStamp)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":                                     "7b664f24-ae55-442e-82a6-711f6cd76c7e",
			"err":                                    err,
			"TestInstructionExecutionStartTimeStamp": testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStartTimeStamp,
		}).Error("Couldn't generate parser layout from 'TestInstructionExecutionStartTimeStamp'")

		// Add a log post
		var logPostText string
		logPostText = fmt.Sprintf("Couldn't generate parser layout from 'TestInstructionExecutionStartTimeStamp'. "+
			"TestCaseExecutionUuid='%s', "+
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s', "+
			"TestInstructionExecutionVersion='%d'. "+
			"testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStartTimeStamp='%s"+
			"Errror='%s'",
			testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStartTimeStamp,
			err.Error())

		// Generate new LogPostUuid
		var logPostUuid uuid.UUID
		logPostUuid, err = uuid.NewRandom()
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "c61f5c4f-8983-4993-ba19-c3475a42116b",
			}).Error("Failed to generate UUID")
		}

		var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
		errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
			LogPostUuid:                         logPostUuid.String(),
			LogPostTimeStamp:                    timestamppb.Now(),
			LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
			LogPostText:                         logPostText,
			FoundVersusExpectedValueForVariable: nil,
		}

		errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification: nil,
			TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
				TestInstructionExecutionUuid,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
			TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
			ResponseVariables:                      nil,
			LogPosts:                               errLogPostsToAdd,
		}

		// At least one error found
		foundError = true

	} else {

		testInstructionExecutionStartTimeStamp, err = time.Parse(timeStampLayoutForParser,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStartTimeStamp)
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":  "bbe31015-dcb5-404d-bd9f-a235ac490555",
				"err": err,
				"TestInstructionExecutionStartTimeStamp": testApiEngineFinalTestInstructionExecutionResult.
					TestInstructionExecutionStartTimeStamp,
			}).Error("Couldn't parse 'TestInstructionExecutionStartTimeStamp' in " +
				"'testApiEngineFinalTestInstructionExecutionResult'")

			// Add a log post
			var logPostText string
			logPostText = fmt.Sprintf("Couldn't parse 'TestInstructionExecutionStartTimeStamp' in "+
				"'testApiEngineFinalTestInstructionExecutionResult'. "+
				"TestCaseExecutionUuid='%s', "+
				"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s', "+
				"TestInstructionExecutionVersion='%d'. "+
				"testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStartTimeStamp='%s"+
				"Errror='%s'",
				testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
				testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStartTimeStamp,
				err.Error())

			// Generate new LogPostUuid
			var logPostUuid uuid.UUID
			logPostUuid, err = uuid.NewRandom()
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id": "49d16433-4591-42ce-88dd-d469ac930c93",
				}).Error("Failed to generate UUID")
			}

			var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
			errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
				LogPostUuid:                         logPostUuid.String(),
				LogPostTimeStamp:                    timestamppb.Now(),
				LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
				LogPostText:                         logPostText,
				FoundVersusExpectedValueForVariable: nil,
			}

			errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

			testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
				ClientSystemIdentification: nil,
				TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
					TestInstructionExecutionUuid,
				TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
					TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
				TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
				TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
				ResponseVariables:                      nil,
				LogPosts:                               errLogPostsToAdd,
			}

			// At least one error found
			foundError = true
		}
	}

	// Convert TestInstructionExecutionEndTimeStamp into time-variable
	var testInstructionExecutionEndTimeStamp time.Time
	timeStampLayoutForParser, err = sharedCode.GenerateTimeStampParserLayout(
		testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionEndTimeStamp)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":  "41c1ff67-e643-4c33-a6e8-0ade0017cc0c",
			"err": err,
			"TestInstructionExecutionEndTimeStamp": testApiEngineFinalTestInstructionExecutionResult.
				TestInstructionExecutionEndTimeStamp,
		}).Error("Couldn't generate parser layout from 'TestInstructionExecutionEndTimeStamp'")

		// Add a log post
		var logPostText string
		logPostText = fmt.Sprintf("Couldn't generate parser layout from 'TestInstructionExecutionEndTimeStamp'. "+
			"TestCaseExecutionUuid='%s', "+
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s', "+
			"TestInstructionExecutionVersion='%d'. "+
			"testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionEndTimeStamp='%s"+
			"Errror='%s'",
			testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionEndTimeStamp,
			err.Error())

		// Generate new LogPostUuid
		var logPostUuid uuid.UUID
		logPostUuid, err = uuid.NewRandom()
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "54a77942-37f1-461d-b100-16d2c1beb651",
			}).Error("Failed to generate UUID")
		}

		var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
		errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
			LogPostUuid:                         logPostUuid.String(),
			LogPostTimeStamp:                    timestamppb.Now(),
			LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
			LogPostText:                         logPostText,
			FoundVersusExpectedValueForVariable: nil,
		}

		errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification: nil,
			TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
				TestInstructionExecutionUuid,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
			TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
			ResponseVariables:                      nil,
			LogPosts:                               errLogPostsToAdd,
		}

		// At least one error found
		foundError = true

	} else {

		testInstructionExecutionStartTimeStamp, err = time.Parse(timeStampLayoutForParser,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionEndTimeStamp)
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":  "4e50dc7d-ee62-43dc-b7a7-8ba9aee64df5",
				"err": err,
				"TestInstructionExecutionEndTimeStamp": testApiEngineFinalTestInstructionExecutionResult.
					TestInstructionExecutionEndTimeStamp,
			}).Error("Couldn't parse 'TestInstructionExecutionEndTimeStamp' in " +
				"'testApiEngineFinalTestInstructionExecutionResult'")

			// Add a log post
			var logPostText string
			logPostText = fmt.Sprintf("Couldn't parse 'TestInstructionExecutionEndTimeStamp' in "+
				"'testApiEngineFinalTestInstructionExecutionResult'. "+
				"TestCaseExecutionUuid='%s', "+
				"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s', "+
				"TestInstructionExecutionVersion='%d'. "+
				"testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionEndTimeStamp='%s"+
				"Errror='%s'",
				testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
				testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionEndTimeStamp,
				err.Error())

			// Generate new LogPostUuid
			var logPostUuid uuid.UUID
			logPostUuid, err = uuid.NewRandom()
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id": "924656e3-ce2b-4415-837e-4fe7533ee03f",
				}).Error("Failed to generate UUID")
			}

			var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
			errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
				LogPostUuid:                         logPostUuid.String(),
				LogPostTimeStamp:                    timestamppb.Now(),
				LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
				LogPostText:                         logPostText,
				FoundVersusExpectedValueForVariable: nil,
			}

			errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

			testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
				ClientSystemIdentification: nil,
				TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
					TestInstructionExecutionUuid,
				TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
					TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
				TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
				TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
				ResponseVariables:                      nil,
				LogPosts:                               errLogPostsToAdd,
			}

			// At least one error found
			foundError = true
		}
	}

	// Convert 'TestInstructionExecutionStatus' into gRPC-variable
	var testInstructionExecutionStatus int32
	var existInMap bool
	testInstructionExecutionStatus, existInMap = fenixExecutionWorkerGrpcApi.
		TestInstructionExecutionStatusEnum_value[testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStatus]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":                             "e14137df-6d55-4996-8547-8052e5269b97",
			"err":                            err,
			"TestInstructionExecutionStatus": testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStatus,
		}).Error("'TestInstructionExecutionStatus' in 'testApiEngineFinalTestInstructionExecutionResult' doesn't " +
			"exist within gRPC-definition")

		// Add a log post
		var logPostText string
		logPostText = fmt.Sprintf("'TestInstructionExecutionStatus' in "+
			"'testApiEngineFinalTestInstructionExecutionResult' doesn't exist within gRPC-definition. "+
			"TestCaseExecutionUuid='%s', "+
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s' <> "+
			"testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionUUID='%s', "+
			"TestInstructionExecutionVersion='%d' <> testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionVersion='%s', "+
			"testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStatus='%s'",
			testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionUUID,
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionVersion,
			testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionStatus)

		// Generate new LogPostUuid
		var logPostUuid uuid.UUID
		logPostUuid, err = uuid.NewRandom()
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "a46eac85-d73d-4224-84eb-f6364447953e",
			}).Error("Failed to generate UUID")
		}

		var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
		errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
			LogPostUuid:                         logPostUuid.String(),
			LogPostTimeStamp:                    timestamppb.Now(),
			LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
			LogPostText:                         logPostText,
			FoundVersusExpectedValueForVariable: nil,
		}

		errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

		testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
			ClientSystemIdentification: nil,
			TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
				TestInstructionExecutionUuid,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
			TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
			ResponseVariables:                      nil,
			LogPosts:                               errLogPostsToAdd,
		}

		// At least one error found
		foundError = true
	}

	// Convert ResponseValue
	var tempResponseVariablesGrpc []*fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_ResponseVariableMessage

	// Loop response variables from TestApiEngine
	for _, tempResponseVariable := range testApiEngineFinalTestInstructionExecutionResult.ResponseVariables {

		// Create gRPC-Response variable
		var tempResponseVariableGrpc *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_ResponseVariableMessage
		tempResponseVariableGrpc = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_ResponseVariableMessage{
			ResponseVariableUuid:          tempResponseVariable.ResponseVariableUUID,
			ResponseVariableName:          tempResponseVariable.ResponseVariableName,
			ResponseVariableTypeUuid:      tempResponseVariable.ResponseVariableTypeUuid,
			ResponseVariableTypeName:      tempResponseVariable.ResponseVariableTypeName,
			ResponseVariableValueAsString: tempResponseVariable.ResponseVariableValueAsString,
		}

		// Append to list of Response variables
		tempResponseVariablesGrpc = append(tempResponseVariablesGrpc, tempResponseVariableGrpc)
	}

	// Convert LogPosts
	var tempLogPostsGrpc []*fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage

	// Loop Logpost from TestApiEngine
	for _, tempLogPost := range testApiEngineFinalTestInstructionExecutionResult.LogPosts {

		// Create gRPC-FoundVersusExpectedValueForVariables
		var tempFoundVersusExpectedValueForVariablesGrpc []*fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage_FoundVersusExpectedValueForVariableMessage

		// Loop Found Versus Expected Values from TestApiEngine
		for _, tempFoundVersusExpectedValueForVariable := range tempLogPost.FoundVersusExpectedValueForVariables {
			var tempFoundVersusExpectedValueForVariableMessageGrpc *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage_FoundVersusExpectedValueForVariableMessage

			// gRPC Found Versus Expected Values variable
			tempFoundVersusExpectedValueForVariableMessageGrpc = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage_FoundVersusExpectedValueForVariableMessage{
				VariableName:        tempFoundVersusExpectedValueForVariable.VariableName,
				VariableDescription: tempFoundVersusExpectedValueForVariable.VariableDescription,
				FoundVersusExpectedValue: &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage_FoundVersusExpectedValueMessage{
					FoundValue:    tempFoundVersusExpectedValueForVariable.FoundVersusExpectedValues.FoundValue,
					ExpectedValue: tempFoundVersusExpectedValueForVariable.FoundVersusExpectedValues.ExpectedValue,
				},
			}
			// Add FoundVsExpected to list
			tempFoundVersusExpectedValueForVariablesGrpc = append(tempFoundVersusExpectedValueForVariablesGrpc,
				tempFoundVersusExpectedValueForVariableMessageGrpc)

		}

		// Convert 'LogPostStatus' into gRPC-variable
		var tempLogPostStatus int32
		tempLogPostStatus, existInMap = fenixExecutionWorkerGrpcApi.
			LogPostStatusEnum_value[tempLogPost.LogPostStatus]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":            "c1bd4194-f527-4758-85b0-9709158cf22e",
				"LogPostStatus": tempLogPost.LogPostStatus,
			}).Error("'LogPostStatus' in 'testApiEngineFinalTestInstructionExecutionResult' doesn't " +
				"exist within gRPC-definition")

			// Add a log post
			var logPostText string
			logPostText = fmt.Sprintf("'LogPostStatus' in "+
				"'testApiEngineFinalTestInstructionExecutionResult' doesn't exist within gRPC-definition. "+
				"TestCaseExecutionUuid='%s', "+
				"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s' <> "+
				"testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionUUID='%s', "+
				"TestInstructionExecutionVersion='%d' <> testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionVersion='%s', "+
				"tempLogPost.LogPostStatus='%s'",
				testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
				testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionUUID,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
				testApiEngineFinalTestInstructionExecutionResult.TestInstructionExecutionVersion,
				tempLogPost.LogPostStatus)

			// Generate new LogPostUuid
			var logPostUuid uuid.UUID
			logPostUuid, err = uuid.NewRandom()
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id": "83329796-9654-47f8-80d4-d966378d318d",
				}).Error("Failed to generate UUID")
			}

			var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
			errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
				LogPostUuid:                         logPostUuid.String(),
				LogPostTimeStamp:                    timestamppb.Now(),
				LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
				LogPostText:                         logPostText,
				FoundVersusExpectedValueForVariable: nil,
			}

			errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

			testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
				ClientSystemIdentification: nil,
				TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
					TestInstructionExecutionUuid,
				TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
					TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
				TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
				TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
				ResponseVariables:                      nil,
				LogPosts:                               errLogPostsToAdd,
			}

			// At least one error found
			foundError = true
		}

		// Convert TestInstructionExecutionStartTimeStamp into time-variable
		var logPostTimeStamp time.Time
		timeStampLayoutForParser, err = sharedCode.GenerateTimeStampParserLayout(
			tempLogPost.LogPostTimeStamp)
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":               "669665e6-f8f5-40bc-a4a6-dce19b969b86",
				"err":              err,
				"LogPostTimeStamp": tempLogPost.LogPostTimeStamp,
			}).Error("Couldn't generate parser layout from 'LogPostTimeStamp'")

			// Add a log post
			var logPostText string
			logPostText = fmt.Sprintf("Couldn't generate parser layout from 'LogPostTimeStamp'. "+
				"TestCaseExecutionUuid='%s', "+
				"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s', "+
				"TestInstructionExecutionVersion='%d'. "+
				"tempLogPost.LogPostTimeStamp='%s"+
				"Errror='%s'",
				testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
				testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
				tempLogPost.LogPostTimeStamp,
				err.Error())

			// Generate new LogPostUuid
			var logPostUuid uuid.UUID
			logPostUuid, err = uuid.NewRandom()
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id": "670d24b1-eec6-431b-ac67-4ca1627892c4",
				}).Error("Failed to generate UUID")
			}

			var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
			errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
				LogPostUuid:                         logPostUuid.String(),
				LogPostTimeStamp:                    timestamppb.Now(),
				LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
				LogPostText:                         logPostText,
				FoundVersusExpectedValueForVariable: nil,
			}

			errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

			testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
				ClientSystemIdentification: nil,
				TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
					TestInstructionExecutionUuid,
				TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
					TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
				TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
				TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
				ResponseVariables:                      nil,
				LogPosts:                               errLogPostsToAdd,
			}

			// At least one error found
			foundError = true

		} else {

			logPostTimeStamp, err = time.Parse(timeStampLayoutForParser,
				tempLogPost.LogPostTimeStamp)
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"Id":               "a1b99570-526c-4960-9bec-79ff6310dc70",
					"err":              err,
					"LogPostTimeStamp": tempLogPost.LogPostTimeStamp,
				}).Error("Couldn't parse 'LogPostTimeStamp' in " +
					"'testApiEngineFinalTestInstructionExecutionResult'")

				// Add a log post
				var logPostText string
				logPostText = fmt.Sprintf("Couldn't parse 'LogPostTimeStamp' in "+
					"'testApiEngineFinalTestInstructionExecutionResult'. "+
					"TestCaseExecutionUuid='%s', "+
					"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid='%s', "+
					"TestInstructionExecutionVersion='%d'. "+
					"tempLogPost.LogPostTimeStamp='%s"+
					"Errror='%s'",
					testInstructionExecutionPubSubRequest.TestCaseExecutionUuid,
					testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionUuid,
					testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionExecutionVersion,
					tempLogPost.LogPostTimeStamp,
					err.Error())

				// Generate new LogPostUuid
				var logPostUuid uuid.UUID
				logPostUuid, err = uuid.NewRandom()
				if err != nil {
					sharedCode.Logger.WithFields(logrus.Fields{
						"id": "cc2a2a1b-cf00-491c-be89-4c4103eefcec",
					}).Error("Failed to generate UUID")
				}

				var errLogPostToAdd *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
				errLogPostToAdd = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
					LogPostUuid:                         logPostUuid.String(),
					LogPostTimeStamp:                    timestamppb.Now(),
					LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR,
					LogPostText:                         logPostText,
					FoundVersusExpectedValueForVariable: nil,
				}

				errLogPostsToAdd = append(errLogPostsToAdd, errLogPostToAdd)

				testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{
					ClientSystemIdentification: nil,
					TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
						TestInstructionExecutionUuid,
					TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
						TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION,
					TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp,
					TestInstructionExecutionEndTimeStamp:   timestamppb.Now(),
					ResponseVariables:                      nil,
					LogPosts:                               errLogPostsToAdd,
				}

				// At least one error found
				foundError = true
			}
		}

		// At least one error found
		if foundError == true {
			return testInstructionExecutionResultMessage
		}

		// Generate new LogPostUuid
		var logPostUuid uuid.UUID
		logPostUuid, err = uuid.NewRandom()
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "cc6f7a43-3777-4f65-a9a0-d7c17fde815b",
			}).Error("Failed to generate UUID")
		}

		// Create gRPC-LogPost variable
		var tempLogPostGrpc *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage
		tempLogPostGrpc = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_LogPostMessage{
			LogPostUuid:                         logPostUuid.String(),
			LogPostTimeStamp:                    timestamppb.New(logPostTimeStamp),
			LogPostStatus:                       fenixExecutionWorkerGrpcApi.LogPostStatusEnum(tempLogPostStatus),
			LogPostText:                         tempLogPost.LogPostText,
			FoundVersusExpectedValueForVariable: tempFoundVersusExpectedValueForVariablesGrpc,
		}

		// Add LogPost to list of LogPosts
		tempLogPostsGrpc = append(tempLogPostsGrpc, tempLogPostGrpc)
	}

	testInstructionExecutionResultMessage = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage{

		ClientSystemIdentification: nil,
		TestInstructionExecutionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
			GetTestInstructionExecutionUuid(),
		TestInstructionExecutionVersion: int32(testInstructionExecutionPubSubRequest.GetTestInstruction().
			GetTestInstructionExecutionVersion()),
		MatureTestInstructionUuid: testInstructionExecutionPubSubRequest.GetTestInstruction().
			GetMatureTestInstructionUuid(),
		TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum(
			testInstructionExecutionStatus),
		TestInstructionExecutionStartTimeStamp: timestamppb.New(testInstructionExecutionStartTimeStamp),
		TestInstructionExecutionEndTimeStamp:   timestamppb.New(testInstructionExecutionEndTimeStamp),
		ResponseVariables:                      tempResponseVariablesGrpc,
		LogPosts:                               tempLogPostsGrpc,
	}

	return testInstructionExecutionResultMessage
}
