include MakeFileSecretVariables

RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6673

filename :=
filenamePartFirst := FenixSubCustodyConnector
filenamePartLast := .exe
datetime := `date +'%y%m%d_%H%M%S'`

githubUsername := jlambert68
repositoryOverAll := FenixGrpcApi
repositorySpecifc := FenixSubCustodyConnector
branchOverAll := master
branschSpecific := master

localJsonSchemaPath := json-schmas

jsonSchemaOverAll := FenixExecutionServer/fenixExecutionConnectorGrpcApi/json-schema/FinalTestInstructionExecutionResultMessage.json-schema.json
localJsonSchemaOverAllLName := FinalTestInstructionExecutionResultMessage.json-schema.json



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
	@curl -L -o  $(localJsonSchemaPath)/$(localJsonSchemaOverAllLName) "https://github.com/$(githubUsername)/$(repositoryOverall)/$(branchOverAll)/$(jsonSchemaOverAll)"
