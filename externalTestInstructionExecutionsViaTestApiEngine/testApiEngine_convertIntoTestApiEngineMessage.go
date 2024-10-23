package executeTestInstructionUsingTestApiEngine

import (
	"FenixSubCustodyConnector/sharedCode"
	"errors"
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	TestApiEngineClassesAndMethodsAndAttributes "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	testApiEngineClassesAndMethods "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
	"strconv"
)

// ConvertTestInstructionExecutionIntoTestApiEngineRestCallMessage
// Converts an 'processTestInstructionExecutionPubSubRequest' into message to be sent to TestApiEngine
func ConvertTestInstructionExecutionIntoTestApiEngineRestCallMessage(
	processTestInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest,
	maximumExecutionDurationInSeconds int64) ( //timeoutTimeInSeconds int
	TestApiEngineRestApiMessageValues *TestApiEngineRestApiMessageStruct,
	err error) {

	// Extract UUID:s from 'TestInstructionExecutionRequest' and TestInstructionExecutionVersion
	var (
		testInstructionUuid             string
		testCaseExecutionUuid           string
		testInstructionExecutionUuid    string
		testInstructionExecutionVersion uint32
	)
	testInstructionUuid = processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetTestInstructionOriginalUuid()
	testCaseExecutionUuid = processTestInstructionExecutionPubSubRequest.GetTestCaseExecutionUuid()
	testInstructionExecutionUuid = processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetTestInstructionExecutionUuid()
	testInstructionExecutionVersion = processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetTestInstructionExecutionVersion()

	// Extract TestInstructionAttributes from 'TestInstructionExecutionRequest'
	var testInstructionAttributes []*fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest_TestInstructionAttributeMessage
	testInstructionAttributes = processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetTestInstructionAttributes()

	// Extract relevant testApiEngineData, from TestApiEngine-definitions to be used in mapping,
	var TestApiEngineClassesAndMethodsAndAttributesMap TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
	var testApiEngineTestInstructionDataMapPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
	var existsInMap bool

	TestApiEngineClassesAndMethodsAndAttributesMap = *TestApiEngineClassesAndMethodsAndAttributesMapPtr
	testApiEngineTestInstructionDataMapPtr, existsInMap =
		TestApiEngineClassesAndMethodsAndAttributesMap[TypeAndStructs.OriginalElementUUIDType(testInstructionUuid)] //LocalExecutionMethods.FullTestApiEngineClassesMethodsAttributesVersionMap[TypeAndStructs.OriginalElementUUIDType(testInstructionUuid)]
	if existsInMap != true {
		// Must exist in map
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                  "acbfdd00-ed23-4882-893c-0e6b4e61338f",
			"testInstructionUuid": testInstructionUuid,
		}).Error("Couldn't find correct testApiEngineData in 'testApiEngineClassesMethodsAttributesMap'")

		errorID := "4faf3e89-f647-494c-8cd2-3f0623db68c6"
		err = errors.New(fmt.Sprintf("couldn't find correct testApiEngineData in 'testApiEngineClassesMethodsAttributesMap' for TestInstructionUuid:'%s', [ErrorID='%s']", testInstructionUuid, errorID))

		return nil, err
	}

	// Map holding all versions of the specific TestInstruction, with data for mapping towards TestApiEngine
	var testApiEngineTestInstructionDataMap TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
	testApiEngineTestInstructionDataMap = *testApiEngineTestInstructionDataMapPtr

	// Version as 'string'
	var versionNumberAsString TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodNameVersion_SubCustody_Type // "1_0" or "13_3" ...)
	// Create the version number as a 'string'
	versionNumberAsString = TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodNameVersion_SubCustody_Type(
		strconv.Itoa(int(processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetMajorVersionNumber())) + "_" +
			strconv.Itoa(int(processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetMinorVersionNumber())))

	// Pointer to the structure holding Classes, Methods and Attribute conversion information
	var tempTestApiEngineClassesMethodsAttributesPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct
	tempTestApiEngineClassesMethodsAttributesPtr, existInMap := testApiEngineTestInstructionDataMap[versionNumberAsString]
	if existInMap == false {
		// Must exist in map
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                  "89922406-110e-4792-8c4b-9d48474bfc19",
			"testInstructionUuid": testInstructionUuid,
		}).Error("Couldn't find correct testApiEngineData in 'testApiEngineClassesMethodsAttributesMap'")

		errorID := "4faf3e89-f647-494c-8cd2-3f0623db68c6"
		err = errors.New(fmt.Sprintf("couldn't find correct testApiEngineData in 'testApiEngineClassesMethodsAttributesMap' for TestInstructionUuid:'%s', [ErrorID='%s']", testInstructionUuid, errorID))

		return nil, err
	}

	// Structure holding Classes, Methods and Attribute conversion information
	var tempTestApiEngineClassesMethodsAttributes TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct
	tempTestApiEngineClassesMethodsAttributes = *tempTestApiEngineClassesMethodsAttributesPtr

	// Values to be used in RestCall to TestApiEngine
	TestApiEngineRestApiMessageValues = &TestApiEngineRestApiMessageStruct{
		TestCaseExecutionUuid:                testCaseExecutionUuid,
		TestInstructionExecutionUuid:         testInstructionExecutionUuid,
		TestInstructionExecutionVersion:      testInstructionExecutionVersion,
		TestApiEngineClassName:               tempTestApiEngineClassesMethodsAttributes.TestApiEngineClassNameNAME,
		TestApiEngineMethodName:              tempTestApiEngineClassesMethodsAttributes.TestApiEngineMethodNameNAME,
		TestApiEngineExpectedToBePassedValue: "",
		TestInstructionAttribute:             nil,
		TestApiEngineAttributes:              make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*testApiEngineClassesAndMethods.TestApiEngineAttributesStruct),
		MaximumExecutionDurationInSeconds:    maximumExecutionDurationInSeconds,
	}

	// Loop all Attributes and populate message to be used for RestCall to TestApiEngine
	for _, testInstructionAttribute := range testInstructionAttributes {

		fmt.Println("Jobbar på: ", testInstructionAttribute.TestInstructionAttributeName, testInstructionAttribute.GetTestInstructionAttributeUuid())

		// Separate Attribute 'ExpectedToBePassed', which is used in url instead as a parameter in the body of the rest call
		// But also as an attribute in json-request-body
		if testInstructionAttribute.TestInstructionAttributeName == string(TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed) {
			// Attribute is 'ExpectedToBePassedValue' (in url)
			TestApiEngineRestApiMessageValues.TestApiEngineExpectedToBePassedValue = TypeAndStructs.AttributeValueAsStringType(testInstructionAttribute.AttributeValueAsString)
		}

		// Create and add Attribute with value
		var tempTestInstructionAttributesUuidAndValue TestInstructionAttributesUuidAndValueStruct
		tempTestInstructionAttributesUuidAndValue = TestInstructionAttributesUuidAndValueStruct{
			TestInstructionAttributeUUID:          TypeAndStructs.TestInstructionAttributeUUIDType(testInstructionAttribute.TestInstructionAttributeUuid),
			TestInstructionAttributeName:          TypeAndStructs.TestInstructionAttributeNameType(testInstructionAttribute.TestInstructionAttributeName),
			TestInstructionAttributeValueAsString: TypeAndStructs.AttributeValueAsStringType(testInstructionAttribute.AttributeValueAsString),
		}

		// Get TestApiEngine-attribute conversion-map from pointer
		var attributesMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
		attributesMap = *tempTestApiEngineClassesMethodsAttributes.Attributes

		// Extract TestApiEngineAttribute-data pointer
		var tempTestApiEngineAttributesPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
		tempTestApiEngineAttributesPtr, existsInMap = attributesMap[TypeAndStructs.TestInstructionAttributeUUIDType(testInstructionAttribute.TestInstructionAttributeUuid)]
		if existsInMap != true {
			// Must exist in map
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                           "c3fdcf73-f3b7-4ed2-9a9c-c25a64151bbc",
				"testInstructionUuid":          testInstructionUuid,
				"TestInstructionAttributeUuid": TypeAndStructs.TestInstructionAttributeUUIDType(testInstructionAttribute.TestInstructionAttributeUuid),
			}).Error("Couldn't find correct attribute in 'tempTestApiEngineClassesMethodsAttributes.Attributes'")

			errorID := "0ce6aea7-ce07-4ae9-8c9f-227efb621e92"
			err = errors.New(fmt.Sprintf("couldn't find correct testApiEngineData in 'testApiEngineData.Attributes' for TestInstructionAttributeUuid:'%s', [ErrorID='%s']", testInstructionUuid, errorID))

			return nil, err
		}

		// Extract TestApiEngineAttribute-data
		var tempTestApiEngineAttributes TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
		tempTestApiEngineAttributes = *tempTestApiEngineAttributesPtr

		// Check this attribute should be sent to TestApiEngine
		if tempTestApiEngineAttributes.AttributeShouldBeSentToTestApiEngine == true {

			// Create and add reference between Attribute and TestApiEngineAttribute-name to be used in RestRequest
			var tempTestApiEngineAttributesToUse *testApiEngineClassesAndMethods.TestApiEngineAttributesStruct
			tempTestApiEngineAttributesToUse = &testApiEngineClassesAndMethods.TestApiEngineAttributesStruct{
				TestInstructionAttributeUUID:         tempTestApiEngineAttributes.TestInstructionAttributeUUID,
				TestInstructionAttributeName:         tempTestApiEngineAttributes.TestInstructionAttributeName,
				TestInstructionAttributeTypeUUID:     tempTestApiEngineAttributes.TestInstructionAttributeTypeUUID,
				TestApiEngineAttributeNameUUID:       tempTestApiEngineAttributes.TestApiEngineAttributeNameUUID,
				TestApiEngineAttributeNameName:       tempTestApiEngineAttributes.TestApiEngineAttributeNameName,
				AttributeShouldBeSentToTestApiEngine: tempTestApiEngineAttributes.AttributeShouldBeSentToTestApiEngine,
			}
			// Add Attribute
			TestApiEngineRestApiMessageValues.TestApiEngineAttributes[TypeAndStructs.TestInstructionAttributeUUIDType(testInstructionAttribute.TestInstructionAttributeUuid)] = tempTestApiEngineAttributesToUse

			TestApiEngineRestApiMessageValues.TestInstructionAttribute = append(
				TestApiEngineRestApiMessageValues.TestInstructionAttribute,
				tempTestInstructionAttributesUuidAndValue)

			fmt.Println("From Attribute:", tempTestApiEngineAttributes)
			fmt.Println("To Attribute to use:", tempTestApiEngineAttributesToUse)
		} else {
			fmt.Println("tempTestApiEngineAttributes", tempTestApiEngineAttributes)
		}
	}

	return TestApiEngineRestApiMessageValues, err
}
