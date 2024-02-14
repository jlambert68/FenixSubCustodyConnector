package executeTestInstructionUsingTestApiEngine

// Environment Variables used locally
var (
	// TurnAllCommunicationWithWorker
	// When 'true' all communication to and from Worker is turned off.
	// Is used for doing testing of local TestInstructionExecutions calls towards TestApiEngine
	TurnAllCommunicationWithWorker bool

	// UseInternalWebServerForTestInsteadOfCallingTestApiEngine
	// When 'true' calls for TestApiEngine will be sent to a local web server instead
	// Used for doing testing
	UseInternalWebServerForTestInsteadOfCallingTestApiEngine bool

	// LocalWebServerAddress
	// Address used by local web server
	// is used for testing instead if calling TestApiEngine
	LocalWebServerAddress string

	// LocalWebServerPort
	// Port used by local web server
	// is used for testing instead if calling TestApiEngine
	LocalWebServerPort string
)
