package executeTestInstructionUsingTestApiEngine

import (
	_ "embed"
)

// TestApiEngine-response
//
//go:embed json-schemas/TestApiEngineResponse-json-schema.json
var testApiEngineResponseMessageJsonSchemaAsByteArray []byte

// FinalTestInstructionExecutionResultMessage
//
//go:embed json-schemas/FinalTestInstructionExecutionResultMessage.json-schema.json
var finalTestInstructionExecutionResultMessageJsonSchemaAsByteArray []byte

// FenixGeneral_SendTestDataToThisDomain
//
//go:embed json-schemas/FenixGeneral_SendTestDataToThisDomain_v1_0_Request.json-schema.json
var fenixGeneral_SendTestDataToThisDomain_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/FenixGeneral_SendTestDataToThisDomain_v1_0_RequestMethodParameters.json-schema.json
var fenixGeneral_SendTestDataToThisDomainl_v1_0_RequestMethodParametersMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/FenixGeneral_SendTestDataToThisDomain_v1_0_ResponseVariables.json-schema.json
var fenixGeneral_SendTestDataToThisDomain_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendOnMQTypeMT_SendGeneral
//
//go:embed json-schemas/SendOnMQTypeMT_SendGeneral_v1_0_Request.json-schema.json
var sendOnMQTypeMT_SendGeneral_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/SendOnMQTypeMT_SendGeneral_v1_0_RequestMethodParameters.json-schema.json
var sendOnMQTypeMT_SendGeneral_v1_0_RequestMethodParametersMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/SendOnMQTypeMT_SendGeneral_v1_0_ResponseVariables.json-schema.json
var sendOnMQTypeMT_SendGeneral_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendMT540
//
//go:embed json-schemas/SendMT540_v1_0_Request.json-schema.json
var sendMT540_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/Send540_v1_0_RequestMethodParameters.json-schema.json
var sendMT540_v1_0_RequestMethodParametersMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/SendMT540_v1_0_ResponseVariables.json-schema.json
var sendMT540_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendMT542
//
//go:embed json-schemas/SendMT542_v1_0_Request.json-schema.json
var sendMT542_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/Send542_v1_0_RequestMethodParameters.json-schema.json
var sendMT542_v1_0_RequestMethodParametersMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/SendMT542_v1_0_ResponseVariables.json-schema.json
var sendMT542_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendMT544
//
//go:embed json-schemas/ValidateMT544_v1_0_RequestMethodParameters.json-schema.json
var validateMT544_v1_0_RequestMethodParametersMessageJsonSchemaAsByteArray []byte

//
//go:embed json-schemas/ValidateMT544_v1_0_Request.json-schema.json
var validateMT544_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/ValidateMT544_v1_0_ResponseVariables.json-schema.json
var validateMT544_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendMT546
//
//go:embed json-schemas/ValidateMT546_v1_0_Request.json-schema.json
var validateMT546_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/ValidateMT546_v1_0_RequestMethodParameters.json-schema.json
var validateMT546_v1_0_RequestMethodParametersMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/ValidateMT546_v1_0_ResponseVariables.json-schema.json
var validateMT546_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendMT548
//
//go:embed json-schemas/ValidateMT548_v1_0_Request.json-schema.json
var validateMT548_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/ValidateMT548_v1_0_RequestMethodParameters.json-schema.json
var validateMT548_v1_0_RequestMethodParametersMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/ValidateMT548_v1_0_ResponseVariables.json-schema.json
var validateMT548_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendVerifyReceivedTypeMT5xx
//
//go:embed json-schemas/VerifyReceivedTypeMT5xx_v1_0_Request.json-schema.json
var verifyReceivedTypeMT5xx_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/VerifyReceivedTypeMT5xx_v1_0_RequestMethodParameters.json-schema.json
var verifyReceivedTypeMT5xx_v1_0_RequestMethodParametersMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/VerifyReceivedTypeMT5xx_v1_0_ResponseVariables.json-schema.json
var verifyReceivedTypeMT5xx_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte
