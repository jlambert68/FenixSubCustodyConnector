package executionOrchestrator

import (
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
)

// Generates the Simple" TestData that will be sent via gRPC to Worker
// Should return 'nil' if there is no 'simple' TestData
func generateSimpleTestData() []*testDataEngine.TestDataFromSimpleTestDataAreaStruct {

	var testDataFiles []*testDataEngine.TestDataFromSimpleTestDataAreaStruct
	var fileHash string

	for _, testDataFile := range simpleTestDataFiles {

		var simpleTestDataFile testDataEngine.TestDataFromSimpleTestDataAreaStruct
		simpleTestDataFile = testDataEngine.ImportEmbeddedSimpleCsvTestDataFile(testDataFile, ';')

		// Get file hash and add to data
		fileHash = fenixSyncShared.HashSingleValue(string(testDataFile))
		simpleTestDataFile.TestDataFileSha256Hash = fileHash

		// Add TestData to slice of all TestData
		testDataFiles = append(testDataFiles, &simpleTestDataFile)
	}

	return testDataFiles

}
