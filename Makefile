include MakeFileSecretVariables

RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6673

filename :=
filenamePartFirst := FenixSubCustodyConnector
filenamePartLast := .exe
datetime := `date +'%y%m%d_%H%M%S'`

# Remote OverAll
remoteUrl := https://raw.githubusercontent.com
githubUsername := jlambert68

repositoryOverAll := FenixGrpcApi
branchOverAll := master

# Remote Json-Schemas
jsonSchemaPathOverAll := FenixExecutionServer/fenixExecutionConnectorGrpcApi/json-schema/
jsonSchemaFileNameOverAllL := FinalTestInstructionExecutionResultMessage.json-schema.json


# Full Remote Path and FilePath - OverAll
fullRemotePathOverAll := $(remoteUrl)/$(githubUsername)/$(repositoryOverAll)/$(branchOverAll)/$(jsonSchemaPathOverAll)
fullRemoteFilePathOverAll := $(fullRemotePathOverAll)/$(jsonSchemaFileNameOverAllL)

# Remote Specific
repositorySubCustody := FenixSubCustodyTestInstructionAdmin
branchSpecific := master

# **** Remote Json-Schemas SendMT540 ****
jsonSchemaPathSendMT540_v1_0 := TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540/version_1_0/
jsonSchemaFileNameSendMT540Request := Send540.json-schema.json
jsonSchemaFileNameSendMT540Response := SendMT540_ResponseVariables.json-schema.json

# Full Remote Path and FilePath - SendMT540
fullRemotePathSendMT540_v1_0 := $(remoteUrl)/$(githubUsername)/$(repositorySubCustody)/$(branchSpecific)/$(jsonSchemaPathSendMT540_v1_0)
fullRemoteFilePathSendMT540Response_v1_0 := $(fullRemotePathSendMT540_v1_0)/$(jsonSchemaFileNameSendMT540Response)

# General Local path
localJsonSchemaPath := externalTestInstructionExecutionsViaTestApiEngine/json-schemas
fullLocalPath := $(localJsonSchemaPath)
fullLocalFilePathOverAll := $(fullLocalPath)/$(jsonSchemaFileNameOverAllL)

# Local - SendMT540
localJsonSchemaFileNameSendMT540_v1_0 := SendMT540_v1_0_ResponseVariables.json-schema.json
fullLocalFilePathSendMT540_v1_0 := $(fullLocalPath)/$(localJsonSchemaFileNameSendMT540_v1_0)

# **** Remote Json-Schemas SendMT542 ****
jsonSchemaPathSendMT542_v1_0 := TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0/
jsonSchemaFileNameSendMT542Request := Send542.json-schema.json
jsonSchemaFileNameSendMT542Response := SendMT542_ResponseVariables.json-schema.json

# Full Remote Path and FilePath - SendMT542
fullRemotePathSendMT542_v1_0 := $(remoteUrl)/$(githubUsername)/$(repositorySubCustody)/$(branchSpecific)/$(jsonSchemaPathSendMT542_v1_0)
fullRemoteFilePathSendMT542Response_v1_0 := $(fullRemotePathSendMT542_v1_0)/$(jsonSchemaFileNameSendMT542Response)

# Local - SendMT542
localJsonSchemaFileNameSendMT542_v1_0 := SendMT542_v1_0_ResponseVariables.json-schema.json
fullLocalFilePathSendMT542_v1_0 := $(fullLocalPath)/$(localJsonSchemaFileNameSendMT542_v1_0)

# **** Remote Json-Schemas ValidateMT544 ****
jsonSchemaPathValidateMT544_v1_0 := TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0/
jsonSchemaFileNameValidateMT544Request := Validate544.json-schema.json
jsonSchemaFileNameValidateMT544Response := ValidateMT544_ResponseVariables.json-schema.json

# Full Remote Path and FilePath - ValidateMT544
fullRemotePathValidateMT544_v1_0 := $(remoteUrl)/$(githubUsername)/$(repositorySubCustody)/$(branchSpecific)/$(jsonSchemaPathValidateMT544_v1_0)
fullRemoteFilePathValidateMT542Response_v1_0 := $(fullRemotePathValidateMT544_v1_0)/$(jsonSchemaFileNameValidateMT544Response)

# Local - ValidateMT544
localJsonSchemaFileNameValidateMT544_v1_0 := ValidateMT544_v1_0_ResponseVariables.json-schema.json
fullLocalFilePathValidateMT544_v1_0 := $(fullLocalPath)/$(localJsonSchemaFileNameValidateMT544_v1_0)







GenerateDateTime:
	$(eval fileName := $(filenamePartFirst)$(datetime)$(filenamePartLast))

	echo $(fileName)

GenerateTrayIcons:
	./bundleIcons.sh

PrintInjectValues:
	echo "$(InjectValue)"

#BuildExeForWindows:
#	fyne-cross windows -arch=amd64 --ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.executionConnectorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.executionWorkerAddress=fenixexecutionworker-ca-nwxrrpoxea-lz.a.run.app' -X 'main.executionWorkerPort=443' -X 'main.gcpAuthentication=false'"
#	GOOD=windows GOARCH=amd64 go build -o FenixSCConnectorWindow.exe -ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.executionConnectorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.executionWorkerAddress=fenixexecutionworker-ca-nwxrrpoxea-lz.a.run.app' -X 'main.executionWorkerPort=443' -X  'main.gcpAuthentication=true' -X 'main.caEngineAddress=127.0.0.1' -X 'main.caEngineAddressPath=/"
	GOOD=windows GOARCH=amd64 go build -o FenixSubCustodyConnector.$(datetime).WindowsExe.exe -ldflags="$(InjectValue)"
BuildExeForLinux:
	GOOD=linux GOARCH=amd64 go build -o FenixSubCustodyConnector.$(datetime).LinuxExe -ldflags="$(InjectValue)"

CrossBuildForWindows_SEB_test:
	$(eval fileName := $(filenamePartFirst)$(datetime)$(filenamePartLast))
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o $(fileName) -ldflags="$(InjectValue)" .

Download-json-schemas:
	echo "$(fullLocalFilePathSendMT540_v1_0)"
	echo "$(fullRemoteFilePathSendMT540_v1_0_Response)"
	@curl -L -o $(fullLocalFilePathOverAll) "$(fullRemoteFilePathOverAll)"
	@curl -L -o $(fullLocalFilePathSendMT540_v1_0) "$(fullRemoteFilePathSendMT540Response_v1_0)"
	@curl -L -o $(fullLocalFilePathSendMT542_v1_0) "$(fullRemoteFilePathSendMT542Response_v1_0)"
	@curl -L -o $(fullLocalFilePathValidateMT544_v1_0) "$(fullRemoteFilePathValidateMT542Response_v1_0)"
#$(localJsonSchemaPath)/$(localJsonSchemaOverAllLName)