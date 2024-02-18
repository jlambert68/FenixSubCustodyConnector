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

# Remote SendMT540
# Remote Json-Schemas
jsonSchemaPathSendMT540_v1_0 := TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540/version_1_0/
jsonSchemaFileNameSendMT540 := SendMT540_ResponseVariables.json-schema.json



# Full Remote Path and FilePath - SendMT540
fullRemotePathSendMT540_v1_0 := $(remoteUrl)/$(githubUsername)/$(repositorySubCustody)/$(branchSpecific)/$(jsonSchemaPathSendMT540_v1_0)
fullRemoteFilePathSendMT540_v1_0 := $(fullRemotePathSendMT540_v1_0)/$(jsonSchemaFileNameSendMT540)

# Local
localJsonSchemaPath := externalTestInstructionExecutionsViaTestApiEngine/json-schemas
localJsonSchemaFileNameSendMT540_v1_0 := SendMT540_v1_0_ResponseVariables.json-schema.json

fullLocalPath := $(localJsonSchemaPath)
fullLocalFilePathOverAll := $(fullLocalPath)/$(jsonSchemaFileNameOverAllL)
fullLocalFilePathSendMT540_v1_0 := $(fullLocalPath)/$(localJsonSchemaFileNameSendMT540_v1_0)



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
	echo "$(fullRemoteFilePathSendMT540_v1_0)"
	@curl -L -o $(fullLocalFilePathOverAll) "$(fullRemoteFilePathOverAll)"
	@curl -L -o $(fullLocalFilePathSendMT540_v1_0) "$(fullRemoteFilePathSendMT540_v1_0)"
#$(localJsonSchemaPath)/$(localJsonSchemaOverAllLName)