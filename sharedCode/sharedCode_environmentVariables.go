package sharedCode

// Environment Variables used locally
var (

	// UseInternalWebServerForTestInsteadOfCallingTestApiEngine
	// When 'true' calls for TestApiEngine will be sent to a local web server instead
	// Used for doing testing
	UseInternalWebServerForTestInsteadOfCallingTestApiEngine bool
)
