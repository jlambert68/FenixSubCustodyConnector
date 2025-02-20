package executeTestInstructionUsingTestApiEngine

import (
	"FenixSubCustodyConnector/sharedCode"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func PostTestInstructionUsingRestCall(
	testApiEngineRestApiMessageValues *TestApiEngineRestApiMessageStruct,
	requestMessageToTestApiEngineJsonSchema *string,
	requestMethodParametersMessageToTestApiEngineJsonSchema *string,
	testApiEngineResponseMessageJsonSchema *string,
	finalTestInstructionExecutionResultJsonSchema *string,
	responseVariablesJsonSchema *string,
	testInstructionVersion string) (
	testApiEngineFinalTestInstructionExecutionResult TestApiEngineFinalTestInstructionExecutionResultStruct,
	err error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "97a1ab6a-9994-4004-b457-5abbb6816f9e",
	}).Debug("2. Performing Http Post...")

	// Generate json-body for RestCall, need to do it manually, because of strange json-structure with parameters just added instead of using an array
	attributesMap := make(map[string]string)

	for _, testInstructionAttribute := range testApiEngineRestApiMessageValues.TestInstructionAttribute {

		// Add attribute to Map
		attributesMap[string(testInstructionAttribute.TestInstructionAttributeName)] = string(testInstructionAttribute.
			TestInstructionAttributeValueAsString)
	}

	// Create the struct holding full Rest-request to TestApiEngine
	type testApiEngineRequestStruct struct {
		TestStepClassName     string            `json:"testStepClassName"`
		TestStepActionMethod  string            `json:"testStepActionMethod"`
		TestDataParameterType string            `json:"testDataParameterType"`
		ExpectedToBePassed    bool              `json:"expectedToBePassed"`
		MethodParameters      map[string]string `json:"methodParameters"`
	}

	// Add TestCaseExecutionUuid, TestInstructionExecutionUuid, TestInstructionExecutionVersion,
	// TestInstructionVersion and TimeoutTimeInSeconds
	attributesMap["TestStepActionMethod"] = string(testApiEngineRestApiMessageValues.TestApiEngineMethodName)
	attributesMap["TestInstructionVersion"] = testInstructionVersion
	attributesMap["TestCaseExecutionUuid"] = testApiEngineRestApiMessageValues.TestCaseExecutionUuid
	attributesMap["TestInstructionExecutionUuid"] = testApiEngineRestApiMessageValues.TestInstructionExecutionUuid
	attributesMap["TestInstructionExecutionVersion"] = strconv.Itoa(int(testApiEngineRestApiMessageValues.
		TestInstructionExecutionVersion))
	//Only add 'ExpectedToBePassed' if it is used as an attribute
	if len(testApiEngineRestApiMessageValues.TestApiEngineExpectedToBePassedValue) > 0 {
		attributesMap["ExpectedToBePassed"] = string(testApiEngineRestApiMessageValues.TestApiEngineExpectedToBePassedValue)
	}
	attributesMap["TimeoutTimeInSeconds"] = strconv.Itoa(int(testApiEngineRestApiMessageValues.MaximumExecutionDurationInSeconds))

	// Generate json from AttributesMap
	var attributesMapAsJson []byte

	attributesMapAsJson, err = json.Marshal(attributesMap)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":            "5776274d-df22-4e6d-951a-e2680b754b5a",
			"attributesMap": attributesMap,
		}).Error("Couldn't Marshal Attributes-map into json request")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Request to be sent to TestApiEngine
	var tempTestApiEngineRequest testApiEngineRequestStruct
	tempTestApiEngineRequest = testApiEngineRequestStruct{
		TestStepClassName:     string(testApiEngineRestApiMessageValues.TestApiEngineClassName),
		TestStepActionMethod:  string(testApiEngineRestApiMessageValues.TestApiEngineMethodName),
		TestDataParameterType: "FixedValue",
		ExpectedToBePassed:    true,
		MethodParameters:      map[string]string{"MethodParametersJsonAsString": string(attributesMapAsJson)},
	}

	// Marshal to json as byte-array
	var testApiEngineRequestAsJson []byte
	testApiEngineRequestAsJson, err = json.Marshal(tempTestApiEngineRequest)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                       "e1e74131-5040-43fa-abfc-1023f09d4388",
			"tempTestApiEngineRequest": tempTestApiEngineRequest,
		}).Error("Couldn't Marshal TestApiEngineRequest into json request")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Validate rest-request and convert
	err = validateRestRequest(
		&testApiEngineRequestAsJson,
		&attributesMapAsJson,
		requestMessageToTestApiEngineJsonSchema,
		requestMethodParametersMessageToTestApiEngineJsonSchema)

	if err != nil {
		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
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
	var testApiEngineUrl string
	testApiEngineUrl = LocalExecutionMethods.TestApiEngineUrlPath

	// Use Local web server for test or TestApiEngine
	if UseInternalWebServerForTestInsteadOfCallingTestApiEngine == true {
		// Use Local web server for testing
		testApiEngineUrl = "http://:" + LocalWebServerPort + testApiEngineUrl

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "724d84e8-ec94-4947-ac74-0c7e5c17cfb6",
		}).Debug("Posting TestInstruction to local web server for Tests")

	} else {
		// Use TestApiEngine address and port
		testApiEngineUrl = LocalExecutionMethods.TestApiEngineAddress + ":" + LocalExecutionMethods.TestApiEnginePort + testApiEngineUrl

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":               "fa461ab4-789f-4b2a-a215-2653567fe319",
			"testApiEngineUrl": testApiEngineUrl,
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
	var restResponse *http.Response
	restResponse, err = httpClient.Post(
		testApiEngineUrl,
		"application/json; charset=utf-8",
		bytes.NewBuffer(testApiEngineRequestAsJson))
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":               "12b846ad-e8bf-41e0-8893-a1a7cef5f396",
			"testApiEngineUrl": testApiEngineUrl,
		}).Error("Couldn't do call to Rest-execution-server")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Read the body
	defer restResponse.Body.Close()
	body, err := ioutil.ReadAll(restResponse.Body)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":               "838f461a-d207-450b-a46d-dc4557f64422",
			"testApiEngineUrl": testApiEngineUrl,
		}).Error("Couldn't extract json-body")
	}

	var bodyAsString string
	bodyAsString = string(body)

	// Validate rest-response and convert into 'TestApiEngineFinalTestInstructionExecutionResultStruct'
	testApiEngineFinalTestInstructionExecutionResult, err = validateAndTransformRestResponse(
		&bodyAsString,
		testApiEngineResponseMessageJsonSchema,
		finalTestInstructionExecutionResultJsonSchema,
		responseVariablesJsonSchema)

	if err != nil {
		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	return testApiEngineFinalTestInstructionExecutionResult, err
}

