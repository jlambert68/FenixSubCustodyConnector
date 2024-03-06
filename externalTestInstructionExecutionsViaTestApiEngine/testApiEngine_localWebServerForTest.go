package executeTestInstructionUsingTestApiEngine

import (
	"FenixSubCustodyConnector/sharedCode"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"strconv"
	"time"

	//"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Structs used when converting json messages in RestAPI

// RestUserMessageStruct
type RestUserMessageStruct struct {
	ParameterMap map[string]string `json:"ParameterMap,omitempty"`
}

func RestAPIServer() {

	// Wait until Logger has been initiated
	for {
		if sharedCode.Logger == nil {
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "028012db-71a4-4585-900b-4b5986f7a4bc",
	}).Info("starting API server for Test")

	//create a new router
	router := mux.NewRouter()
	router.UseEncodedPath()

	//specify endpoints
	router.HandleFunc("/health-check", healthCheck).Methods("GET")
	router.HandleFunc("/TestCaseExecution/ExecuteTestAction", testApiEngineClassTestApiEngineMethod).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(notFound)
	//router.HandleFunc("/*", allOtherRoutes).Methods("POST")
	/*
		router.HandleFunc("/ExampleTestStepClass/DoSomething1{expectedToBePassed}", doSomething1).Methods("POST")
		router.HandleFunc("/ExampleTestStepClass/DoSomething2{expectedToBePassed}", doSomething2).Methods("POST")
		router.HandleFunc("/ExampleTestStepClass/DoSomethingWithTestException{expectedToBePassed}", doSomethingWithTestException).Methods("POST")
		router.HandleFunc("/ExampleTestStepClass/DoSomethingWithException{expectedToBePassed}", doSomethingWithException).Methods("POST")
	*/

	http.Handle("/", router)

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                 "4961d4f2-20f5-43be-8a9b-074ec79f5075",
		"LocalWebServerPort": LocalWebServerPort,
	}).Debug("Starting Local Web Server for Test, instead of doing calls to TestApiEngine")

	//start and listen to requests
	http.ListenAndServe(":"+LocalWebServerPort, router)
}

// RestApi to check if local TestWeb-server is up and running
func notFound(w http.ResponseWriter, r *http.Request) {
	// curl --request GET localhost:8080/health-check

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":           "42c2cdca-4ce1-4802-888d-ccc6eb82996f",
		"http.Request": r,
	}).Debug("Incoming 'RestApi - *'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "fab7676d-c303-4b20-8980-397d7a59282e",
	}).Debug("Outgoing 'RestApi - *'")

	// Create base for response body
	var responseBody map[string]string
	responseBody = make(map[string]string)
	responseBody["type"] = "FenixConnector - internal Web Server"

	// Create Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	// Create Response message
	responseBody["title"] = "Error - Not Found"
	responseBody["status"] = "404"
	responseBody["detail"] = "The address used is not valid"
	responseBody["traceId"] = "6d7c074e-2110-49ef-a45a-2a41a5a83b2c"

	responseBodydata, _ := json.Marshal(responseBody)

	fmt.Fprintf(w, string(responseBodydata))

	return

}

// RestApi to check if local TestWeb-server is up and running
func healthCheck(w http.ResponseWriter, r *http.Request) {
	// curl --request GET localhost:8080/health-check

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "fb3c1ecb-3da8-4d27-b1c4-16d5120e7125",
	}).Debug("Incoming 'RestApi - /health-check'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "fab7676d-c303-4b20-8980-397d7a59282e",
	}).Debug("Outgoing 'RestApi - /health-check'")

	// Create base for response body
	var responseBody map[string]string
	responseBody = make(map[string]string)
	responseBody["type"] = "FenixSCConnector - internal Web Server"

	// Create Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Create Response message
	responseBody["title"] = "OK"
	responseBody["status"] = "200"
	responseBody["detail"] = "The Connectors Internal Test API is up and running"
	responseBody["traceId"] = "5a263f5c-8a86-4552-bc61-3b5cce52b208"

	responseBodydata, _ := json.Marshal(responseBody)

	fmt.Fprintf(w, string(responseBodydata))

	return
	// Create Response message
	fmt.Fprintf(w, "API is up and running")
}

