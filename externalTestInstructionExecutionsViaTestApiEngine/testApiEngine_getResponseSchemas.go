package executeTestInstructionUsingTestApiEngine

import (
	"FenixSubCustodyConnector/sharedCode"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
)

// Depending on the TestInstruction to be executed, chose correct json-schemas
func GetResponseSchemasToUse(
	testInstructionUUID TypeAndStructs.OriginalElementUUIDType,
	testInstructionVersion string) (
	finalTestInstructionExecutionResultJsonSchema *string,
	responseVariablesJsonSchema *string) {

	// Get the json-schema for 'FinalTestInstructionExecutionResult'
	*finalTestInstructionExecutionResultJsonSchema = string(finalTestInstructionExecutionResultMessageJsonSchema)

	// Chose Response Schema depending on TestInstruction to be executed
	switch testInstructionUUID {

	// Send a MT540 on MQ
	case TestInstruction_SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540:

		// Extract json-schema depending on version
		switch testInstructionVersion {
		case "1_0":
			*responseVariablesJsonSchema = string(sendMT540_v1_0_ResponseVariablesMessageJsonSchema)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                     "dec47e30-5f2e-4891-a228-b00580d3dc31",
				"testInstructionUUID":    testInstructionUUID,
				"testInstructionVersion": testInstructionVersion,
			}).Error("Unhandled version")

		}

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                  "bca82e83-53bf-4803-9196-31e78966920e",
			"testInstructionUUID": testInstructionUUID,
		}).Fatal("Unknown TestInstruction Uuid")

	}

	return finalTestInstructionExecutionResultJsonSchema, responseVariablesJsonSchema
}
