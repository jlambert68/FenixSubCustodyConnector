package executionOrchestrator

import (
	executeTestInstructionUsingTestApiEngine "FenixSubCustodyConnector/externalTestInstructionExecutionsViaTestApiEngine"
	"FenixSubCustodyConnector/sharedCode"
	"fmt"
	"github.com/google/uuid"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	testInstruction_SendTestDataToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540"
	TestInstruction_SendOnMQTypeMT_SendMT540_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540/version_1_0"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542"
	TestInstruction_SendOnMQTypeMT_SendMT542_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT544_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544/version_1_0"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT546"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT546_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT546/version_1_0"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT548"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT548_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT548/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

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

	// General TestInstruction that can be forced to Connector by user
	// TestInstruction holds the TestData that the TestCase is using
	case testInstruction_SendTestDataToThisDomain_version_1_0.TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain:

		// Just log out the Data
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                                    "865f4047-b11d-4a04-886c-5ee9a2cd800a",
			"testInstructionExecutionPubSubRequest": testInstructionExecutionPubSubRequest,
		}).Info("The TestInstruction for TestData was sent to Connector")

	// Send a MT54x on MQ or Validate MT54x
	case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540,
		TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542,
		TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544,
		TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546,
		TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstructionUUID_SubCustody_ValidateMT548:

		// Extract the maximum allowed time before timeout occurs
		var maximumExecutionDurationInSeconds int64
		switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {
		case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:
			maximumExecutionDurationInSeconds = TestInstruction_SendOnMQTypeMT_SendMT540_version1_0.ExpectedMaxTestInstructionExecutionDurationInSeconds

		case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:
			maximumExecutionDurationInSeconds = TestInstruction_SendOnMQTypeMT_SendMT542_version1_0.ExpectedMaxTestInstructionExecutionDurationInSeconds

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:
			maximumExecutionDurationInSeconds = TestInstruction_ValidateMQTypeMT54x_ValidateMT544_version1_0.ExpectedMaxTestInstructionExecutionDurationInSeconds

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546:
			maximumExecutionDurationInSeconds = TestInstruction_ValidateMQTypeMT54x_ValidateMT546_version1_0.ExpectedMaxTestInstructionExecutionDurationInSeconds

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstructionUUID_SubCustody_ValidateMT548:
			maximumExecutionDurationInSeconds = TestInstruction_ValidateMQTypeMT54x_ValidateMT548_version1_0.ExpectedMaxTestInstructionExecutionDurationInSeconds

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                                    "cb71e52e-d27c-4c59-a12f-cebcf577ba0e",
				"testInstructionExecutionPubSubRequest": testInstructionExecutionPubSubRequest,
				"TestInstructionUuid":                   TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid),
			}).Fatalln("Unhandled 'TestInstructionUuid' when extracting Timeout-time to be used towards TestApiEngine")
		}

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
		var requestMethodParametersMessageToTestApiEngineJsonSchema *string
		var testApiEngineResponseMessageJsonSchema *string
		var finalTestInstructionExecutionResultJsonSchema *string
		var responseVariablesJsonSchema *string

		// Get correct Response Schema depending on message type
		switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {

		// Send a MT540 on MQ
		case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:
			requestMessageToTestApiEngineJsonSchema, requestMethodParametersMessageToTestApiEngineJsonSchema,
				testApiEngineResponseMessageJsonSchema, finalTestInstructionExecutionResultJsonSchema,
				responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540,
					testInstructionVersion)

		case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:
			requestMessageToTestApiEngineJsonSchema, requestMethodParametersMessageToTestApiEngineJsonSchema,
				testApiEngineResponseMessageJsonSchema, finalTestInstructionExecutionResultJsonSchema,
				responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542,
					testInstructionVersion)

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:
			requestMessageToTestApiEngineJsonSchema, requestMethodParametersMessageToTestApiEngineJsonSchema,
				testApiEngineResponseMessageJsonSchema, finalTestInstructionExecutionResultJsonSchema,
				responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544,
					testInstructionVersion)

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546:
			requestMessageToTestApiEngineJsonSchema, requestMethodParametersMessageToTestApiEngineJsonSchema,
				testApiEngineResponseMessageJsonSchema, finalTestInstructionExecutionResultJsonSchema,
				responseVariablesJsonSchema =
				executeTestInstructionUsingTestApiEngine.GetResponseSchemasToUse(
					TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546,
					testInstructionVersion)

		case TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstructionUUID_SubCustody_ValidateMT548:
			requestMessageToTestApiEngineJsonSchema, requestMethodParametersMessageToTestApiEngineJsonSchema,
				testApiEngineResponseMessageJsonSchema, finalTestInstructionExecutionResultJsonSchema,
				responseVariablesJsonSchema =
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
				requestMethodParametersMessageToTestApiEngineJsonSchema,
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
