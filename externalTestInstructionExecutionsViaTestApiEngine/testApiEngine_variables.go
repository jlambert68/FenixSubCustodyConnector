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

// TestApiEngineFinalTestInstructionExecutionResultStruct
// Specify the structure that the Json-response, from TestApiEngine, will be 'converted into
type TestApiEngineFinalTestInstructionExecutionResultStruct struct {
	TestInstructionExecutionUUID           string                   `json:"TestInstructionExecutionUuid"`
	TestInstructionExecutionVersion        string                   `json:"TestInstructionExecutionVersion"`
	TestInstructionExecutionStatus         string                   `json:"TestInstructionExecutionStatus"`
	TestInstructionExecutionStartTimeStamp string                   `json:"TestInstructionExecutionStartTimeStamp"`
	TestInstructionExecutionEndTimeStamp   string                   `json:"TestInstructionExecutionEndTimeStamp"`
	ResponseVariables                      []ResponseVariableStruct `json:"ResponseVariables"`
	LogPosts                               []LogPostStruct          `json:"LogPosts"`
}

// LogPostStruct, within 'TestApiEngineFinalTestInstructionExecutionResultStruct'
// Hold one logpost item
type LogPostStruct struct {
	LogPostTimeStamp                     string                                      `json:"LogPostTimeStamp"`
	LogPostStatus                        string                                      `json:"LogPostStatus"`
	LogPostText                          string                                      `json:"LogPostText"`
	FoundVersusExpectedValueForVariables []FoundVersusExpectedValueForVariableStruct `json:"FoundVersusExpectedValueForVariables"`
}

// FoundVersusExpectedValueForVariableStruct within 'LogPostStruct'
// Holds one variables and its expected value vs found value
type FoundVersusExpectedValueForVariableStruct struct {
	VariableName              string                           `json:"VariableName"`
	VariableDescription       string                           `json:"VariableDescription"`
	FoundVersusExpectedValues []FoundVersusExpectedValueStruct `json:"FoundVersusExpectedValues"`
}

// FoundVersusExpectedValueStruct within 'LogPostStruct'
// Holds one variables and its expected value vs found value
type FoundVersusExpectedValueStruct struct {
	FoundValue    string `json:"FoundValue"`
	ExpectedValue string `json:"ExpectedValue"`
}

// ResponseVariableStruct within 'TestApiEngineFinalTestInstructionExecutionResultStruct'
// Holds one response variable and its value
type ResponseVariableStruct struct {
	ResponseVariableUUID          string `json:"ResponseVariableUuid"`
	ResponseVariableName          string `json:"ResponseVariableName"`
	ResponseVariableValueAsString string `json:"ResponseVariableValueAsString"`
}
