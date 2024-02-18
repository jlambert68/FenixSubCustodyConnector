package executeTestInstructionUsingTestApiEngine

import (
	_ "embed"
)

//go:embed json-schemas/FinalTestInstructionExecutionResultMessage.json-schema.json
var finalTestInstructionExecutionResultMessageJsonSchema []byte

//go:embed json-schemas/SendMT540_v1_0_ResponseVariables.json-schema.json
var sendMT540_v1_0_ResponseVariablesMessageJsonSchema []byte