func doSomething(w http.ResponseWriter, r *http.Request) {

	// curl -X POST localhost:8080/ExampleTestStepClass/DoSomething?expectedToBePassed=true -H 'Content-Type: application/json' -d '{"UserId":"s41797", "TestInstructionUuid":"myUuid", "TestInstructionName":"myName"}'

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "2472dda1-701d-4b23-8326-757e43df4af4",
	}).Debug("Incoming 'RestApi - (POST) /DoSomething")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "db318ff4-ad36-43d4-a8d4-3e0ac4ff08c6",
	}).Debug("Outgoing 'RestApi - (POST) /DoSomething'")

	// Variable where Rest-json-payload will end up in
	//jsonData := &RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct{}

	// Create base for response body
	var responseBody map[string]string
	responseBody = make(map[string]string)
	responseBody["type"] = "FenixSCConnector - internal Web Server"

	// read message body
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		fmt.Println(error)
		return
	}

	// close message body
	r.Body.Close()

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(body, &jsonMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract 'expectedToBePassedSlice'
	variables := r.URL.Query() //mux.Vars(r)
	expectedToBePassedSlice, existInMap := variables["expectedToBePassed"]

	// Missing parameter 'expectedToBePassedSlice'
	if existInMap == false {

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		// Create Response message
		responseBody["title"] = "Error - Bad Request"
		responseBody["status"] = "400"
		responseBody["detail"] = "Missing parameter 'expectedToBePassed'"
		responseBody["traceId"] = "15f7f628-c80e-4010-8853-66df1ffa1a59"

		responseBodydata, _ := json.Marshal(responseBody)

		fmt.Fprintf(w, string(responseBodydata))

		return
	}

	// Exact one parameter 'expectedToBePassed' must exist
	if len(expectedToBePassedSlice) != 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		// Create Response message
		fmt.Fprintf(w, "Parameter 'expectedToBePassed' must contain exactly one value")

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		// Create Response message
		responseBody["title"] = "Error - Bad Request"
		responseBody["status"] = "400"
		responseBody["detail"] = "Parameter 'expectedToBePassed' must contain exactly one value"
		responseBody["traceId"] = "dcdfc951-1eb5-4ed9-8c54-5f22bb718ae7"

		responseBodydata, _ := json.Marshal(responseBody)

		fmt.Fprintf(w, string(responseBodydata))

		return
	}

	// Parameter 'expectedToBePassed' should be 'true' or 'false'
	if expectedToBePassedSlice[0] != "true" && expectedToBePassedSlice[0] != "false" {

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		// Create Response message
		responseBody["title"] = "Error - Bad Request"
		responseBody["status"] = "400"
		responseBody["detail"] = "Parameter 'expectedToBePassed' should be 'true' or 'false'"
		responseBody["traceId"] = "2c82ed7f-18f6-4362-8ca7-a4c3602d81ac"

		responseBodydata, _ := json.Marshal(responseBody)

		fmt.Fprintf(w, string(responseBodydata))

		return
	}

	// 'expectedToBePassed' should be 'true'
	if expectedToBePassedSlice[0] == "true" {

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Create Response message
		responseBody["title"] = "OK"
		responseBody["status"] = "200"
		responseBody["detail"] = "OK Test from Test Web server"
		responseBody["traceId"] = "8f374286-d692-4196-83b4-575f66c12684"

		responseBodydata, _ := json.Marshal(responseBody)

		fmt.Fprintf(w, string(responseBodydata))

		return

	}

	// 'expectedToBePassed' is 'false' - Will allways go in here
	if expectedToBePassedSlice[0] == "false" {

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) //TODO Fang must change this

		// Create Response message
		fmt.Fprintf(w, "Not a OK Test from Test Web server")
		// Create Response message
		responseBody["title"] = "Error - Internal Server Error"
		responseBody["status"] = "500"
		responseBody["detail"] = "Not a OK Test from Test Web server"
		responseBody["traceId"] = "7f139cbd-2fb2-4ba2-9f8b-4d42faefc69f"

		responseBodydata, _ := json.Marshal(responseBody)

		fmt.Fprintf(w, string(responseBodydata))

	}
}

