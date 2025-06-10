package main

import (
	_ "embed"
)

//go:embed supportedMetaData/supportedTestCaseMetaData.json
var embeddedFile_SupportedTestCaseMetaData []byte

//go:embed supportedMetaData/supportedTestSuiteMetaData.json
var embeddedFile_SupportedTestSuiteMetaData []byte
