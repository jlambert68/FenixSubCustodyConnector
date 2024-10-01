package main

import (
	_ "embed"
)

//go:embed simpleTestData/FenixRawTestdata_646rows_211220.csv
var embeddedFile_SubCustody_MainTestDataArea []byte

//go:embed simpleTestData/FenixRawTestdata_3rows_240705.csv
var embeddedFile_SubCustody_ExtraTestDataArea []byte

//go:embed simpleTestData/FenixRawTestdata_10rows_240705.csv
var embeddedFile_CustodyCash_MainTestDataArea []byte

//go:embed simpleTestData/TestData.csv
var embeddedFile_TestData []byte
