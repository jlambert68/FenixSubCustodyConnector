package executionOrchestrator

import (
	executeTestInstructionUsingTestApiEngine "FenixSubCustodyConnector/externalTestInstructionExecutionsViaTestApiEngine"
	"FenixSubCustodyConnector/sharedCode"
	"fmt"
	"github.com/google/uuid"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

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

	// Validate that outgoing and incoming TestInstructionExecution is the same
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

	// Act on Response Variable Type
	switch testApiEngineFinalTestInstructionExecutionResult.ResponseVariableType {

	case executeTestInstructionUsingTestApiEngine.NoResponseVariableType:

		// Do nothing, because don't send back 'NoResponseVariableType'

	case executeTestInstructionUsingTestApiEngine.ResponseVariableType1Type:

		// Loop response variables from TestApiEngine
		for _, tempResponseVariableType1 := range testApiEngineFinalTestInstructionExecutionResult.ResponseVariablesType1 {

			// Create gRPC-Response variable
			var tempResponseVariableGrpc *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_ResponseVariableMessage
			tempResponseVariableGrpc = &fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage_ResponseVariableMessage{
				ResponseVariableUuid:          tempResponseVariableType1.ResponseVariableUUID,
				ResponseVariableName:          tempResponseVariableType1.ResponseVariableName,
				ResponseVariableTypeUuid:      tempResponseVariableType1.ResponseVariableTypeUuid,
				ResponseVariableTypeName:      tempResponseVariableType1.ResponseVariableTypeName,
				ResponseVariableValueAsString: tempResponseVariableType1.ResponseVariableValueAsString,
			}

			// Append to list of Response variables
			tempResponseVariablesGrpc = append(tempResponseVariablesGrpc, tempResponseVariableGrpc)

			continue
		}

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                   "2cfcbeaf-69f0-480e-adf2-d0bbefec42e6",
			"ResponseVariableType": testApiEngineFinalTestInstructionExecutionResult.ResponseVariableType,
			"testApiEngineFinalTestInstructionExecutionResult": testApiEngineFinalTestInstructionExecutionResult,
		}).Error("Unhandled 'ResponseVariableTypeType'")

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
