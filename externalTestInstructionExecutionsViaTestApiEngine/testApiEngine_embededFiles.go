package executeTestInstructionUsingTestApiEngine

import (
	_ "embed"
)

//go:embed json-schemas/FinalTestInstructionExecutionResultMessage.json-schema.json
var finalTestInstructionExecutionResultMessageJsonSchema []byte

// SendMT540
//
//go:embed json-schemas/SendMT540_v1_0_Request.json-schema.json
var sendMT540_v1_0_RequestMessageJsonSchema []byte

//go:embed json-schemas/SendMT540_v1_0_ResponseVariables.json-schema.json
var sendMT540_v1_0_ResponseVariablesMessageJsonSchema []byte

// SendMT542
//
//go:embed json-schemas/SendMT542_v1_0_Request.json-schema.json
var sendMT542_v1_0_RequestMessageJsonSchema []byte

//go:embed json-schemas/SendMT542_v1_0_ResponseVariables.json-schema.json
var sendMT542_v1_0_ResponseVariablesMessageJsonSchema []byte

// SendMT544
//
//go:embed json-schemas/ValidateMT544_v1_0_Request.json-schema.json
var validateMT544_v1_0_RequestMessageJsonSchema []byte

//go:embed json-schemas/ValidateMT544_v1_0_ResponseVariables.json-schema.json
var validateMT544_v1_0_ResponseVariablesMessageJsonSchema []byte
