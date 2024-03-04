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
	TestInstructionExecutionVersion      uint32
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

// TestApiEngineResponseStruct
// Specify the structure that the Json-response, from TestApiEngine, will be 'converted into
// This is the response from TestApiEngine
type TestApiEngineResponseWithResponseValueAsStringStruct struct {
	TestStepExecutionStatus TestStepExecutionStatusStruct `json:"testStepExecutionStatus"`
	Details                 string                        `json:"details"`
	ResponseValue           string                        `json:"responseValue"`
	ExecutionTimeStamp      string                        `json:"executionTimeStamp"`
}

type TestApiEngineResponseWithResponseValueAsTestApiEngineFinalTestInstructionExecutionResultStruct struct {
	TestStepExecutionStatus TestStepExecutionStatusStruct                          `json:"testStepExecutionStatus"`
	Details                 string                                                 `json:"details"`
	ResponseValue           TestApiEngineFinalTestInstructionExecutionResultStruct `json:"responseValue"`
	ExecutionTimeStamp      string                                                 `json:"executionTimeStamp"`
}

// TestStepExecutionStatusStruct
// Holds Status code and Text for TestApiEngine response
type TestStepExecutionStatusStruct struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
}

// TestApiEngineFinalTestInstructionExecutionResultStruct
// Specify the structure that the Json-response, from TestApiEngine, will be 'converted into
// This is the unique Fenix-parts of the TestApiEngine-response, which exists as an inner json 'ResponseValue'
type TestApiEngineFinalTestInstructionExecutionResultStruct struct {
	TestApiEngineResponseJsonSchemaVersion string                   `json:"TestApiEngineResponseJsonSchemaVersion"`
	TestInstructionExecutionUUID           string                   `json:"TestInstructionExecutionUuid"`
	TestInstructionExecutionVersion        string                   `json:"TestInstructionExecutionVersion"`
	TestInstructionExecutionStatus         string                   `json:"TestInstructionExecutionStatus"`
	TestInstructionExecutionStartTimeStamp string                   `json:"TestInstructionExecutionStartTimeStamp"`
	TestInstructionExecutionEndTimeStamp   string                   `json:"TestInstructionExecutionEndTimeStamp"`
	ResponseVariables                      []ResponseVariableStruct `json:"ResponseVariables"`
	LogPosts                               []LogPostStruct          `json:"LogPosts"`
}

// LogPostStruct
// within 'TestApiEngineFinalTestInstructionExecutionResultStruct'
// Hold one logpost item
type LogPostStruct struct {
	LogPostTimeStamp                     string                                      `json:"LogPostTimeStamp"`
	LogPostStatus                        string                                      `json:"LogPostStatus"`
	LogPostText                          string                                      `json:"LogPostText"`
	FoundVersusExpectedValueForVariables []FoundVersusExpectedValueForVariableStruct `json:"FoundVersusExpectedValue"`
}

// FoundVersusExpectedValueForVariableStruct within 'LogPostStruct'
// Holds one variables and its expected value vs found value
type FoundVersusExpectedValueForVariableStruct struct {
	VariableName              string                         `json:"VariableName"`
	VariableDescription       string                         `json:"VariableDescription"`
	FoundVersusExpectedValues FoundVersusExpectedValueStruct `json:"FoundVersusExpectedValues"`
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
	TestInstructionVersion        string `json:"TestInstructionVersion"`
	ResponseVariableUUID          string `json:"ResponseVariableUuid"`
	ResponseVariableName          string `json:"ResponseVariableName"`
	ResponseVariableTypeUuid      string `json:"ResponseVariableTypeUuid"`
	ResponseVariableTypeName      string `json:"ResponseVariableTypeName"`
	ResponseVariableValueAsString string `json:"ResponseVariableValueAsString"`
}
