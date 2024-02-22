package main

// Variables that can be populated while building the application
// Used when building an executable to be run by itself
var (
	// This variable tells if the application should expect injected environment variables or not
	useInjectedEnvironmentVariables string

	// Injected Environment Variables
	injectedAuthClientId                                                                          string
	injectedAuthClientSecret                                                                      string
	injectedExecutionConnectorPort                                                                string
	injectedExecutionLocationForConnector                                                         string
	injectedExecutionLocationForWorker                                                            string
	injectedExecutionWorkerAddress                                                                string
	injectedExecutionWorkerPort                                                                   string
	injectedForceNewBaseLineForTestInstructionsAndTestInstructionContainers                       string
	injectedGCPAuthentication                                                                     string
	injectedGcpProject                                                                            string
	injectedLocalServiceAccountPath                                                               string
	injectedLocalWebServerAddress                                                                 string
	injectedLocalWebServerPort                                                                    string
	injectedLoggingLevel                                                                          string
	injectedRelativePathToAllowedUsersList                                                        string
	injectedShouldPubSubReceiverBeStarted                                                         string
	injectedTestApiEngineAddress                                                                  string
	injectedTestApiEnginePort                                                                     string
	injectedTestApiEngineUrlPath                                                                  string
	injectedTestInstructionExecutionPubSubTopicBase                                               string
	injectedThisConnectorIsTheOneThatPublishSupportedTestInstructionsAndTestInstructionContainers string
	injectedThisDomainsUuid                                                                       string
	injectedThisExecutionDomainUuid                                                               string
	injectedTurnOffAllCommunicationWithWorker                                                     string
	injectedUseInternalWebServerForTestInsteadOfCallingTestApiEngine                              string
	injectedUseNativeGcpPubSubClientLibrary                                                       string
	injectedUsePubSubToReceiveMessagesFromWorker                                                  string
	injectedUseServiceAccount                                                                     string
)

// Used for hard coding if Injected or real Environment Variables are expected
var falseOrTrueValue string = "false"

// Map used when Extracting the value of the injected variable
var injectedVariablesMap = map[string]*string{
	// Will Injected values och real Environment Variables be used
	"useInjectedEnvironmentVariables": &falseOrTrueValue,

	// Environment Variables used for testing "FenixSubCustodyTestInstructionAdmin" by itself
	"Injected_AuthClientId":                                                    &injectedAuthClientId,
	"Injected_AuthClientSecret":                                                &injectedAuthClientSecret,
	"Injected_ExecutionConnectorPort":                                          &injectedExecutionConnectorPort,
	"Injected_ExecutionLocationForConnector":                                   &injectedExecutionLocationForConnector,
	"Injected_ExecutionLocationForWorker":                                      &injectedExecutionLocationForWorker,
	"Injected_ExecutionWorkerAddress":                                          &injectedExecutionWorkerAddress,
	"Injected_ExecutionWorkerPort":                                             &injectedExecutionWorkerPort,
	"Injected_ForceNewBaseLineForTestInstructionsAndTestInstructionContainers": &injectedForceNewBaseLineForTestInstructionsAndTestInstructionContainers,
	"Injected_GCPAuthentication":                                               &injectedGCPAuthentication,
	"Injected_GcpProject":                                                      &injectedGcpProject,
	"Injected_LocalServiceAccountPath":                                         &injectedLocalServiceAccountPath,
	"Injected_LocalWebServerAddress":                                           &injectedLocalWebServerAddress,
	"Injected_LocalWebServerPort":                                              &injectedLocalWebServerPort,
	"Injected_LoggingLevel":                                                    &injectedLoggingLevel,
	"Injected_RelativePathToAllowedUsersList":                                  &injectedRelativePathToAllowedUsersList,
	"Injected_ShouldPubSubReceiverBeStarted":                                   &injectedShouldPubSubReceiverBeStarted,
	"Injected_TestApiEngineAddress":                                            &injectedTestApiEngineAddress,
	"Injected_TestApiEnginePort":                                               &injectedTestApiEnginePort,
	"Injected_TestApiEngineUrlPath":                                            &injectedTestApiEngineUrlPath,
	"Injected_TestInstructionExecutionPubSubTopicBase":                         &injectedTestInstructionExecutionPubSubTopicBase,
	"Injected_ThisConnectorIsTheOneThatPublishSupportedTestInstructionsAndTestInstructionContainers": &injectedThisConnectorIsTheOneThatPublishSupportedTestInstructionsAndTestInstructionContainers,
	"Injected_ThisDomainsUuid":                                          &injectedThisDomainsUuid,
	"Injected_ThisExecutionDomainUuid":                                  &injectedThisExecutionDomainUuid,
	"Injected_TurnOffAllCommunicationWithWorker":                        &injectedTurnOffAllCommunicationWithWorker,
	"Injected_UseInternalWebServerForTestInsteadOfCallingTestApiEngine": &injectedUseInternalWebServerForTestInsteadOfCallingTestApiEngine,
	"Injected_UseNativeGcpPubSubClientLibrary":                          &injectedUseNativeGcpPubSubClientLibrary,
	"Injected_UsePubSubToReceiveMessagesFromWorker":                     &injectedUsePubSubToReceiveMessagesFromWorker,
	"Injected_UseServiceAccount":                                        &injectedUseServiceAccount,
}
