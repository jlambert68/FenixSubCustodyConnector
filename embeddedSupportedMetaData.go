package main

import (
	_ "embed"
)

//go:embed supportedMetaData/supportedMetaData.json
var embeddedFile_SupportedMetaData []byte
