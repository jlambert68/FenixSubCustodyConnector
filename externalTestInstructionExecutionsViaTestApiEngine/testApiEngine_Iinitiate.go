package executeTestInstructionUsingTestApiEngine

// InitiateRestCallsToCAEngine
// Do all initiation to have restEngine be able to do RestCalls to Sub Custodys TestApiEngine
func InitiateTestApiEngine() {

	// Extract environment variables regarding TestApiEngine
	Init()

	// Start local web server for tests instead od doing call to TestApiEngine
	if UseInternalWebServerForTestInsteadOfCallingTestApiEngine == true {
		go RestAPIServer()
	}

}
