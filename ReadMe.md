MIT License

Copyright (c) 2024 Jonas Lambert

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

***

# Fenix Inception

## External Connector
A External Connector is the component that has two main purpose. First of all do it expose the TestInstructions and TestInstructionContainers that is supported by the external TestAuto framwork. The other important task for the External Connector is to act as the bridge between the Fenix Inception world in GCP and the external TestAuto framework. 

![Fenix Inception - Standard External Connector](./Documentation/FenixInception-Overview-NonDetailed-ExternalConnector.png "Fenix Inception - External Connector")

The following environment variable is needed for the Connector to be able to run.

| Environment variable                                                                  | Example value                                                                   | comment                                       |
|---------------------------------------------------------------------------------------|---------------------------------------------------------------------------------|-----------------------------------------------|
| AuthClientId                                                                          | 94234523512385-aj56askfjs87asd9f8FJKi2j7o7ltbna0rc53.apps.googleusercontent.com |                                               |
| AuthClientSecret                                                                      | FJIONSO-M5FKJsdfu09dfFLKlkki                                                    |                                               |
| CAEngineAddress                                                                       | #                                                                               |                                               |
| CAEngineAddressPath                                                                   | #                                                                               |                                               |
| ENVIRONMENT                                                                           | dev                                                                             |                                               |
| ExecutionConnectorPort                                                                | 6673                                                                            |                                               |
| ExecutionEnvironmentPlatform                                                          | Other                                                                           |                                               |
| ExecutionLocationForConnector                                                         | GCP                                                                             |                                               |
| ExecutionLocationForWorker                                                            | GCP                                                                             |                                               |
| ExecutionWorkerAddress                                                                | fenixexecutionworkerserver-must-be-logged-in-ffafweeerg-lz.a.run.app            |                                               |
| ExecutionWorkerPort                                                                   | 443                                                                             |                                               |
| ForceNewBaseLineForTestInstructionsAndTestInstructionContainers                       | false                                                                           |                                               |
| GCPAuthentication                                                                     | true                                                                            |                                               |
| GcpProject                                                                            | mycloud-run-project                                                             |                                               |
| GenerateNewPrivateKeyForAcc                                                           | false                                                                           |                                               |
| GenerateNewPrivateKeyForDev                                                           | false                                                                           |                                               |
| GitHubApiKeys                                                                         | #                                                                               |                                               |
| LocalServiceAccountPath                                                               | #                                                                               |                                               |
| LoggingLevel                                                                          | DebugLevel                                                                      |                                               |
| PrivateKey                                                                            | ApEJHDjksdioqta+Bank/PKLifDSJAlaksdaksdKq/YA                                    |                                               |   
| ProxyServerURL                                                                        | #                                                                               |                                               |
| RelativePathToAllowedUsersList                                                        | allowedUsers/allowedUsers.json                                                  |                                               |
| ShouldProxyServerBeUsed                                                               | false                                                                           |                                               |
| ShouldPubSubReceiverBeStarted                                                         | false                                                                           |                                               |
| TestApiEngineAddress                                                                  | http://127.0.0.1                                                                | Address to external TestAuto framework        |
| TestApiEngineUrlPath                                                                  | 5000                                                                            | Port used by external TestAuto framework      |
| TestApiEngineUrlPath                                                                  | /TestCaseExecution/ExecuteTestActionMethod                                      | Rest-path used by external TestAuto framework |
| TestInstructionExecutionPubSubTopicBase                                               | ProcessTestInstructionExecutionRequest                                          |                                               |
| ThisConnectorIsTheOneThatPublishSupportedTestInstructionsAndTestInstructionContainers | true                                                                            |                                               |
| ThisDomainsUuid                                                                       | e81b9734-5dce-43c9-8d77-3368940cf126                                            |                                               |
| ThisExecutionDomainUuid                                                               | e81b9734-5dce-43c9-8d77-3368940cf126                                            |                                               |
| TurnOffAllCommunicationWithWorker                                                     | false                                                                           |                                               |
| TurnOffCallToWorker                                                                   | false                                                                           |                                               |
| UseInternalWebServerForTest                                                           | false                                                                           |                                               |
| UseNativeGcpPubSubClientLibrary                                                       | true                                                                            |                                               |
| UsePubSubToReceiveMessagesFromWorker                                                  | true                                                                            |                                               |
| UseServiceAccount                                                                     | false                                                                           |                                               |