// Validate the json-request to be sent to TestApiEngine
// Validation os done with supported json-schemas
func validateRestRequest(
	testApiEngineRequestAsJson *[]byte,
	attributesMapAsJson *[]byte,
	requestMessageToTestApiEngineJsonSchema *string,
	requestMethodParametersMessageToTestApiEngineJsonSchema *string) (
	err error) {

	var tempJsonSchema string
	var testApiEngineRequestAsJsonByteArray []byte

	// 	// *** First Step - Validate the overall Request***
	tempJsonSchema = *requestMessageToTestApiEngineJsonSchema

	// Compile json-schema 'requestMessageToTestApiEngineJsonSchema'
	var jsonSchemaValidatorRequest *jsonschema.Schema
	jsonSchemaValidatorRequest, err = jsonschema.CompileString("schema.json", tempJsonSchema)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "b090d10f-4880-45ee-b015-43e8789bc1ea",
			"err": err,
			"requestMessageToTestApiEngineJsonSchema": *requestMessageToTestApiEngineJsonSchema,
		}).Error("Couldn't compile the json-schema for 'requestMessageToTestApiEngineJsonSchema'")
	}

	// Second convert to object that can be validated
	testApiEngineRequestAsJsonByteArray = *testApiEngineRequestAsJson
	var jsonObjectedToBeValidated interface{}
	err = json.Unmarshal(testApiEngineRequestAsJsonByteArray, &jsonObjectedToBeValidated)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                                  "e0c23da2-322c-477b-bcca-7aca9f39aa9b",
			"err":                                 err,
			"string(*testApiEngineRequestAsJson)": string(*testApiEngineRequestAsJson),
		}).Error("Couldn't Unmarshal the json, testApiEngineRequestAsJsonByteArray, into object that can be validated")

		return err
	}

	// Thirds validate that the 'Request' is valid -'requestMessageToTestApiEngineJsonSchema'
	err = jsonSchemaValidatorRequest.Validate(jsonObjectedToBeValidated)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                                  "bb0ffb11-77e6-4739-8911-fcb70ff714f2",
			"err":                                 err,
			"string(*testApiEngineRequestAsJson)": string(*testApiEngineRequestAsJson),
		}).Error("'string(*testApiEngineRequestAsJson)' is not valid to json " +
			"'requestMessageToTestApiEngineJsonSchema'")

		return err

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                                  "18c74f8c-3de2-4333-9ebb-618f5908adcd",
			"string(*testApiEngineRequestAsJson)": string(*testApiEngineRequestAsJson),
		}).Debug("'string(*testApiEngineRequestAsJson)' is valid to json-schema " +
			"'requestMessageToTestApiEngineJsonSchema'")
	}

	// 	// *** Second Step - Validate the MethodParameters***

	tempJsonSchema = *requestMethodParametersMessageToTestApiEngineJsonSchema

	// Compile json-schema 'requestMessageToTestApiEngineJsonSchema'
	var jsonSchemaValidatorMethodParametersRequest *jsonschema.Schema
	jsonSchemaValidatorMethodParametersRequest, err = jsonschema.CompileString("schema.json", tempJsonSchema)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "f519829c-0e68-4824-92d7-abd5cac6033f",
			"err": err,
			"requestMethodParametersMessageToTestApiEngineJsonSchema": *requestMethodParametersMessageToTestApiEngineJsonSchema,
		}).Error("Couldn't compile the json-schema for 'requestMethodParametersMessageToTestApiEngineJsonSchema'")
	}

	// Extract the MethodParameters which is embedded as string and should be converted into a json
	var attributesMapAsJsonAsByteArray []byte

	// Second convert to object that can be validated
	attributesMapAsJsonAsByteArray = *attributesMapAsJson
	err = json.Unmarshal(attributesMapAsJsonAsByteArray, &jsonObjectedToBeValidated)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                           "e6d9ab52-9342-4344-8765-b9856df56427",
			"err":                          err,
			"string(*attributesMapAsJson)": string(*attributesMapAsJson),
		}).Error("Couldn't Unmarshal the json, attributesMapAsJson, into object that can be validated")

		return err
	}

	// Thirds validate that the 'Request' is valid -'requestMessageToTestApiEngineJsonSchema'
	err = jsonSchemaValidatorMethodParametersRequest.Validate(jsonObjectedToBeValidated)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                           "e375ccfe-8228-4e88-a922-c37429e96110",
			"err":                          err,
			"string(*attributesMapAsJson)": string(*attributesMapAsJson),
		}).Error("'string(*attributesMapAsJson)' is not valid to json " +
			"'requestMethodParametersMessageToTestApiEngineJsonSchema'")

		return err

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                           "4293eab6-6c3e-4164-b853-d9c999396acb",
			"string(*attributesMapAsJson)": string(*attributesMapAsJson),
		}).Debug("'string(*attributesMapAsJson)' is valid to json-schema " +
			"'requestMethodParametersMessageToTestApiEngineJsonSchema'")
	}

	return err
}

