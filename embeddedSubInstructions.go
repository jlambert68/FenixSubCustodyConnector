package main

import (
	_ "embed"
)

//go:embed supportedSubInstructions/supportedSubInstructions_All.json
var embeddedFile_SupportedSubInstructions []byte

//go:embed supportedSubInstructions/supportedSubInstructions_VerifyMQTypeMT_VerifyContentTypeMT5xx.json
var embeddedFiles_SupportedSubInstructions_VerifyMQTypeMT_VerifyContentTypeMT5xx []byte

var embeddedFiles_SupportedSubInstructionsPerTestInstructionSlice = [][]byte{
	embeddedFiles_SupportedSubInstructions_VerifyMQTypeMT_VerifyContentTypeMT5xx}