func testApiEngineClassTestApiEngineMethod(w http.ResponseWriter, r *http.Request) {

	// curl -X POST localhost:8080/GeneralSetupTearDown/SetupexpectedToBePassed=true -H 'Content-Type: application/json' -d '{"UserId":"s41797", "TestInstructionUuid":"myUuid", "TestInstructionName":"myName"}'

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "5c68e681-73c1-438f-aed5-cf7f6f4f9072",
	}).Debug("Incoming 'RestApi - (POST) /TestApiEngineClass/TestApiEngineMethod'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "9195d621-eb4a-477f-8f68-1109c4aa69c1",
	}).Debug("Outgoing 'RestApi - (POST) /TestApiEngineClass/TestApiEngineMethod'")

	// StartTime for TestInstructionExecution
	var tempTestInstructionExecutionStartTimeStamp time.Time
	tempTestInstructionExecutionStartTimeStamp = time.Now()

	// Variable where Rest-json-payload will end up in
	//jsonData := &RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct{}

	// Create base for response body
	var responseBody map[string]string
	responseBody = make(map[string]string)
	responseBody["type"] = "FenixConnector - internal Web Server"

	// read message body
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		fmt.Println(error)
		return
	}

	// close message body
	r.Body.Close()

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(body, &jsonMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		// Just print Incoming parameter
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":      "12789a24-f02c-494d-b8d8-18642d7588ef",
			"jsonMap": jsonMap,
		}).Debug("Incoming Parameters")
	}

	// Variables to be extracted from jsonMap
	var (
		testStepClassName              string
		testStepActionMethod           string
		testDataParameterType          string
		expectedToBePassedTestApiLevel bool
		methodParameter                map[string]interface{}
	)

	var (
		expectedToBePassedAsString              string
		testCaseExecutionUuid                   string
		testInstructionExecutionUuid            string
		testInstructionExecutionVersionAsString string
		testInstructionVersion                  string
		timeoutTimeInSecondsAsString            string
		timeoutTimeInSeconds                    int
		relatedReference_54x_20CRELA            string
	)

	var existInMap bool
	var tempMapVariable interface{}
	var canCastTempMapVariableToCorrectVariableType bool
	var variableName string
	var variableType string

	// *** Extract TestStepClassName ****
	variableName = "testStepClassName"
	variableType = "string"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = jsonMap[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "e3d3c6f9-1374-4ca6-9625-c5bc34f8265f",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'",
			variableName))
	}

	// Transform variable into correct type
	testStepClassName, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "e3d3c6f9-1374-4ca6-9625-c5bc34f8265f",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// *** Extract TestStepActionMethod ****
	variableName = "testStepActionMethod"
	variableType = "string"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = jsonMap[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "e3d3c6f9-1374-4ca6-9625-c5bc34f8265f",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'",
			variableName))

	}

	// Transform variable into correct type
	testStepActionMethod, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "e3d3c6f9-1374-4ca6-9625-c5bc34f8265f",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// *** Extract testDataParameterType ****
	variableName = "testDataParameterType"
	variableType = "string"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = jsonMap[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "7867e66a-19af-4217-8c8f-8f3cb6f88c1b",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

	}

	// Transform variable into correct type
	testDataParameterType, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "620f1e40-9ac4-47b4-8547-25584da0ad71",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// Verify that value is set to 'FixedValue'
	if testDataParameterType != "FixedValue" {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "a11a2549-5c16-45e8-bf04-0b0d6a233d65",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' doesn't have value 'FixedValue'",
			variableName))

	}

	// *** Extract expectedToBePassedTestApiLevel ****
	variableName = "expectedToBePassed"
	variableType = "boolean"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = jsonMap[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "566b3de7-ec40-4348-8d3b-39ef92e7f828",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))
	}

	// Transform variable into correct type
	expectedToBePassedTestApiLevel, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(bool)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "620f1e40-9ac4-47b4-8547-25584da0ad71",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// Verify that value is set to 'true'
	if expectedToBePassedTestApiLevel != true {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "a11a2549-5c16-45e8-bf04-0b0d6a233d65",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' doesn't have value 'true'",
			variableName))

	}

	// *** Extract methodParameter ****
	variableName = "methodParameters"
	variableType = "map[string]interface{}"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = jsonMap[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "d1e918b7-9e42-4020-a657-ae56df7e8d78",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

	}

	// Transform variable into correct type
	methodParameter, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(map[string]interface{})
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "806c60d4-900f-4e9f-9cbb-f91fb830fa73",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// *** Extract ExpectedToBePassed ****
	variableName = "ExpectedToBePassed"
	variableType = "string"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = methodParameter[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "0db1bff6-7583-4b3d-a9f3-6a498b4c65dc",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

	}

	// Transform variable into correct type
	expectedToBePassedAsString, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "574dd3fb-a0ff-4a11-aa43-34f2aa7477fa",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// Validate that value is a boolean
	if expectedToBePassedAsString != "true" && expectedToBePassedAsString != "false" {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "138d84de-e1aa-4ab7-9300-4d30b7d3debf",
		}).Fatalln(fmt.Sprintf("parameter '%s' is not any of 'true' or 'false'. The value is '%s'",
			variableName, expectedToBePassedAsString))

	}

	// *** Extract TestCaseExecutionUuid ****
	variableName = "TestCaseExecutionUuid"
	variableType = "string"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = methodParameter[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "0db1bff6-7583-4b3d-a9f3-6a498b4c65dc",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

	}

	// Transform variable into correct type
	testCaseExecutionUuid, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "d343b69f-75b7-4c74-b337-eeada05f7289",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))
	}

	// *** Extract TestInstructionExecutionUuid ****
	variableName = "TestInstructionExecutionUuid"
	variableType = "string"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = methodParameter[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "0db1bff6-7583-4b3d-a9f3-6a498b4c65dc",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

	}

	// Transform variable into correct type
	testInstructionExecutionUuid, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "14c93b8f-f9bd-42dd-ba72-30b850fba39b",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// *** Extract TestInstructionExecutionVersion ****
	variableName = "TestInstructionExecutionVersion"
	variableType = "Integer"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = methodParameter[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "e20def48-50a4-4a54-adb1-68ae23302e26",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

	}

	// Transform variable into correct type
	testInstructionExecutionVersionAsString, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "976b32b9-1f75-4d46-9ec2-0bd0c5cf983f",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// Validate that value is an integer
	_, err = strconv.Atoi(testInstructionExecutionVersionAsString)
	if err != nil {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":  "d627108c-55de-42b9-a291-a6c28efd1e5c",
			"err": err,
			"testInstructionExecutionVersionAsString": testInstructionExecutionVersionAsString,
		}).Fatalln(fmt.Sprintf("Couldn't convert 'testInstructionExecutionVersionAsString' into an integer",
			variableName, expectedToBePassedAsString))

	}

	// *** Extract TestInstructionVersion ****
	variableName = "TestInstructionVersion"
	variableType = "String"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = methodParameter[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "e20def48-50a4-4a54-adb1-68ae23302e26",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

	}

	// Transform variable into correct type
	testInstructionVersion, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "87c1a80c-f5da-47a1-8d25-2c232ee7f770",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
			variableName,
			variableType))

	}

	// *** Extract TimeoutTimeInSeconds ****
	variableName = "TimeoutTimeInSeconds"
	variableType = "String"

	// Verify that Variable exist in json-map and extract variable
	tempMapVariable, existInMap = methodParameter[variableName]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "6f6d8909-f7bf-4174-be87-f8a030fbfd59",
		}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

	}

	// Transform variable into correct type
	timeoutTimeInSecondsAsString, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
	if canCastTempMapVariableToCorrectVariableType == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "4148e257-16a2-4dfa-931f-cc295275909d",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into an '%s'",
			variableName,
			variableType))

	}

	// Transform 'String' into 'Int
	timeoutTimeInSeconds, err = strconv.Atoi(timeoutTimeInSecondsAsString)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "4ff0ad02-0f61-4b8f-8e40-8434a4d7bacd",
		}).Fatalln(fmt.Sprintf(" Parameter '%s' with value '%s' couldn't be converted from 'String' into an 'Integer'",
			variableName,
			timeoutTimeInSecondsAsString))

	}

	// Depending on 'testStepActionMethod', extract extra incoming parameters
	switch testStepActionMethod {

	case "SendMT540_v1_0", "SendMT542_v1_0":
		// No extra parameters

	case "ValidateMT544_v1_0", "ValidateMT546_v1_0", "ValidateMT548_v1_0":
		// Extract 'RelatedReference_54x_20CRELA'

		// *** Extract RelatedReference_54x_20CRELA ****
		variableName = "RelatedReference_54x_20CRELA"
		variableType = "String"

		// Verify that Variable exist in json-map and extract variable
		tempMapVariable, existInMap = methodParameter[variableName]
		if existInMap == false {

			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "923321b9-a04d-46b5-8621-07c2f4e9a91a",
			}).Fatalln(fmt.Sprintf("Missing parameter '%s'", variableName))

		}

		// Transform variable into correct type
		relatedReference_54x_20CRELA, canCastTempMapVariableToCorrectVariableType = tempMapVariable.(string)
		if canCastTempMapVariableToCorrectVariableType == false {

			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "87c1a80c-f5da-47a1-8d25-2c232ee7f770",
			}).Fatalln(fmt.Sprintf(" Parameter '%s' couldn't be transformed into a '%s'",
				variableName,
				variableType))

		}

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "beee23b8-5cab-4f8b-9cea-0a2301e78c06",
		}).Info(fmt.Sprintf("Got 'RelatedReference_54x_20CRELA' = '%s' as input",
			relatedReference_54x_20CRELA))

	default:
		// Unhandled Extra parameters
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                   "6b85e2f1-3faa-4880-b89f-b50741ccca72",
			"testStepActionMethod": testStepActionMethod,
		}).Fatalln(fmt.Sprintf("Couldn't find 'relatedReference_54x_20CRELA' as input. Exiting..."))
	}

	// Got this as input
	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                              "df221e75-8089-47f0-8fec-7cbc68dc2620",
		"testStepClassName":               testStepClassName,
		"testStepActionMethod":            testStepActionMethod,
		"testDataParameterType":           testDataParameterType,
		"expectedToBePassedTestApiLevel":  expectedToBePassedTestApiLevel,
		"methodParameter":                 methodParameter,
		"expectedToBePassed":              expectedToBePassedAsString,
		"testCaseExecutionUuid":           testCaseExecutionUuid,
		"testInstructionExecutionUuid":    testInstructionExecutionUuid,
		"testInstructionExecutionVersion": testInstructionExecutionVersionAsString,
		"testInstructionVersion":          testInstructionVersion,
		"relatedReference_54x_20CRELA":    relatedReference_54x_20CRELA,
	}).Info(fmt.Sprintf("Got these parameters as input"))

	// **** Start creating Responses ***

	// If this TestStep is not expected to be passed then respond her
	// 'expectedToBePassed' is 'false' - Will always go in here
	// There are special rules for 'SendMT540_v1_0' and 'SendMT542_v1_0', see below
	if expectedToBePassedAsString == "false" &&
		testStepActionMethod != "SendMT540_v1_0" && testStepActionMethod != "SendMT542_v1_0" {

		// Create the Response from the "Fenix-code"
		var tempTestApiEngineFinalTestInstructionExecutionResult TestApiEngineFinalTestInstructionExecutionResultStruct
		tempTestApiEngineFinalTestInstructionExecutionResult = TestApiEngineFinalTestInstructionExecutionResultStruct{
			TestApiEngineResponseJsonSchemaVersion: "v1.0",
			TestInstructionExecutionUUID:           testInstructionExecutionUuid,
			TestInstructionExecutionVersion:        testInstructionExecutionVersionAsString,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_FINISHED_NOT_OK)],
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp.Format(time.RFC3339),
			TestInstructionExecutionEndTimeStamp:   time.Now().Format(time.RFC3339),
			ResponseVariables:                      []ResponseVariableStruct{},
			LogPosts: []LogPostStruct{
				{
					LogPostTimeStamp:                     time.Now().Format(time.RFC3339),
					LogPostStatus:                        fenixExecutionWorkerGrpcApi.LogPostStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR)],
					LogPostText:                          "Variable 'expectedToBePassed' is 'false' and in the mocked version then the TestInstructionExecution fails!",
					FoundVersusExpectedValueForVariables: nil,
				},
			},
		}

		var jsonBytes []byte
		jsonBytes, err = json.Marshal(tempTestApiEngineFinalTestInstructionExecutionResult)
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":  "90b567d0-be45-4819-b1a6-fb8c74541034",
				"err": err,
				"tempTestApiEngineFinalTestInstructionExecutionResult": tempTestApiEngineFinalTestInstructionExecutionResult,
			}).Fatalln("Couldn't convert 'tempTestApiEngineFinalTestInstructionExecutionResult' into json. Exiting...")
		}

		// Convert '"' into '\"'
		//var jsonAsString string
		//jsonAsString = strings.ReplaceAll(string(jsonBytes), `"`, `\"`)

		// Create the TestApiEngine Response
		var tempTestApiEngineResponseWithResponseValueAsString TestApiEngineResponseWithResponseValueAsStringStruct
		tempTestApiEngineResponseWithResponseValueAsString = TestApiEngineResponseWithResponseValueAsStringStruct{
			TestStepExecutionStatus: TestStepExecutionStatusStruct{
				StatusCode: 4,
				StatusText: "FETSE_FINISHED_OK",
			},
			Details:            "",
			ResponseValue:      string(jsonBytes),
			ExecutionTimeStamp: time.Now().Format(time.RFC3339),
		}

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //TODO Fang must change this

		// Create Response message
		responseBodydata, _ := json.Marshal(tempTestApiEngineResponseWithResponseValueAsString)

		fmt.Fprintf(w, string(responseBodydata))

		return

	}

	// If this TestStep is not expected to be passed then respond her
	// 'expectedToBePassed' is 'false' and special rules for 'SendMT540_v1_0'
	// Here are the special rules for 'SendMT540_v1_0'
	// 'SendMT540_v1_0' will time out on MaxTimeOutWaitTime
	if expectedToBePassedAsString == "false" &&
		testStepActionMethod == "SendMT540_v1_0" {

		// Sleep the maximum expected time
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "a552b97c-a438-4338-84ed-75e515bb79c9",
		}).Info(fmt.Sprintf("...Sleeping for 'MaxTimeOutWaitTime' = '%s' seconds", timeoutTimeInSecondsAsString))
		time.Sleep(time.Duration(timeoutTimeInSeconds) * time.Second)

		// Create the Response from the "Fenix-code"
		var tempTestApiEngineFinalTestInstructionExecutionResult TestApiEngineFinalTestInstructionExecutionResultStruct
		tempTestApiEngineFinalTestInstructionExecutionResult = TestApiEngineFinalTestInstructionExecutionResultStruct{
			TestApiEngineResponseJsonSchemaVersion: "v1.0",
			TestInstructionExecutionUUID:           testInstructionExecutionUuid,
			TestInstructionExecutionVersion:        testInstructionExecutionVersionAsString,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_TIMEOUT_INTERRUPTION)],
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp.Format(time.RFC3339),
			TestInstructionExecutionEndTimeStamp:   time.Now().Format(time.RFC3339),
			ResponseVariables:                      []ResponseVariableStruct{},
			LogPosts: []LogPostStruct{
				{
					LogPostTimeStamp: time.Now().Format(time.RFC3339),
					LogPostStatus:    fenixExecutionWorkerGrpcApi.LogPostStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR)],
					LogPostText: "Variable 'expectedToBePassed' is 'false' and in the mocked " +
						"version then 'SendMT540_v1_0' will timeout on MaxTimeOutWaitTime",
					FoundVersusExpectedValueForVariables: nil,
				},
			},
		}

		var jsonBytes []byte
		jsonBytes, err = json.Marshal(tempTestApiEngineFinalTestInstructionExecutionResult)
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":  "607701a4-d33d-48d2-9c60-eea7fdc7da64",
				"err": err,
				"tempTestApiEngineFinalTestInstructionExecutionResult": tempTestApiEngineFinalTestInstructionExecutionResult,
			}).Fatalln("Couldn't convert 'tempTestApiEngineFinalTestInstructionExecutionResult' into json. Exiting...")
		}

		// Convert '"' into '\"'
		//var jsonAsString string
		//jsonAsString = strings.ReplaceAll(string(jsonBytes), `"`, `\"`)

		// Create the TestApiEngine Response
		var tempTestApiEngineResponseWithResponseValueAsString TestApiEngineResponseWithResponseValueAsStringStruct
		tempTestApiEngineResponseWithResponseValueAsString = TestApiEngineResponseWithResponseValueAsStringStruct{
			TestStepExecutionStatus: TestStepExecutionStatusStruct{
				StatusCode: 4,
				StatusText: "FETSE_FINISHED_OK",
			},
			Details:            "",
			ResponseValue:      string(jsonBytes),
			ExecutionTimeStamp: time.Now().Format(time.RFC3339),
		}

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //TODO Fang must change this

		// Create Response message
		responseBodydata, _ := json.Marshal(tempTestApiEngineResponseWithResponseValueAsString)

		fmt.Fprintf(w, string(responseBodydata))

		return

	}

	// If this TestStep is not expected to be passed then respond her
	// 'expectedToBePassed' is 'false' and special rules for 'SendMT542_v1_0'
	// Here are the special rules for 'SendMT542_v1_0'
	// 'SendMT542_v1_0' will time out after MaxTimeOutWaitTime
	if expectedToBePassedAsString == "false" &&
		testStepActionMethod == "SendMT542_v1_0" {

		// Sleep the maximum expected time + 3 minutes
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "67ba63cd-18d0-45ce-966d-a9db3bf95f5b",
		}).Info(fmt.Sprintf("...Sleeping for 'MaxTimeOutWaitTime' = '%s' seconds + 3 minutes", timeoutTimeInSecondsAsString))

		time.Sleep(time.Duration(timeoutTimeInSeconds+3*60) * time.Second)

		// Create the Response from the "Fenix-code"
		var tempTestApiEngineFinalTestInstructionExecutionResult TestApiEngineFinalTestInstructionExecutionResultStruct
		tempTestApiEngineFinalTestInstructionExecutionResult = TestApiEngineFinalTestInstructionExecutionResultStruct{
			TestApiEngineResponseJsonSchemaVersion: "v1.0",
			TestInstructionExecutionUUID:           testInstructionExecutionUuid,
			TestInstructionExecutionVersion:        testInstructionExecutionVersionAsString,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_TIMEOUT_INTERRUPTION)],
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp.Format(time.RFC3339),
			TestInstructionExecutionEndTimeStamp:   time.Now().Format(time.RFC3339),
			ResponseVariables:                      []ResponseVariableStruct{},
			LogPosts: []LogPostStruct{
				{
					LogPostTimeStamp: time.Now().Format(time.RFC3339),
					LogPostStatus: fenixExecutionWorkerGrpcApi.LogPostStatusEnum_name[int32(
						fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR)],
					LogPostText: "Variable 'expectedToBePassed' is 'false' and in the mocked " +
						"version then 'SendMT542_v1_0' will timeout after MaxTimeOutWaitTime",
					FoundVersusExpectedValueForVariables: nil,
				},
			},
		}

		var jsonBytes []byte
		jsonBytes, err = json.Marshal(tempTestApiEngineFinalTestInstructionExecutionResult)
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":  "68526630-d788-4db6-853a-55a00102928d",
				"err": err,
				"tempTestApiEngineFinalTestInstructionExecutionResult": tempTestApiEngineFinalTestInstructionExecutionResult,
			}).Fatalln("Couldn't convert 'tempTestApiEngineFinalTestInstructionExecutionResult' into json. Exiting...")
		}

		// Convert '"' into '\"'
		//var jsonAsString string
		//jsonAsString = strings.ReplaceAll(string(jsonBytes), `"`, `\"`)

		// Create the TestApiEngine Response
		var tempTestApiEngineResponseWithResponseValueAsString TestApiEngineResponseWithResponseValueAsStringStruct
		tempTestApiEngineResponseWithResponseValueAsString = TestApiEngineResponseWithResponseValueAsStringStruct{
			TestStepExecutionStatus: TestStepExecutionStatusStruct{
				StatusCode: 4,
				StatusText: "FETSE_FINISHED_OK",
			},
			Details:            "",
			ResponseValue:      string(jsonBytes),
			ExecutionTimeStamp: time.Now().Format(time.RFC3339),
		}

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //TODO Fang must change this

		// Create Response message
		responseBodydata, _ := json.Marshal(tempTestApiEngineResponseWithResponseValueAsString)

		fmt.Fprintf(w, string(responseBodydata))

		return

	}

	// Depending on 'testStepActionMethod', create correct response extra incoming parameters
	switch testStepActionMethod {

	case "SendMT540_v1_0", "SendMT542_v1_0", "ValidateMT544_v1_0", "ValidateMT546_v1_0", "ValidateMT548_v1_0":

		// Create correct Response Variable
		var resoponseVariables ResponseVariableStruct
		switch testStepActionMethod {

		case "SendMT540_v1_0":
			resoponseVariables = ResponseVariableStruct{
				TestInstructionVersion:        "v1.0",
				ResponseVariableUUID:          "24fa2f84-827a-4c01-a86c-da42d888c295",
				ResponseVariableName:          ":20C::SEME//",
				ResponseVariableTypeUuid:      "0f6e945e-1556-4cb0-80e5-e021ebc5d8c1",
				ResponseVariableTypeName:      "54x_:20C::SEME type",
				ResponseVariableValueAsString: "MT540_" + uuid.NewString()[:8],
			}

		case "SendMT542_v1_0":
			resoponseVariables = ResponseVariableStruct{
				TestInstructionVersion:        "v1.0",
				ResponseVariableUUID:          "9dd57f25-75e0-4024-862b-e0728c066604",
				ResponseVariableName:          ":20C::SEME//",
				ResponseVariableTypeUuid:      "0f6e945e-1556-4cb0-80e5-e021ebc5d8c1",
				ResponseVariableTypeName:      "54x_:20C::SEME type",
				ResponseVariableValueAsString: "MT542_" + uuid.NewString()[:8],
			}

		case "ValidateMT544_v1_0":
			resoponseVariables = ResponseVariableStruct{
				TestInstructionVersion:        "v1.0",
				ResponseVariableUUID:          "39818ba1-676d-42d0-87da-e1080e9d5ffd",
				ResponseVariableName:          ":20C::SEME//",
				ResponseVariableTypeUuid:      "0f6e945e-1556-4cb0-80e5-e021ebc5d8c1",
				ResponseVariableTypeName:      "54x_:20C::SEME type",
				ResponseVariableValueAsString: "MT544_" + uuid.NewString()[:8],
			}

		case "ValidateMT546_v1_0":
			resoponseVariables = ResponseVariableStruct{
				TestInstructionVersion:        "v1.0",
				ResponseVariableUUID:          "5dfd7890-a0b4-4528-804a-451a77f542ad",
				ResponseVariableName:          ":20C::SEME//",
				ResponseVariableTypeUuid:      "0f6e945e-1556-4cb0-80e5-e021ebc5d8c1",
				ResponseVariableTypeName:      "54x_:20C::SEME type",
				ResponseVariableValueAsString: "MT546_" + uuid.NewString()[:8],
			}
		case "ValidateMT548_v1_0":
			resoponseVariables = ResponseVariableStruct{
				TestInstructionVersion:        testInstructionVersion,
				ResponseVariableUUID:          "8ed1ead9-741b-4115-9f78-f8a7db1d6274",
				ResponseVariableName:          ":20C::SEME//",
				ResponseVariableTypeUuid:      "0f6e945e-1556-4cb0-80e5-e021ebc5d8c1",
				ResponseVariableTypeName:      "54x_:20C::SEME type",
				ResponseVariableValueAsString: "MT548_" + uuid.NewString()[:8],
			}

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "0ade709b-4b34-408c-a2e3-4eb653efbf39",
			}).Fatalln(fmt.Sprintf("Unhandeled 'testStepActionMethod'='%s'. Exiting...", testStepActionMethod))
		}

		// Create the Response from the "Fenix-code"
		var tempTestApiEngineFinalTestInstructionExecutionResult TestApiEngineFinalTestInstructionExecutionResultStruct
		tempTestApiEngineFinalTestInstructionExecutionResult = TestApiEngineFinalTestInstructionExecutionResultStruct{
			TestApiEngineResponseJsonSchemaVersion: "v1.0",
			TestInstructionExecutionUUID:           testInstructionExecutionUuid,
			TestInstructionExecutionVersion:        testInstructionExecutionVersionAsString,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_FINISHED_OK)],
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp.Format(time.RFC3339),
			TestInstructionExecutionEndTimeStamp:   time.Now().Format(time.RFC3339),
			ResponseVariables: []ResponseVariableStruct{
				resoponseVariables,
			},
			LogPosts: []LogPostStruct{
				{
					LogPostTimeStamp:                     time.Now().Format(time.RFC3339),
					LogPostStatus:                        fenixExecutionWorkerGrpcApi.LogPostStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_OK)],
					LogPostText:                          fmt.Sprintf("Execution of '%s' was a success", testStepActionMethod),
					FoundVersusExpectedValueForVariables: nil,
				},
			},
		}

		var jsonBytes []byte
		jsonBytes, err = json.Marshal(tempTestApiEngineFinalTestInstructionExecutionResult)
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":  "099e199d-2dd5-4c6d-bc46-c0deee0e11e3",
				"err": err,
				"tempTestApiEngineFinalTestInstructionExecutionResult": tempTestApiEngineFinalTestInstructionExecutionResult,
			}).Fatalln("Couldn't convert 'tempTestApiEngineFinalTestInstructionExecutionResult' into json. Exiting...")
		}

		// Convert '"' into '\"'
		//var jsonAsString string
		//jsonAsString = strings.ReplaceAll(string(jsonBytes), `"`, `\"`)

		// Create the Final Response
		var tempTestApiEngineResponseWithResponseValueAsString TestApiEngineResponseWithResponseValueAsStringStruct
		tempTestApiEngineResponseWithResponseValueAsString = TestApiEngineResponseWithResponseValueAsStringStruct{
			TestStepExecutionStatus: TestStepExecutionStatusStruct{
				StatusCode: 4,
				StatusText: "FETSE_FINISHED_OK",
			},
			Details:            "",
			ResponseValue:      string(jsonBytes),
			ExecutionTimeStamp: time.Now().Format(time.RFC3339),
		}

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //TODO Fang must change this

		// Create Response message
		responseBodydata, _ := json.Marshal(tempTestApiEngineResponseWithResponseValueAsString)

		fmt.Fprintf(w, string(responseBodydata))

		return

	default:
		// Unhandled 'testStepActionMethod'

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "0ade709b-4b34-408c-a2e3-4eb653efbf39",
		}).Fatalln(fmt.Sprintf("Unhandeled 'testStepActionMethod'='%s'. Exiting...", testStepActionMethod))

		var tempTestApiEngineFinalTestInstructionExecutionResult TestApiEngineFinalTestInstructionExecutionResultStruct
		tempTestApiEngineFinalTestInstructionExecutionResult = TestApiEngineFinalTestInstructionExecutionResultStruct{
			TestApiEngineResponseJsonSchemaVersion: "v1.0",
			TestInstructionExecutionUUID:           testInstructionExecutionUuid,
			TestInstructionExecutionVersion:        testInstructionExecutionVersionAsString,
			TestInstructionExecutionStatus: fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.
				TestInstructionExecutionStatusEnum_TIE_FINISHED_NOT_OK)],
			TestInstructionExecutionStartTimeStamp: tempTestInstructionExecutionStartTimeStamp.Format(time.RFC3339),
			TestInstructionExecutionEndTimeStamp:   time.Now().Format(time.RFC3339),
			ResponseVariables:                      []ResponseVariableStruct{},
			LogPosts: []LogPostStruct{
				{
					LogPostTimeStamp:                     time.Now().Format(time.RFC3339),
					LogPostStatus:                        fenixExecutionWorkerGrpcApi.LogPostStatusEnum_name[int32(fenixExecutionWorkerGrpcApi.LogPostStatusEnum_EXECUTION_ERROR)],
					LogPostText:                          fmt.Sprintf("Unhandeled 'testStepActionMethod'='%s'", testStepActionMethod),
					FoundVersusExpectedValueForVariables: nil,
				},
			},
		}

		var jsonBytes []byte
		jsonBytes, err = json.Marshal(tempTestApiEngineFinalTestInstructionExecutionResult)
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id":  "099e199d-2dd5-4c6d-bc46-c0deee0e11e3",
				"err": err,
				"tempTestApiEngineFinalTestInstructionExecutionResult": tempTestApiEngineFinalTestInstructionExecutionResult,
			}).Fatalln("Couldn't convert 'tempTestApiEngineFinalTestInstructionExecutionResult' into json. Exiting...")
		}

		// Convert '"' into '\"'
		//var jsonAsString string
		//jsonAsString = strings.ReplaceAll(string(jsonBytes), `"`, `\"`)

		// Create the Final Response
		var tempTestApiEngineResponseWithResponseValueAsString TestApiEngineResponseWithResponseValueAsStringStruct
		tempTestApiEngineResponseWithResponseValueAsString = TestApiEngineResponseWithResponseValueAsStringStruct{
			TestStepExecutionStatus: TestStepExecutionStatusStruct{
				StatusCode: 4,
				StatusText: "FETSE_FINISHED_OK",
			},
			Details:            "",
			ResponseValue:      string(jsonBytes),
			ExecutionTimeStamp: time.Now().Format(time.RFC3339),
		}

		// Create Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //TODO Fang must change this

		// Create Response message
		responseBodydata, _ := json.Marshal(tempTestApiEngineResponseWithResponseValueAsString)

		fmt.Fprintf(w, string(responseBodydata))

		return

	}

}

func extractAndValidateJsonBody(responseWriterPointer *http.ResponseWriter, httpRequest *http.Request, myInputTypeVariable interface{}) (err error) {
	// If the Content-Type header is present, check that it has the value
	// information in the header.
	responseWriter := *responseWriterPointer
	if httpRequest.Header.Get("Content-Type") != "" {
		httpRequest.Header.Get("Content-Type")
		value := httpRequest.Header.Get("Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(responseWriter, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	httpRequest.Body = http.MaxBytesReader(responseWriter, httpRequest.Body, 1048576)

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(httpRequest.Body)
	dec.DisallowUnknownFields()

	var p = myInputTypeVariable //RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct
	err = dec.Decode(&p)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(responseWriter, msg, http.StatusRequestEntityTooLarge)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			log.Println(err.Error())
			http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return err
	}

	// Call decode again, using a pointer to an empty anonymous struct as
	// the destination. If the request body only contained a single JSON
	// object this will return an io.EOF error. So if we get anything else,
	// we know that there is additional data in the request body.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		http.Error(responseWriter, msg, http.StatusBadRequest)
		return
	}

	//fmt.Fprintf(responseWriter, "RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct: %+v", p)

	return nil
}
