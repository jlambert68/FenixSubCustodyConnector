package executeTestInstructionUsingTestApiEngine

import (
	"FenixSubCustodyConnector/sharedCode"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	TestApiEngineClassesAndMethodsAndAttributes "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	testApiEngineClassesAndMethods "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func ConvertTestInstructionExecutionIntoTestApiEngineRestCallMessage(
	processTestInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest) (
	TestApiEngineRestApiMessageValues *TestApiEngineRestApiMessageStruct,
	err error) {

	// Extract UUID:s from 'TestInstructionExecutionRequest'
	var (
		testInstructionUuid          string
		testCaseExecutionUuid        string
		testInstructionExecutionUuid string
	)
	testInstructionUuid = processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetTestInstructionUuid()
	testCaseExecutionUuid = processTestInstructionExecutionPubSubRequest.GetTestCaseExecutionUuid()
	testInstructionExecutionUuid = processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetTestInstructionExecutionUuid()

	// Extract TestInstructionAttributes from 'TestInstructionExecutionRequest'
	var testInstructionAttributes []*fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest_TestInstructionAttributeMessage
	testInstructionAttributes = processTestInstructionExecutionPubSubRequest.GetTestInstruction().GetTestInstructionAttributes()

	// Extract relevant testApiEngineData, from TestApiEngine-definitions to be used in mapping,
	var testApiEngineTestInstructionDataMapPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
	var existsInMap bool
	testApiEngineTestInstructionDataMapPtr, existsInMap = LocalExecutionMethods.FullTestApiEngineClassesMethodsAttributesVersionMap[TypeAndStructs.OriginalElementUUIDType(testInstructionUuid)]
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
		TestApiEngineClassNameNAME:           tempTestApiEngineClassesMethodsAttributes.TestApiEngineClassNameNAME,
		TestApiEngineMethodNameNAME:          tempTestApiEngineClassesMethodsAttributes.TestApiEngineMethodNameNAME,
		TestApiEngineExpectedToBePassedValue: "",
		TestInstructionAttribute:             nil,
		TestApiEngineAttributes:              make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*testApiEngineClassesAndMethods.TestApiEngineAttributesStruct),
	}

	// Loop all Attributes and populate message to be used for RestCall to TestApiEngine
	for _, testInstructionAttribute := range testInstructionAttributes {

		// Separate Attribute 'ExpectedToBePassed', which is used in url instead as a parameter in the body of the rest call
		if testInstructionAttribute.TestInstructionAttributeName != string(TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed) {

			// Create and add Attribute with value
			var tempTestInstructionAttributesUuidAndValue TestInstructionAttributesUuidAndValueStruct
			tempTestInstructionAttributesUuidAndValue = TestInstructionAttributesUuidAndValueStruct{
				TestInstructionAttributeUUID:          TypeAndStructs.TestInstructionAttributeUUIDType(testInstructionAttribute.TestInstructionAttributeUuid),
				TestInstructionAttributeName:          TypeAndStructs.TestInstructionAttributeNameType(testInstructionAttribute.TestInstructionAttributeName),
				TestInstructionAttributeValueAsString: TypeAndStructs.AttributeValueAsStringType(testInstructionAttribute.AttributeValueAsString),
			}
			TestApiEngineRestApiMessageValues.TestInstructionAttribute = append(
				TestApiEngineRestApiMessageValues.TestInstructionAttribute,
				tempTestInstructionAttributesUuidAndValue)

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

			// Create and add reference between Attribute and TestApiEngineAttribute-name to be used in RestRequest
			var tempTestApiEngineAttributesToUse *testApiEngineClassesAndMethods.TestApiEngineAttributesStruct
			tempTestApiEngineAttributesToUse = &testApiEngineClassesAndMethods.TestApiEngineAttributesStruct{
				TestInstructionAttributeUUID:     tempTestApiEngineAttributes.TestInstructionAttributeUUID,
				TestInstructionAttributeName:     tempTestApiEngineAttributes.TestInstructionAttributeName,
				TestInstructionAttributeTypeUUID: tempTestApiEngineAttributes.TestInstructionAttributeTypeUUID,
				TestApiEngineAttributeNameUUID:   tempTestApiEngineAttributes.TestApiEngineAttributeNameUUID,
				TestApiEngineAttributeNameName:   tempTestApiEngineAttributes.TestApiEngineAttributeNameName,
			}
			// Add Attribute
			TestApiEngineRestApiMessageValues.TestApiEngineAttributes[TypeAndStructs.TestInstructionAttributeUUIDType(testInstructionAttribute.TestInstructionAttributeUuid)] = tempTestApiEngineAttributesToUse

		} else {
			// Attribute is 'ExpectedToBePassedValue'
			TestApiEngineRestApiMessageValues.TestApiEngineExpectedToBePassedValue = TypeAndStructs.AttributeValueAsStringType(testInstructionAttribute.AttributeValueAsString)
		}

	}

	return TestApiEngineRestApiMessageValues, err
}

