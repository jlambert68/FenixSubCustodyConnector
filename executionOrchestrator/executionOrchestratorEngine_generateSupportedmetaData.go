package executionOrchestrator

// Generates the 'SupportedMetaData' that will be sent via gRPC to Worker
func generateSupportedTestCaseMetaData() *[]byte {

	return &supportedTestCaseMetaData

}

// Generates the 'SupportedMetaData' that will be sent via gRPC to Worker
func generateSupportedTestSuiteMetaData() *[]byte {

	return &supportedTestSuiteMetaData

}
