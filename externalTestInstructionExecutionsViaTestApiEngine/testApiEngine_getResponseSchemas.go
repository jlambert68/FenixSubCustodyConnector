package executeTestInstructionUsingTestApiEngine

import "github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"

// Depending on the TestInstruction to be executed, chose correct json-schemas
func getResponseSchemasToUse(
	testInstructionUUID TypeAndStructs.OriginalElementUUIDType) (
	finalTestInstructionExecutionResultAsJson *string,
	finalTestInstructionExecutionResultJsonSchema *string) {

	return finalTestInstructionExecutionResultAsJson, finalTestInstructionExecutionResultJsonSchema
}