func PostTestInstructionUsingRestCall(
	testApiEngineRestApiMessageValues *TestApiEngineRestApiMessageStruct) (
	restResponse *http.Response, err error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "97a1ab6a-9994-4004-b457-5abbb6816f9e",
	}).Debug("2. Performing Http Post...")

	// Generate json-body for RestCall, need to do it manually, because of strange json-structure with parameters just added instead of using an array
	attributesMap := make(map[string]string)

	for _, testInstructionAttribute := range testApiEngineRestApiMessageValues.TestInstructionAttribute {

		// Add attribute to Map
		attributesMap[string(testInstructionAttribute.TestInstructionAttributeName)] = string(testInstructionAttribute.TestInstructionAttributeValueAsString)
	}

	attributesAsJson, err := json.Marshal(attributesMap)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":            "e1e74131-5040-43fa-abfc-1023f09d4388",
			"attributesMap": attributesMap,
		}).Error("Couldn't Marshal Attributes-map into json request")

		return nil, err
	}

	// Create request-url
	/*
		https://igc-SubCustodytestauto-cax-test.sebshift.dev.sebank.se/TestCaseExecution/ExecuteTestActionMethod/a/b?expectedToBePassed=true

		curl -X 'POST' \
		  'https://igc-SubCustodytestauto-cax-test.sebshift.dev.sebank.se/TestCaseExecution/ExecuteTestActionMethod/a/b?expectedToBePassed=true' \
		  -H 'accept: text/plain' \
		  -H 'Content-Type: application/json' \
		  -d '{
		  "additionalProp1": "string",
		  "additionalProp2": "string",
		  "additionalProp3": "string"
		}'
	*/
	var TestApiEngineUrl string
	TestApiEngineUrl = LocalExecutionMethods.TestApiEngineUrlPath + "/" + string(testApiEngineRestApiMessageValues.TestApiEngineClassNameNAME) +
		"/" + string(testApiEngineRestApiMessageValues.TestApiEngineMethodNameNAME) +
		"?" + "expectedToBePassed=" + string(testApiEngineRestApiMessageValues.TestApiEngineExpectedToBePassedValue)

	// Use Local web server for test or TestApiEngine
	if sharedCode.UseInternalWebServerForTestInsteadOfCallingTestApiEngine == true {
		// Use Local web server for testing
		TestApiEngineUrl = "http://:" + LocalWebServerPort + TestApiEngineUrl

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "724d84e8-ec94-4947-ac74-0c7e5c17cfb6",
		}).Debug("Posting TestInstruction to local web server for Tests")

	} else {
		// Use TestApiEngine address and port
		TestApiEngineUrl = LocalExecutionMethods.TestApiEngineAddress + ":" + LocalExecutionMethods.TestApiEnginePort + TestApiEngineUrl

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "fa461ab4-789f-4b2a-a215-2653567fe319",
		}).Debug("Posting TestInstruction to TestApiEngine")
	}
	// Create Http-client
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// Do RestCall to TestApiEngine
	restResponse, err = httpClient.Post(TestApiEngineUrl, "application/json; charset=utf-8", bytes.NewBuffer(attributesAsJson))
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":               "12b846ad-e8bf-41e0-8893-a1a7cef5f396",
			"TestApiEngineUrl": TestApiEngineUrl,
		}).Error("Couldn't do call to Rest-execution-server")

		return restResponse, err
	}

	return restResponse, err
}
