include MakeFileSecretVariables
include MakeFIleJsonSchemas

RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6673

filename :=
filenamePartFirst := FenixSubCustodyConnector
filenamePartLast := .exe
datetime := `date +'%y%m%d_%H%M%S'`

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

	echo "$(fullLocalFilePathOverAll)"
	echo "$(fullRemoteFilePathOverAll)"
	echo "......."
	echo "$(fullLocalFilePathFenixGeneral_SendTestDataToThisDomainRequest_v1_0)"
	echo "$(fullRemoteFilePathFenixGeneral_SendTestDataToThisDomain_v1_0)"
	echo "$(fullLocalFilePathFenixGeneral_SendTestDataToThisDomainMethodParameters_v1_0)"
	echo "$(fullRemoteFilePathFenixGeneral_SendTestDataToThisDomainMethodParameters_v1_0)"
	echo "$(fullLocalFilePathFenixGeneral_SendTestDataToThisDomain_v1_0)"
	echo "$(fullRemoteFilePathFenixGeneral_SendGeneralResponse_v1_0)"

	@curl -L -o $(fullLocalFilePathOverAll) "$(fullRemoteFilePathOverAll)"
	@curl -L -o $(fullLocalFilePathTestApiEngineResponse) "$(fullRemoteFilePathTestApiEngineResponse)"

	@curl -L -o $(fullLocalFilePathSendOnMQTypeMT_SendGeneralRequest_v1_0) "$(fullRemoteFilePathSendOnMQTypeMT_SendGeneralRequest_v1_0)"
	@curl -L -o $(fullLocalFilePathSendOnMQTypeMT_SendGeneralRequestMethodParameters_v1_0) "$(fullRemoteFilePathSendOnMQTypeMT_SendGeneralRequestMethodParameters_v1_0)"
	@curl -L -o $(fullLocalFilePathSendOnMQTypeMT_SendGeneral_v1_0) "$(fullRemoteFilePathSendOnMQTypeMT_SendGeneralResponse_v1_0)"

	@curl -L -o $(fullLocalFilePathFenixGeneral_SendTestDataToThisDomainRequest_v1_0) "$(fullRemoteFilePathFenixGeneral_SendTestDataToThisDomain_v1_0)"
	@curl -L -o $(fullLocalFilePathFenixGeneral_SendTestDataToThisDomainMethodParameters_v1_0) "$(fullRemoteFilePathFenixGeneral_SendTestDataToThisDomainMethodParameters_v1_0)"
	@curl -L -o $(fullLocalFilePathFenixGeneral_SendTestDataToThisDomain_v1_0) "$(fullRemoteFilePathFenixGeneral_SendGeneralResponse_v1_0)"

	@curl -L -o $(fullLocalFilePathSendMT540Request_v1_0) "$(fullRemoteFilePathSendMT540Request_v1_0)"
	@curl -L -o $(fullLocalFilePathSendMT540RequestMethodParameters_v1_0) "$(fullRemoteFilePathSendMT540RequestMethodParameters_v1_0)"
	@curl -L -o $(fullLocalFilePathSendMT540_v1_0) "$(fullRemoteFilePathSendMT540Response_v1_0)"

	@curl -L -o $(fullLocalFilePathSendMT542Request_v1_0) "$(fullRemoteFilePathSendMT542Request_v1_0)"
	@curl -L -o $(fullLocalFilePathSendMT542RequestMethodParameters_v1_0) "$(fullRemoteFilePathSendMT542RequestMethodParameters_v1_0)"
	@curl -L -o $(fullLocalFilePathSendMT542_v1_0) "$(fullRemoteFilePathSendMT542Response_v1_0)"

	@curl -L -o $(fullLocalFilePathValidateMT544Request_v1_0) "$(fullRemoteFilePathValidateMT544Request_v1_0)"
	@curl -L -o $(fullLocalFilePathValidateMT544RequestMethodParameters_v1_0) "$(fullRemoteFilePathValidateMT544RequestMethodParameters_v1_0)"
	@curl -L -o $(fullLocalFilePathValidateMT544_v1_0) "$(fullRemoteFilePathValidateMT544_v1_0)"

	@curl -L -o $(fullLocalFilePathValidateMT546Request_v1_0) "$(fullRemoteFilePathValidateMT546Request_v1_0)"
	@curl -L -o $(fullLocalFilePathValidateMT546RequestMethodParameters_v1_0) "$(fullRemoteFilePathValidateMT546RequestMethodParameters_v1_0)"
	@curl -L -o $(fullLocalFilePathValidateMT546_v1_0) "$(fullRemoteFilePathValidateMT546_v1_0)"
    	
	@curl -L -o $(fullLocalFilePathValidateMT548Request_v1_0) "$(fullRemoteFilePathValidateMT548Request_v1_0)"
	@curl -L -o $(fullLocalFilePathValidateMT548RequestMethodParameters_v1_0) "$(fullRemoteFilePathValidateMT548RequestMethodParameters_v1_0)"
	@curl -L -o $(fullLocalFilePathValidateMT548_v1_0) "$(fullRemoteFilePathValidateMT548_v1_0)"

	@curl -L -o $(fullLocalFilePathVerifyReceivedTypeMT5xxRequest_v1_0) "$(fullRemoteFilePathVerifyReceivedTypeMT5xxRequest_v1_0)"
	@curl -L -o $(fullLocalFilePathVerifyReceivedTypeMT5xxRequestMethodParameters_v1_0) "$(fullRemoteFilePathVerifyReceivedTypeMT5xxRequestMethodParameters_v1_0)"
	@curl -L -o $(fullLocalFilePathVerifyReceivedTypeMT5xx_v1_0) "$(fullRemoteFilePathVerifyReceivedTypeMT5xx_v1_0)"
#$(localJsonSchemaPath)/$(localJsonSchemaOverAllLName)