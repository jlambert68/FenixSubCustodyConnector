package executeTestInstructionUsingTestApiEngine

import (
	"FenixSubCustodyConnector/sharedCode"
	"bytes"
	"crypto/tls"
	"encoding/json"
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

	// Add TestCaseExecutionUuid, TestInstructionExecutionUuid, TestInstructionExecutionVersion and TestInstructionVersion
	attributesMap["TestInstructionVersion"] = testInstructionVersion
	attributesMap["TestCaseExecutionUuid"] = testApiEngineRestApiMessageValues.TestCaseExecutionUuid
	attributesMap["TestInstructionExecutionUuid"] = testApiEngineRestApiMessageValues.TestInstructionExecutionUuid
	attributesMap["TestInstructionExecutionVersion"] = strconv.Itoa(int(testApiEngineRestApiMessageValues.
		TestInstructionExecutionVersion))
	attributesMap["ExpectedToBePassed"] = string(testApiEngineRestApiMessageValues.TestApiEngineExpectedToBePassedValue)

	// Request to be sent to TestApiEngine
	var tempTestApiEngineRequest testApiEngineRequestStruct
	tempTestApiEngineRequest = testApiEngineRequestStruct{
		TestStepClassName:     string(testApiEngineRestApiMessageValues.TestApiEngineClassNameNAME),
		TestStepActionMethod:  string(testApiEngineRestApiMessageValues.TestApiEngineMethodNameNAME),
		TestDataParameterType: "FixedValue",
		ExpectedToBePassed:    true,
		MethodParameters:      attributesMap,
	}

	// Marshal to json as byte-array
	var attributesAsJson []byte
	attributesAsJson, err = json.Marshal(tempTestApiEngineRequest)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":            "e1e74131-5040-43fa-abfc-1023f09d4388",
			"attributesMap": attributesMap,
		}).Error("Couldn't Marshal Attributes-map into json request")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Validate rest-request and convert
	err = validateRestRequest(
		&attributesAsJson,
		requestMessageToTestApiEngineJsonSchema)

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
	testApiEngineUrl = LocalExecutionMethods.TestApiEngineUrlPath + "/" + string(testApiEngineRestApiMessageValues.TestApiEngineClassNameNAME) +
		"/" + string(testApiEngineRestApiMessageValues.TestApiEngineMethodNameNAME) +
		"?" + "expectedToBePassed=" + string(testApiEngineRestApiMessageValues.TestApiEngineExpectedToBePassedValue)

	// Use Local web server for test or TestApiEngine
	if UseInternalWebServerForTestInsteadOfCallingTestApiEngine == true {
		// Use Local web server for testing
		testApiEngineUrl = "http://" + LocalWebServerPort + testApiEngineUrl

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "724d84e8-ec94-4947-ac74-0c7e5c17cfb6",
		}).Debug("Posting TestInstruction to local web server for Tests")

	} else {
		// Use TestApiEngine address and port
		testApiEngineUrl = "https://" + LocalExecutionMethods.TestApiEngineAddress + ":" + LocalExecutionMethods.TestApiEnginePort + testApiEngineUrl

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
		bytes.NewBuffer(attributesAsJson))
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":               "12b846ad-e8bf-41e0-8893-a1a7cef5f396",
			"testApiEngineUrl": testApiEngineUrl,
		}).Error("Couldn't do call to Rest-execution-server")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Read the body
	restResponse.Body.Close()
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
	attributesAsJson *[]byte,
	requestMessageToTestApiEngineJsonSchema *string) (
	err error) {

	var tempJsonSchema string
	var tempAttributesAsByteArray []byte

	// 	// *** First Step ***
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
	tempAttributesAsByteArray = *attributesAsJson
	var jsonObjectedToBeValidated interface{}
	err = json.Unmarshal(tempAttributesAsByteArray, &jsonObjectedToBeValidated)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                        "e0c23da2-322c-477b-bcca-7aca9f39aa9b",
			"err":                       err,
			"string(*attributesAsJson)": string(*attributesAsJson),
		}).Error("Couldn't Unmarshal the json, tempAttributesAsByteArray, into object that can be validated")

		return err
	}

	// Thirds validate that the 'Request' is valid -'requestMessageToTestApiEngineJsonSchema'
	err = jsonSchemaValidatorRequest.Validate(jsonObjectedToBeValidated)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                        "bb0ffb11-77e6-4739-8911-fcb70ff714f2",
			"err":                       err,
			"string(*attributesAsJson)": string(*attributesAsJson),
		}).Error("'string(*attributesAsJson)' is not valid to json " +
			"'requestMessageToTestApiEngineJsonSchema'")

		return err

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                        "18c74f8c-3de2-4333-9ebb-618f5908adcd",
			"string(*attributesAsJson)": string(*attributesAsJson),
		}).Debug("'string(*attributesAsJson)' is valid to json-schema " +
			"'requestMessageToTestApiEngineJsonSchema'")
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

	// Convert to object that can be validated
	var jsonObjectedToBeValidated interface{}
	err = json.Unmarshal([]byte(*testApiEngineResponseMessageJson), &jsonObjectedToBeValidated)

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
			"id":                               "2a2cb32d-84d7-4a7f-aced-2b1466982513",
			"testApiEngineResponseMessageJson": *testApiEngineResponseMessageJson,
		}).Debug("'testApiEngineResponseMessageJson' is valid to json-schema " +
			"'testApiEngineResponseMessageJsonSchema'")
	}

	// UmMarshal TestApiEngine-json into Go-struct
	var tempTestApiEngineResponse TestApiEngineResponseStruct
	err = json.Unmarshal([]byte(*testApiEngineResponseMessageJson),
		&tempTestApiEngineResponse)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "18b982b4-cb71-4309-83fd-5af5a30cf456",
			"testApiEngineResponseMessageJson": *testApiEngineResponseMessageJson,
		}).Error("Couldn't Unmarshal 'testApiEngineResponseMessageJson' into Go-struct")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Extract Fenix-response, which is 'hidden' in a string TestApiEngine-json
	var cleanedResponseVariablesAsString string

	// Clean from '\"' into '"'. The reason is that all '"' are escaped within the string
	cleanedResponseVariablesAsString = strings.ReplaceAll(tempTestApiEngineResponse.ResponseVariables, `\\\"`, `"`)

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
	err = json.Unmarshal([]byte(cleanedResponseVariablesAsString), &jsonObjectedToBeValidated)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "3e01379c-166b-4c58-a48c-fad7692387da",
			"err":                              err,
			"cleanedResponseVariablesAsString": cleanedResponseVariablesAsString,
		}).Error("Couldn't Unmarshal the json, cleanedResponseVariablesAsString, into object that can be validated")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Validate that the overall message is valid -'finalTestInstructionExecutionResultJsonSchema'
	err = jsonSchemaValidatorFinalTestInstructionExecutionResult.Validate(jsonObjectedToBeValidated)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "0bd7fcff-ac59-4747-aeaf-f60ef4e0aa37",
			"err":                              err,
			"cleanedResponseVariablesAsString": cleanedResponseVariablesAsString,
		}).Error("'cleanedResponseVariablesAsString' is not valid to json-schema " +
			"'finalTestInstructionExecutionResultJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "2a2cb32d-84d7-4a7f-aced-2b1466982513",
			"cleanedResponseVariablesAsString": cleanedResponseVariablesAsString,
		}).Debug("'cleanedResponseVariablesAsString' is valid to json-schema " +
			"'finalTestInstructionExecutionResultJsonSchema'")
	}

	// UmMarshal TestApiEngine-json into Go-struct
	err = json.Unmarshal([]byte(cleanedResponseVariablesAsString),
		&testApiEngineFinalTestInstructionExecutionResult)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "4997b271-fcf0-44fa-ac29-cdab53f7cdbb",
			"cleanedResponseVariablesAsString": cleanedResponseVariablesAsString,
		}).Error("Couldn't Unmarshal 'cleanedResponseVariablesAsString' into Go-struct")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Extract Response Variables
	var testAPiEngineResponseVariables []ResponseVariableStruct
	testAPiEngineResponseVariables = testApiEngineFinalTestInstructionExecutionResult.ResponseVariables

	// Convert 'Response Variables' into json, to be validated towards json-schema
	var testAPiEngineResponseVariablesAsJsonByteArray []byte
	testAPiEngineResponseVariablesAsJsonByteArray, err = json.Marshal(testAPiEngineResponseVariables)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                             "4997b271-fcf0-44fa-ac29-cdab53f7cdb",
			"err":                            err,
			"testAPiEngineResponseVariables": testAPiEngineResponseVariables,
		}).Error("Couldn't marshal 'testAPiEngineResponseVariables' into json")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	}

	// 	// *** Third Step ***

	// Compile json-schema 'responseVariablesJsonSchema'
	var jsonSchemaValidatorResponseVariables *jsonschema.Schema
	jsonSchemaValidatorResponseVariables, err = jsonschema.CompileString("schema.json", *responseVariablesJsonSchema)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                          "a852d237-13e0-4925-893d-c185c316ed17",
			"err":                         err,
			"responseVariablesJsonSchema": *responseVariablesJsonSchema,
		}).Fatal("Couldn't compile the json-schema for 'responseVariablesJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Convert to object that can be validated
	err = json.Unmarshal(testAPiEngineResponseVariablesAsJsonByteArray, &jsonObjectedToBeValidated)

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                             "ae00dc85-24ee-4b8f-b300-ba17d31dba6f",
			"err":                            err,
			"testAPiEngineResponseVariables": testAPiEngineResponseVariables,
		}).Error("Couldn't Unmarshal the json, testAPiEngineResponseVariablesAsJsonByteArray, into object that can be validated")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err
	}

	// Second validate that the 'Response Variables' is valid -'responseVariablesJsonSchema'
	err = jsonSchemaValidatorResponseVariables.Validate(
		strings.NewReader(string(testAPiEngineResponseVariablesAsJsonByteArray)))
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "74916f92-fb0f-445c-ae0a-6a5bd2bd0fcf",
			"err": err,
			"string(testAPiEngineResponseVariablesAsJsonByteArray)": string(testAPiEngineResponseVariablesAsJsonByteArray),
		}).Error("'string(testAPiEngineResponseVariablesAsJsonByteArray)' is not valid to json-schema " +
			"'responseVariablesJsonSchema'")

		return TestApiEngineFinalTestInstructionExecutionResultStruct{}, err

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "e498ecb2-8d50-4d99-8984-3171dbd9ed6c",
			"string(testAPiEngineResponseVariablesAsJsonByteArray)": string(testAPiEngineResponseVariablesAsJsonByteArray),
		}).Debug("'string(testAPiEngineResponseVariablesAsJsonByteArray)' is valid to json-schema " +
			"'responseVariablesJsonSchema'")
	}

	return testApiEngineFinalTestInstructionExecutionResult, err
}