// Validate the json-response from the Rest-call to TestApiEngine
// Validation os done with supported json-schemas
// First that the overall message is valid
// Second that the Response Variable message is valid
func validateAndTransformRestResponse(
	testApiEngineResponseMessageJson *string,
	testApiEngineResponseMessageJsonSchema *string,
	finalTestInstructionExecutionResultJsonSchema *string,
	responseVariablesJsonSchema *string) (
	testApiEngineFinalTestInstructionExecutionResult TestApiEngineFinalTestInstructionExecutionResultStruct,
	err error) {

	// *** First Step ***

	// Compile json-schema 'testApiEngineResponseMessageJsonSchema'
	var jsonSchemaValidatorTestApiEngineResponseMessage *jsonschema.Schema
	jsonSchemaValidatorTestApiEngineResponseMessage, err = jsonschema.CompileString("schema.json", *testApiEngineResponseMessageJsonSchema)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                                     "2b42469d-7ccc-4f0e-9fec-3e992ba8febb",
			"err":                                    err,
			"testApiEngineResponseMessageJsonSchema": *testApiEngineResponseMessageJsonSchema,
		}).Error("Couldn't compile the json-schema for 'testApiEngineResponseMessageJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	}

	// Clean from '\r\n' into ''. The reason is that all '"' are escaped within the string
	var cleanedTestApiEngineResponseMessageJson string
	cleanedTestApiEngineResponseMessageJson = strings.ReplaceAll(*testApiEngineResponseMessageJson, "\\r\\n", "")

	// Convert to object that can be validated
	var jsonObjectedToBeValidated interface{}
	err = json.Unmarshal([]byte(cleanedTestApiEngineResponseMessageJson), &jsonObjectedToBeValidated)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "5f868efa-a2df-4b7c-92b3-1e47de1b476f",
			"err":                              err,
			"testApiEngineResponseMessageJson": *testApiEngineResponseMessageJson,
		}).Error("Couldn't Unmarshal the json, testApiEngineResponseMessageJson, into object that can be validated")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Validate that the overall message is valid -'testApiEngineResponseMessageJson' with 'testApiEngineResponseMessageJsonSchema'
	err = jsonSchemaValidatorTestApiEngineResponseMessage.Validate(jsonObjectedToBeValidated)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                                     "771b579b-3191-4877-b126-64affafe7fc0",
			"err":                                    err,
			"testApiEngineResponseMessageJsonSchema": *testApiEngineResponseMessageJsonSchema,
		}).Error("'testApiEngineResponseMessageJson' is not valid to json-schema " +
			"'testApiEngineResponseMessageJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "27ff673d-f311-4753-8ae8-ab922ce896c9",
			"testApiEngineResponseMessageJson": *testApiEngineResponseMessageJson,
		}).Debug("'testApiEngineResponseMessageJson' is valid to json-schema " +
			"'testApiEngineResponseMessageJsonSchema'")
	}

	// UmMarshal TestApiEngine-json into Go-struct
	var tempTestApiEngineResponseWithResponseValueAsString TestApiEngineResponseWithResponseValueAsStringStruct
	err = json.Unmarshal([]byte(cleanedTestApiEngineResponseMessageJson),
		&tempTestApiEngineResponseWithResponseValueAsString)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "c91706b3-d462-4943-929e-d424a11a7258",
			"cleanedTestApiEngineResponseMessageJson": cleanedTestApiEngineResponseMessageJson,
		}).Error("Couldn't Unmarshal 'cleanedTestApiEngineResponseMessageJson' into Go-struct of type 'TestApiEngineResponseWithResponseValueAsStringStruct'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Extract Fenix-response, which is 'hidden' in a string TestApiEngine-json
	var fenixMessageInTestApiEngineResponseValueAsJsonString string
	fenixMessageInTestApiEngineResponseValueAsJsonString = tempTestApiEngineResponseWithResponseValueAsString.ResponseValue

	// Clean from '\"' into '"'. The reason is that all '"' are escaped within the string
	var cleanedFenixMessageInTestApiEngineResponseValueAsJsonString string
	cleanedFenixMessageInTestApiEngineResponseValueAsJsonString = strings.ReplaceAll(fenixMessageInTestApiEngineResponseValueAsJsonString, `\"`, `"`)

	// Create a proper json for the string of "responseValue"
	cleanedFenixMessageInTestApiEngineResponseValueAsJsonString = strings.ReplaceAll(cleanedFenixMessageInTestApiEngineResponseValueAsJsonString, `responseValue":"{`, `responseValue":{`)
	cleanedFenixMessageInTestApiEngineResponseValueAsJsonString = strings.ReplaceAll(cleanedFenixMessageInTestApiEngineResponseValueAsJsonString, `}","executionTimeStamp`, `},"executionTimeStamp`)
	cleanedFenixMessageInTestApiEngineResponseValueAsJsonString = strings.ReplaceAll(cleanedFenixMessageInTestApiEngineResponseValueAsJsonString, `\\r\\n`, "")

	// *** Second Step ***

	// Compile json-schema 'finalTestInstructionExecutionResultJsonSchema'
	var jsonSchemaValidatorFinalTestInstructionExecutionResult *jsonschema.Schema
	jsonSchemaValidatorFinalTestInstructionExecutionResult, err = jsonschema.CompileString("schema.json", *finalTestInstructionExecutionResultJsonSchema)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "df414c69-7ee7-454c-83ba-c6350440fd66",
			"err": err,
			"finalTestInstructionExecutionResultJsonSchema": *finalTestInstructionExecutionResultJsonSchema,
		}).Error("Couldn't compile the json-schema for 'finalTestInstructionExecutionResultJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	}

	// Convert to object that can be validated
	err = json.Unmarshal([]byte(cleanedFenixMessageInTestApiEngineResponseValueAsJsonString), &jsonObjectedToBeValidated)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "3e01379c-166b-4c58-a48c-fad7692387da",
			"err": err,
			"cleanedFenixMessageInTestApiEngineResponseValueAsJsonString": cleanedFenixMessageInTestApiEngineResponseValueAsJsonString,
		}).Error("Couldn't Unmarshal the json, cleanedFenixMessageInTestApiEngineResponseValueAsJsonString, into object that can be validated")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Validate that the overall message is valid -'finalTestInstructionExecutionResultJsonSchema'
	err = jsonSchemaValidatorFinalTestInstructionExecutionResult.Validate(jsonObjectedToBeValidated)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "0bd7fcff-ac59-4747-aeaf-f60ef4e0aa37",
			"err": err,
			"cleanedFenixMessageInTestApiEngineResponseValueAsJsonString": cleanedFenixMessageInTestApiEngineResponseValueAsJsonString,
			"finalTestInstructionExecutionResultJsonSchema":               *finalTestInstructionExecutionResultJsonSchema,
		}).Error("'cleanedFenixMessageInTestApiEngineResponseValueAsJsonString' is not valid to json-schema " +
			"'finalTestInstructionExecutionResultJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "2a2cb32d-84d7-4a7f-aced-2b1466982513",
			"err": err,
			"cleanedFenixMessageInTestApiEngineResponseValueAsJsonString": cleanedFenixMessageInTestApiEngineResponseValueAsJsonString,
		}).Debug("'cleanedFenixMessageInTestApiEngineResponseValueAsJsonString' is valid to json-schema " +
			"'finalTestInstructionExecutionResultJsonSchema'")
	}

	// 	 *** Third Step ***

	// UmMarshal 'cleanedFenixMessageInTestApiEngineResponseValueAsJsonString' into Go-struct
	err = json.Unmarshal([]byte(cleanedFenixMessageInTestApiEngineResponseValueAsJsonString),
		&testApiEngineFinalTestInstructionExecutionResult)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "b25bcf70-3740-490a-ba68-5483cfa94427",
			"cleanedFenixMessageInTestApiEngineResponseValueAsJsonString": cleanedFenixMessageInTestApiEngineResponseValueAsJsonString,
		}).Error("Couldn't Unmarshal 'cleanedFenixMessageInTestApiEngineResponseValueAsJsonString' into Go-struct of type 'TestApiEngineFinalTestInstructionExecutionResultStruct'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Compile json-schema 'responseVariablesJsonSchema'
	var jsonSchemaValidatorResponseVariables *jsonschema.Schema
	jsonSchemaValidatorResponseVariables, err = jsonschema.CompileString("schema.json", *responseVariablesJsonSchema)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                          "955f13ca-f450-44ed-9d17-ca8bfb3f601c",
			"err":                         err,
			"responseVariablesJsonSchema": *responseVariablesJsonSchema,
		}).Fatal("Couldn't compile the json-schema for 'responseVariablesJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Extract Response Variables
	var testAPiEngineResponseVariablesType1 []ResponseVariableType1Struct
	var testAPiEngineNoResponseVariables []NoResponseVariableStruct
	var responseVariablesAsJsonByteArray []byte

	for _, responseVariable := range testApiEngineFinalTestInstructionExecutionResult.ResponseVariables {
		switch v := responseVariable.(type) {
		case ResponseVariableType1Struct:

			// Check if the ResponseVariable can be cast into correct response struct
			if resVar, ok := responseVariable.(ResponseVariableType1Struct); ok {
				testAPiEngineResponseVariablesType1 = append(testAPiEngineResponseVariablesType1, resVar)
			} else {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id":               "10c35e90-fd04-47fd-a0da-2188c0bf1f4a",
					"resVar":           resVar,
					"responseVariable": responseVariable,
				}).Fatal("ResponseVariable can't be cast into 'ResponseVariableType1Struct'")
			}

		case NoResponseVariableStruct:
			// Check if the ResponseVariable can be cast into correct response struct
			if resVar, ok := responseVariable.(NoResponseVariableStruct); ok {
				testAPiEngineNoResponseVariables = append(testAPiEngineNoResponseVariables, resVar)
			} else {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id":               "9650a892-7ec4-4329-a9ea-fc3fd1f8fe94",
					"resVar":           resVar,
					"responseVariable": responseVariable,
				}).Fatal("ResponseVariable can't be cast into 'NoResponseVariableStruct'")
			}

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                      "c9356f5b-2568-42ff-b651-032cc59b3aeb",
				"responseVariable.(type)": v,
			}).Fatal("Unknown response variable type")
		}
	}

	// Validate that there are at least one response variable
	if len(testApiEngineFinalTestInstructionExecutionResult.ResponseVariables) == 0 {

		err = errors.New("expected at least one response variable")

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "6ede6867-29eb-4a9d-ac71-e0823de9aee1",
			"testApiEngineFinalTestInstructionExecutionResult.ResponseVariables": testApiEngineFinalTestInstructionExecutionResult.ResponseVariables,
		}).Error("Expected at least one response variable")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	}

	// Select ResponseVariable-type
	switch v := testApiEngineFinalTestInstructionExecutionResult.ResponseVariables[0].(type) {
	case ResponseVariableType1Struct:
		// Convert Response Variables into json-byte-array
		responseVariablesAsJsonByteArray, err = json.Marshal(testAPiEngineResponseVariablesType1)

	case NoResponseVariableStruct:
		// Convert Response Variables into json-byte-array
		responseVariablesAsJsonByteArray, err = json.Marshal(testAPiEngineNoResponseVariables)

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                      "d7e11066-c1cb-4663-ad7b-f8357455bf5f",
			"responseVariable.(type)": v,
		}).Fatal("Unknown response variable type")
	}

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "41d96c5a-25f0-451d-a869-b2686a91f442",
			"string(responseVariablesAsJsonByteArray)": string(responseVariablesAsJsonByteArray),
		}).Error("Couldn't Marshal testAPiEngineResponseVariables into json")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Convert to object that can be validated
	err = json.Unmarshal(responseVariablesAsJsonByteArray, &jsonObjectedToBeValidated)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "f3604e47-9161-4a38-9fcf-ad078cc22860",
			"err":                              err,
			"responseVariablesAsJsonByteArray": string(responseVariablesAsJsonByteArray),
		}).Error("Couldn't Unmarshal the json, responseVariablesAsJsonByteArray, into object that can be validated")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Second validate that the 'Response Variables' is valid -'responseVariablesJsonSchema'
	err = jsonSchemaValidatorResponseVariables.Validate(jsonObjectedToBeValidated)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "74916f92-fb0f-445c-ae0a-6a5bd2bd0fcf",
			"err":                              err,
			"responseVariablesAsJsonByteArray": string(responseVariablesAsJsonByteArray),
			"responseVariablesJsonSchema":      responseVariablesJsonSchema,
		}).Error("'testAPiEngineResponseVariables' is not valid to json-schema " +
			"'responseVariablesJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                             "e498ecb2-8d50-4d99-8984-3171dbd9ed6c",
			"testAPiEngineResponseVariables": string(responseVariablesAsJsonByteArray),
		}).Debug("'testAPiEngineResponseVariables' is valid to json-schema " +
			"'responseVariablesJsonSchema'")
	}

	return testApiEngineFinalTestInstructionExecutionResult, err
}
