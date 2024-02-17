package executeTestInstructionUsingTestApiEngine

import (
	"FenixSubCustodyConnector/sharedCode"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

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
	if UseInternalWebServerForTestInsteadOfCallingTestApiEngine == true {
		// Use Local web server for testing
		TestApiEngineUrl = "http://:" + LocalWebServerPort + TestApiEngineUrl

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "724d84e8-ec94-4947-ac74-0c7e5c17cfb6",
		}).Debug("Posting TestInstruction to local web server for Tests")

	} else {
		// Use TestApiEngine address and port
		TestApiEngineUrl = "http://:" + LocalExecutionMethods.TestApiEngineAddress + ":" + LocalExecutionMethods.TestApiEnginePort + TestApiEngineUrl

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

// Validate the json-response from the Rest-call to TestApiEngine
// Validation os done with supported json-schemas
// First that the overall message is valid
// Second that the Response Variable message is valid
func validateRestResponse(
	finalTestInstructionExecutionResultAsJson *string,
	finalTestInstructionExecutionResultJsonSchema *string,
	responseVariablesJsonSchema *string) (
	err error) {

	// Load the schema - 'finalTestInstructionExecutionResultJsonSchema'
	var jsonSchemaCompilerFinalTestInstructionExecutionResult *jsonschema.Compiler
	jsonSchemaCompilerFinalTestInstructionExecutionResult = jsonschema.NewCompiler()
	err = jsonSchemaCompilerFinalTestInstructionExecutionResult.AddResource("schema.json",
		strings.NewReader(*finalTestInstructionExecutionResultJsonSchema))
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "cf1f344a-ec0c-464f-9082-24823a06540a",
			"err": err,
			"finalTestInstructionExecutionResultJsonSchema": *finalTestInstructionExecutionResultJsonSchema,
		}).Fatal("Couldn't add json-schema for 'finalTestInstructionExecutionResultJsonSchema' to " +
			"'json-schema compile")

	}
	// Compile json-schema 'finalTestInstructionExecutionResultJsonSchema'
	var jsonSchemaValidatorFinalTestInstructionExecutionResult *jsonschema.Schema
	jsonSchemaValidatorFinalTestInstructionExecutionResult, err =
		jsonSchemaCompilerFinalTestInstructionExecutionResult.Compile("schema.json")
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "eaadce84-0800-4ecc-8a67-29351146060a",
			"err": err,
			"finalTestInstructionExecutionResultJsonSchema": *finalTestInstructionExecutionResultJsonSchema,
		}).Fatal("Couldn't compile the json-schema for 'finalTestInstructionExecutionResultJsonSchema'")
	}

	// First validate that the overall message is valid -'finalTestInstructionExecutionResultJsonSchema'
	err = jsonSchemaValidatorFinalTestInstructionExecutionResult.Validate(
		strings.NewReader(*finalTestInstructionExecutionResultAsJson))
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "d71db97a-740a-4a6f-a429-568e2739496a",
			"err": err,
			"finalTestInstructionExecutionResultAsJson": *finalTestInstructionExecutionResultAsJson,
		}).Error("'finalTestInstructionExecutionResultAsJson' is not valid to json-schema " +
			"'finalTestInstructionExecutionResultJsonSchema'")
	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "f2189bcd-5822-4660-aa9b-e853fd2765e9",
			"finalTestInstructionExecutionResultAsJson": *finalTestInstructionExecutionResultAsJson,
		}).Error("'finalTestInstructionExecutionResultAsJson' is valid to json-schema " +
			"'finalTestInstructionExecutionResultJsonSchema'")
	}

	// UmMarshal TestApiEngine-json into Go-struct
	var testApiEngineFinalTestInstructionExecutionResult TestApiEngineFinalTestInstructionExecutionResultStruct
	err = json.Unmarshal([]byte(*finalTestInstructionExecutionResultAsJson),
		&testApiEngineFinalTestInstructionExecutionResult)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "4997b271-fcf0-44fa-ac29-cdab53f7cdbb",
			"finalTestInstructionExecutionResultAsJson": *finalTestInstructionExecutionResultAsJson,
		}).Fatal("Couldn't Unmarshal 'finalTestInstructionExecutionResultAsJson' into Go-struct")
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
		}).Fatal("Couldn't marshal 'testAPiEngineResponseVariables' into json")

	}

	d
	// Load the schema - 'responseVariablesJsonSchema'
	var jsonSchemaCompilerResponseVariables *jsonschema.Compiler
	jsonSchemaCompilerResponseVariables = jsonschema.NewCompiler()
	err = jsonSchemaCompilerResponseVariables.AddResource("schema.json",
		strings.NewReader(*responseVariablesJsonSchema))
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                          "02c10d0e-5fd1-4d45-9d4a-bc041cb4e04d",
			"err":                         err,
			"responseVariablesJsonSchema": *responseVariablesJsonSchema,
		}).Fatal("Couldn't add json-schema for 'responseVariablesJsonSchema' to " +
			"'json-schema compile")

	}
	// Compile json-schema 'responseVariablesJsonSchema'
	var jsonSchemaValidatorResponseVariables *jsonschema.Schema
	jsonSchemaValidatorResponseVariables, err = jsonSchemaCompilerResponseVariables.
		Compile("schema.json")
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                          "a852d237-13e0-4925-893d-c185c316ed17",
			"err":                         err,
			"responseVariablesJsonSchema": *responseVariablesJsonSchema,
		}).Fatal("Couldn't compile the json-schema for 'responseVariablesJsonSchema'")
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
	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "e498ecb2-8d50-4d99-8984-3171dbd9ed6c",
			"string(testAPiEngineResponseVariablesAsJsonByteArray)": string(testAPiEngineResponseVariablesAsJsonByteArray),
		}).Error("'string(testAPiEngineResponseVariablesAsJsonByteArray)' is valid to json-schema " +
			"'responseVariablesJsonSchema'")
	}

	return err
}
