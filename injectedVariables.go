package main

// Variables that can be populated while building the application
// Used when building an executable to be run by itself
var (
	// This variable tells if the application should expect injected environment variables or not
	useInjectedEnvironmentVariables string

	// Injected Environment Variables
	Injected_AuthClientId                                                                          string
	Injected_AuthClientSecret                                                                      string
	Injected_ExecutionConnectorPort                                                                string
	Injected_ExecutionLocationForConnector                                                         string
	Injected_ExecutionLocationForWorker                                                            string
	Injected_ExecutionWorkerAddress                                                                string
	Injected_ExecutionWorkerPort                                                                   string
	Injected_ForceNewBaseLineForTestInstructionsAndTestInstructionContainers                       string
	Injected_GCPAuthentication                                                                     string
	Injected_GcpProject                                                                            string
	Injected_LocalServiceAccountPath                                                               string
	Injected_LocalWebServerAddress                                                                 string
	Injected_LocalWebServerPort                                                                    string
	Injected_LoggingLevel                                                                          string
	Injected_RelativePathToAllowedUsersList                                                        string
	Injected_ShouldPubSubReceiverBeStarted                                                         string
	Injected_TestApiEngineAddress                                                                  string
	Injected_TestApiEnginePort                                                                     string
	Injected_TestApiEngineUrlPath                                                                  string
	Injected_TestInstructionExecutionPubSubTopicBase                                               string
	Injected_ThisConnectorIsTheOneThatPublishSupportedTestInstructionsAndTestInstructionContainers string
	Injected_ThisDomainsUuid                                                                       string
	Injected_ThisExecutionDomainUuid                                                               string
	Injected_TurnOffAllCommunicationWithWorker                                                     string
	Injected_UseInternalWebServerForTestInsteadOfCallingTestApiEngine                              string
	Injected_UseNativeGcpPubSubClientLibrary                                                       string
	Injected_UsePubSubToReceiveMessagesFromWorker                                                  string
	Injected_UseServiceAccount                                                                     string
)

// Used for hard coding if Injected or real Environment Variables are expected
var falseValue string = "false"

// Map used when Extracting the value of the injected variable
var injectedVariablesMap = map[string]*string{
	// Will Injected values och real Environment Variables be used
	"useInjectedEnvironmentVariables": &falseValue,

	// Environment Variables used for testing "FenixSubCustodyTestInstructionAdmin" by itself
	"Injected_AuthClientId":                                                    &Injected_AuthClientId,
	"Injected_AuthClientSecret":                                                &Injected_AuthClientSecret,
	"Injected_ExecutionConnectorPort":                                          &Injected_ExecutionConnectorPort,
	"Injected_ExecutionLocationForConnector":                                   &Injected_ExecutionLocationForConnector,
	"Injected_ExecutionLocationForWorker":                                      &Injected_ExecutionLocationForWorker,
	"Injected_ExecutionWorkerAddress":                                          &Injected_ExecutionWorkerAddress,
	"Injected_ExecutionWorkerPort":                                             &Injected_ExecutionWorkerPort,
	"Injected_ForceNewBaseLineForTestInstructionsAndTestInstructionContainers": &Injected_ForceNewBaseLineForTestInstructionsAndTestInstructionContainers,
	"Injected_GCPAuthentication":                                               &Injected_GCPAuthentication,
	"Injected_GcpProject":                                                      &Injected_GcpProject,
	"Injected_LocalServiceAccountPath":                                         &Injected_LocalServiceAccountPath,
	"Injected_LocalWebServerAddress":                                           &Injected_LocalWebServerAddress,
	"Injected_LocalWebServerPort":                                              &Injected_LocalWebServerPort,
	"Injected_LoggingLevel":                                                    &Injected_LoggingLevel,
	"Injected_RelativePathToAllowedUsersList":                                  &Injected_RelativePathToAllowedUsersList,
	"Injected_ShouldPubSubReceiverBeStarted":                                   &Injected_ShouldPubSubReceiverBeStarted,
	"Injected_TestApiEngineAddress":                                            &Injected_TestApiEngineAddress,
	"Injected_TestApiEnginePort":                                               &Injected_TestApiEnginePort,
	"Injected_TestApiEngineUrlPath":                                            &Injected_TestApiEngineUrlPath,
	"Injected_TestInstructionExecutionPubSubTopicBase":                         &Injected_TestInstructionExecutionPubSubTopicBase,
	"Injected_ThisConnectorIsTheOneThatPublishSupportedTestInstructionsAndTestInstructionContainers": &Injected_ThisConnectorIsTheOneThatPublishSupportedTestInstructionsAndTestInstructionContainers,
	"Injected_ThisDomainsUuid":                                          &Injected_ThisDomainsUuid,
	"Injected_ThisExecutionDomainUuid":                                  &Injected_ThisExecutionDomainUuid,
	"Injected_TurnOffAllCommunicationWithWorker":                        &Injected_TurnOffAllCommunicationWithWorker,
	"Injected_UseInternalWebServerForTestInsteadOfCallingTestApiEngine": &Injected_UseInternalWebServerForTestInsteadOfCallingTestApiEngine,
	"Injected_UseNativeGcpPubSubClientLibrary":                          &Injected_UseNativeGcpPubSubClientLibrary,
	"Injected_UsePubSubToReceiveMessagesFromWorker":                     &Injected_UsePubSubToReceiveMessagesFromWorker,
	"Injected_UseServiceAccount":                                        &Injected_UseServiceAccount,
}
