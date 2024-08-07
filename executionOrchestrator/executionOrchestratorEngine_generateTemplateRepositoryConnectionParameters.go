package executionOrchestrator

import (
	"FenixSubCustodyConnector/sharedCode"
	"encoding/json"
	fenixConnectorAdminShared_sharedCode "github.com/jlambert68/FenixConnectorAdminShared/common_config"
	"github.com/sirupsen/logrus"
)

// Generates the 'TemplateRepositoryConnectionParameters' that will be sent via gRPC to Worker
func generateTemplateRepositoryConnectionParameters() *fenixConnectorAdminShared_sharedCode.
	RepositoryTemplatePathStruct {

	var allTemplateRepositoryConnectionParameters *fenixConnectorAdminShared_sharedCode.RepositoryTemplatePathStruct

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal(templateUrlParameters, &allTemplateRepositoryConnectionParameters)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                    "a44b894a-3f47-474d-9752-32e96c0b2bf2",
			"err":                   err,
			"templateUrlParameters": string(templateUrlParameters),
		}).Fatalln("Error unmarshalling JSON")
	}

	// Loop

	return allTemplateRepositoryConnectionParameters

}
