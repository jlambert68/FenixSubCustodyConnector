package executeTestInstructionUsingTestApiEngine

import (
	"FenixSubCustodyConnector/sharedCode"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540"
	TestInstruction_SendOnMQTypeMT_SendMT542 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0"
	TestInstruction_ValidateMQTypeMT54x_ValidateMT544 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
)

// Depending on the TestInstruction to be executed, chose correct json-schemas
func GetResponseSchemasToUse(
	testInstructionUUID TypeAndStructs.OriginalElementUUIDType,
	testInstructionVersion string) (
	requestMessageToTestApiEngineJsonSchema *string,
	testApiEngineResponseMessageJsonSchema *string,
	finalTestInstructionExecutionResultJsonSchema *string,
	responseVariablesJsonSchema *string) {

	// Get the json-schema for 'TestApiEngineResponse'
	var tempTestApiEngineResponseMessageJsonSchema string
	tempTestApiEngineResponseMessageJsonSchema = string(testApiEngineResponseMessageJsonSchemaAsByteArray)
	testApiEngineResponseMessageJsonSchema = &tempTestApiEngineResponseMessageJsonSchema

	// Get the json-schema for 'FinalTestInstructionExecutionResult'
	var tempFinalTestInstructionExecutionResultJsonSchema string
	tempFinalTestInstructionExecutionResultJsonSchema = string(finalTestInstructionExecutionResultMessageJsonSchemaAsByteArray)
	finalTestInstructionExecutionResultJsonSchema = &tempFinalTestInstructionExecutionResultJsonSchema

	// Chose Response Schema depending on TestInstruction to be executed
	switch testInstructionUUID {

	// Send a MT540 on MQ
	case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:

		// Extract json-schema depending on version
		switch testInstructionVersion {
		case "v1.0":

			// Outgoing Request
			var tempRequestMessageToTestApiEngineJsonSchema string
			tempRequestMessageToTestApiEngineJsonSchema = string(sendMT540_v1_0_RequestMessageJsonSchemaAsByteArray)
			requestMessageToTestApiEngineJsonSchema = &tempRequestMessageToTestApiEngineJsonSchema

			// Incoming Response
			var tempResponseVariablesJsonSchema string
			tempResponseVariablesJsonSchema = string(sendMT540_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray)
			responseVariablesJsonSchema = &tempResponseVariablesJsonSchema

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                     "dec47e30-5f2e-4891-a228-b00580d3dc31",
				"testInstructionUUID":    testInstructionUUID,
				"testInstructionVersion": testInstructionVersion,
			}).Fatal("Unhandled version")

		}

		// Send a MT542 on MQ
	case TestInstruction_SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542:

		// Extract json-schema depending on version
		switch testInstructionVersion {
		case "v1.0":

			// Outgoing Request
			var tempRequestMessageToTestApiEngineJsonSchema string
			tempRequestMessageToTestApiEngineJsonSchema = string(sendMT542_v1_0_RequestMessageJsonSchemaAsByteArray)
			requestMessageToTestApiEngineJsonSchema = &tempRequestMessageToTestApiEngineJsonSchema

			// Incoming Response
			var tempResponseVariablesJsonSchema string
			tempResponseVariablesJsonSchema = string(sendMT542_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray)
			responseVariablesJsonSchema = &tempResponseVariablesJsonSchema

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                     "4d13b512-baa7-439b-a174-90e2859e259c",
				"testInstructionUUID":    testInstructionUUID,
				"testInstructionVersion": testInstructionVersion,
			}).Fatal("Unhandled version for 'TestInstructionUUID_SubCustody_SendMT542'")

		}

		// Validate a MT544
	case TestInstruction_ValidateMQTypeMT54x_ValidateMT544.TestInstructionUUID_SubCustody_ValidateMT544:

		// Extract json-schema depending on version
		switch testInstructionVersion {
		case "v1.0":

			// Outgoing Request
			var tempRequestMessageToTestApiEngineJsonSchema string
			tempRequestMessageToTestApiEngineJsonSchema = string(validateMT544_v1_0_RequestMessageJsonSchemaAsByteArray)
			requestMessageToTestApiEngineJsonSchema = &tempRequestMessageToTestApiEngineJsonSchema

			// Incoming Response
			var tempResponseVariablesJsonSchema string
			tempResponseVariablesJsonSchema = string(validateMT544_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray)
			responseVariablesJsonSchema = &tempResponseVariablesJsonSchema

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                     "4f75a47a-d2b6-4d63-a198-dd6e3964551d",
				"testInstructionUUID":    testInstructionUUID,
				"testInstructionVersion": testInstructionVersion,
			}).Fatal("Unhandled version for 'TestInstructionUUID_SubCustody_ValidateMT544'")

		}

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                  "bca82e83-53bf-4803-9196-31e78966920e",
			"testInstructionUUID": testInstructionUUID,
		}).Fatal("Unknown TestInstruction Uuid")

	}

	return requestMessageToTestApiEngineJsonSchema, finalTestInstructionExecutionResultJsonSchema, responseVariablesJsonSchema
}
