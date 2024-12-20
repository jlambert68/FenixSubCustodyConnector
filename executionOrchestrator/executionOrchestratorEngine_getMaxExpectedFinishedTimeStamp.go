package executionOrchestrator

import (
	"FenixSubCustodyConnector/sharedCode"
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	// Fenix 'SendTemplateToThisDomain'
	TestInstruction_Standard_SendTemplateToThisDomain "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain"
	testInstruction_SendTemplateToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0"
	// Fenix 'SendTestDataToThisDomain'
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain"
	testInstruction_SendTestDataToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0"

	// SubCustody 'SendMT540'
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540"
	TestInstruction_SendOnMQTypeMT_SendMT540_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540/version_1_0"
	// SubCustody 'SendMT542'
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542"
	TestInstruction_SendOnMQTypeMT_SendMT542_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0"
	// SubCustody 'ValidateMT544'
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT544_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544/version_1_0"
	// SubCustody 'ValidateMT546'
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT546"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT546_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT546/version_1_0"
	// SubCustody 'ValidateMT548'
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT548"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT548_version1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT548/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func getMaxExpectedFinishedTimeStamp(testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.
	ProcessTestInstructionExecutionPubSubRequest) (
	maxExpectedFinishedTimeStamp time.Time,
	err error) {

	// expectedExecutionDurationInSeconds is extracted from TestInstruction-data
	var expectedExecutionDurationInSeconds int64

	// Create Version for TestInstruction
	var version string
	version = fmt.Sprintf("v%s.%s",
		strconv.Itoa(int(testInstructionExecutionPubSubRequest.TestInstruction.GetMajorVersionNumber())),
		strconv.Itoa(int(testInstructionExecutionPubSubRequest.TestInstruction.GetMinorVersionNumber())))

	// Depending on TestInstruction calculate or set 'MaxExpectedFinishedTimeStamp'
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {

	// General TestInstruction, owned by Fenix, that can be forced to Connector by user
	// TestInstruction holds the TestData that the TestCase is using
	case TestInstruction_Standard_SendTestDataToThisDomain.TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain:
		switch version {
		case "v1.0":

			// Extract duration
			expectedExecutionDurationInSeconds = testInstruction_SendTemplateToThisDomain_version_1_0.
				ExpectedMaxTestInstructionExecutionDurationInSeconds

			// Create Max Finished TimeStamp
			maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "a37c16ba-09b4-40c7-97a2-fe282a84071c",
				"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionOriginalUuid,
				"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionName,
				"version": version,
			}).Fatalln("Unhandled version, will exit")

		}

	// General TestInstruction, owned by Fenix, that can be forced to Connector by user
	// TestInstruction holds a Template that is sent to the Connector
	case TestInstruction_Standard_SendTemplateToThisDomain.TestInstructionUUID_FenixSentToUsersDomain_SendTemplateToThisDomain:
		switch version {
		case "v1.0":

			// Extract duration
			expectedExecutionDurationInSeconds = testInstruction_SendTestDataToThisDomain_version_1_0.
				ExpectedMaxTestInstructionExecutionDurationInSeconds

			// Create Max Finished TimeStamp
			maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "a37c16ba-09b4-40c7-97a2-fe282a84071c",
				"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionOriginalUuid,
				"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionName,
				"version": version,
			}).Fatalln("Unhandled version, will exit")

		}

	// General TestInstruction that can be forced to Connector by user
	// TestInstruction holds a Template that the user picked
	/*
		case TestInstruction_SendOnMQTypeMT_SendGeneral.TestInstructionUUID_SendOnMQTypeMT_SendGeneral:
			switch version {
			case "v1.0":

				// Extract duration
				expectedExecutionDurationInSeconds = TestInstruction_SendOnMQTypeMT_SendGeneral_version_1_0.
					ExpectedMaxTestInstructionExecutionDurationInSeconds

				// Create Max Finished TimeStamp
				maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

			default:
				sharedCode.Logger.WithFields(logrus.Fields{
					"id": "52f699bb-5876-40bc-9a32-03dbfa8be55b",
					"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
						TestInstructionOriginalUuid,
					"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
						TestInstructionName,
					"version": version,
				}).Fatalln("Unhandled version, will exit")

			}

	*/

	case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:

		// Extract execution duration based on TestInstruction version
		switch version {
		case "v1.0":

			// Extract duration
			expectedExecutionDurationInSeconds = TestInstruction_SendOnMQTypeMT_SendMT540_version1_0.
				ExpectedMaxTestInstructionExecutionDurationInSeconds

			// Create Max Finished TimeStamp
			maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "86af1deb-795c-4a0a-b4ac-766ff5ab4668",
				"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionOriginalUuid,
				"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionName,
				"version": version,
			}).Fatalln("Unhandled version, will exit")

		}

	case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:

		// Extract execution duration based on TestInstruction version
		switch version {
		case "v1.0":

			// Extract duration
			expectedExecutionDurationInSeconds = TestInstruction_SendOnMQTypeMT_SendMT542_version1_0.
				ExpectedMaxTestInstructionExecutionDurationInSeconds

			// Create Max Finished TimeStamp
			maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "c2b0d15d-6259-4a7c-9503-022d33ffa4a3",
				"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionOriginalUuid,
				"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionName,
				"version": version,
			}).Fatalln("Unhandled version, will exit")

		}

	case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:

		// Extract execution duration based on TestInstruction version
		switch version {

		case "v1.0":

			// Extract duration
			expectedExecutionDurationInSeconds = TestInstruction_ValidateMQTypeMT54x_ValidateMT544_version1_0.
				ExpectedMaxTestInstructionExecutionDurationInSeconds

			// Create Max Finished TimeStamp
			maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "86af1deb-795c-4a0a-b4ac-766ff5ab4668",
				"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionOriginalUuid,
				"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionName,
				"version": version,
			}).Fatalln("Unhandled version, will exit")

		}

	case TestInstruction_ValidateMQTypeMT54x_ValidateMT546.TestInstructionUUID_SubCustody_ValidateMT546:

		// Extract execution duration based on TestInstruction version
		switch version {

		case "v1.0":

			// Extract duration
			expectedExecutionDurationInSeconds = TestInstruction_ValidateMQTypeMT54x_ValidateMT546_version1_0.
				ExpectedMaxTestInstructionExecutionDurationInSeconds

			// Create Max Finished TimeStamp
			maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "b4361d83-baf8-4681-947f-d2a8148223a7",
				"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionOriginalUuid,
				"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionName,
				"version": version,
			}).Fatalln("Unhandled version, will exit")

		}

	case TestInstruction_ValidateMQTypeMT54x_ValidateMT548.TestInstructionUUID_SubCustody_ValidateMT548:

		// Extract execution duration based on TestInstruction version
		switch version {

		case "v1.0":

			// Extract duration
			expectedExecutionDurationInSeconds = TestInstruction_ValidateMQTypeMT54x_ValidateMT548_version1_0.
				ExpectedMaxTestInstructionExecutionDurationInSeconds

			// Create Max Finished TimeStamp
			maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "e01ce31f-f2f9-494d-9824-a66032210c5a",
				"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionOriginalUuid,
				"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionName,
				"version": version,
			}).Fatalln("Unhandled version, will exit")

		}

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "4363d5dc-8901-437e-913a-3aae332f859c",
			"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.
				TestInstruction.TestInstructionOriginalUuid,
			"TestInstructionName": testInstructionExecutionPubSubRequest.
				TestInstruction.TestInstructionName,
		}).Fatalln("Unknown TestInstruction Uuid, will Exit")

	}

	return maxExpectedFinishedTimeStamp, err
}
