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

// SendMT540
//
//go:embed json-schemas/SendMT540_v1_0_Request.json-schema.json
var sendMT540_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/SendMT540_v1_0_ResponseVariables.json-schema.json
var sendMT540_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendMT542
//
//go:embed json-schemas/SendMT542_v1_0_Request.json-schema.json
var sendMT542_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/SendMT542_v1_0_ResponseVariables.json-schema.json
var sendMT542_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte

// SendMT544
//
//go:embed json-schemas/ValidateMT544_v1_0_Request.json-schema.json
var validateMT544_v1_0_RequestMessageJsonSchemaAsByteArray []byte

//go:embed json-schemas/ValidateMT544_v1_0_ResponseVariables.json-schema.json
var validateMT544_v1_0_ResponseVariablesMessageJsonSchemaAsByteArray []byte
