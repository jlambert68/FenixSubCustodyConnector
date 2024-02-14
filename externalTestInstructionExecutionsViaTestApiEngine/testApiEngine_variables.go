package executeTestInstructionUsingTestApiEngine

import (
	testApiEngineClassesAndMethods "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

// TestApiEngineRestApiMessageStruct
// Used when converting a TestInstruction into data to be used in RestCall to TestApiEngine
type TestApiEngineRestApiMessageStruct struct {
	TestCaseExecutionUuid                string
	TestInstructionExecutionUuid         string
	TestApiEngineClassNameNAME           testApiEngineClassesAndMethods.TestApiEngine_ClassName_Name_SubCustody_Type
	TestApiEngineMethodNameNAME          testApiEngineClassesAndMethods.TestApiEngine_MethodName_Name_SubCustody_Type
	TestApiEngineExpectedToBePassedValue TypeAndStructs.AttributeValueAsStringType
	TestInstructionAttribute             []TestInstructionAttributesUuidAndValueStruct
	TestApiEngineAttributes              map[TypeAndStructs.TestInstructionAttributeUUIDType]*testApiEngineClassesAndMethods.TestApiEngineAttributesStruct

	//Title     string `json:"title"`
	//Completed bool   `json:"completed"`
}

// TestInstructionAttributesUuidAndValueStruct
// Holds one Attribute UUID, Name and Value
type TestInstructionAttributesUuidAndValueStruct struct {
	TestInstructionAttributeUUID          TypeAndStructs.TestInstructionAttributeUUIDType
	TestInstructionAttributeName          TypeAndStructs.TestInstructionAttributeNameType
	TestInstructionAttributeValueAsString TypeAndStructs.AttributeValueAsStringType
}
