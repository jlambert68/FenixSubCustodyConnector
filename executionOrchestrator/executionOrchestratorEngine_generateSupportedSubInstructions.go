package executionOrchestrator

// Generates the 'SupportedSubInstructions' that will be sent via gRPC to Worker
func generateSupportedSubInstructions() *[]byte {

	return &supportedSubInstructions

}

// Generates the 'SupportedSubInstructionsPerTestInstruction' that will be sent via gRPC to Worker
func generateSupportedSubInstructionsPerTestInstruction() *[][]byte {

	return &supportedSubInstructionsPerTestInstructionSlice

}
