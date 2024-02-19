package sharedCode

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// Extracts 'ParserLayout' from the TimeStamp(as string)
func GenerateTimeStampParserLayout(timeStampAsString string) (parserLayout string, err error) {
	// "2006-01-02 15:04:05.999999999 -0700 MST"

	var timeStampParts []string
	var timeParts []string
	var numberOfDecimals int

	// Split TimeStamp into separate parts
	timeStampParts = strings.Split(timeStampAsString, " ")

	// Validate that first part is a date with the following form '2006-01-02'
	if len(timeStampParts[0]) != 10 {

		Logger.WithFields(logrus.Fields{
			"Id":                "ffbf0682-ebc7-4e27-8ad1-0e5005fbc364",
			"timeStampAsString": timeStampAsString,
			"timeStampParts[0]": timeStampParts[0],
		}).Error("Date part has not the correct form, '2006-01-02'")

		err = errors.New(fmt.Sprintf("Date part, '%s' has not the correct form, '2006-01-02'", timeStampParts[0]))

		return "", err

	}

	// Add Date to Parser Layout
	parserLayout = "2006-01-02"

	// Add Time to Parser Layout
	parserLayout = parserLayout + " 15:04:05."

	// Split time into time and decimals
	timeParts = strings.Split(timeStampParts[1], ".")

	// Get number of decimals
	numberOfDecimals = len(timeParts[1])

	// Add Decimals to Parser Layout
	parserLayout = parserLayout + strings.Repeat("9", numberOfDecimals)

	// Add time zone, part 1, if that information exists
	if len(timeStampParts) > 2 {
		parserLayout = parserLayout + " -0700"
	}

	// Add time zone, part 2, if that information exists
	if len(timeStampParts) > 3 {
		parserLayout = parserLayout + " MST"
	}

	return parserLayout, err
}
