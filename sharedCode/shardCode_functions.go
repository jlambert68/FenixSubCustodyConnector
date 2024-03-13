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
	// "2006-01-02T15:04:05Z"
	// "2006-01-02T15:04:05Z07:00"

	var timeStampParts []string
	var dateVersion int
	var timeParts []string
	var numberOfDecimals int

	// Split TimeStamp into separate parts
	timeStampParts = strings.Split(timeStampAsString, " ")

	// Validate that first part is a date with the following form '2006-01-02'
	dateVersion = 1
	if len(timeStampParts[0]) != 10 {

		// Check TimeStamp version 2, "2006-01-02T15:04:05Z" or "2006-01-02T15:04:05Z07:00"
		timeStampParts = strings.Split(timeStampAsString, "T")
		if len(timeStampParts[0]) != 10 {

			Logger.WithFields(logrus.Fields{
				"Id":                "dd656c66-dca5-4a1f-b79d-8418c98f75b0",
				"timeStampAsString": timeStampAsString,
				"timeStampParts[0]": timeStampParts[0],
			}).Error("Date part has not the correct form, '2006-01-02 15:04:05' or '2006-01-02T15:04:05'")

			err = errors.New(fmt.Sprintf("Date part, '%s' has not the correct form, '2006-01-02'", timeStampParts[0]))

			return "", err

		} else {

			dateVersion = 2
		}
	}

	// Add Date to Parser Layout
	parserLayout = "2006-01-02"

	// Add Time to Parser Layout depending on DateTime-format
	switch dateVersion {
	case 1:
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

	case 2:
		// Decide "T15:04:05Z" or "T15:04:05Z07:00"
		switch len(timeStampParts[1]) {

		case 9:
			parserLayout = parserLayout + "T15:04:05Z"

		case 10:
			parserLayout = parserLayout + "T15:04:05Z"

		case 14:
			parserLayout = parserLayout + "T15:04:05-07:00"

		default:
			Logger.WithFields(logrus.Fields{
				"Id":                "47013eb0-37b9-400c-924a-e6f144cf7a4f",
				"timeStampAsString": timeStampAsString,
				"timeStampParts":    timeStampParts,
			}).Error("Unhandled length of 'timeStampParts'")

			return "", err
		}

		// More parts shouldn't exist
		if len(timeStampParts) > 2 {
			Logger.WithFields(logrus.Fields{
				"Id":                "fa2ee6da-eea6-4e56-ac0f-179f2f6f0a04",
				"timeStampAsString": timeStampAsString,
				"timeStampParts":    timeStampParts,
			}).Error("Unexpected number of parts in TimeStamp")
		}

	default:
		Logger.WithFields(logrus.Fields{
			"Id":                "ffbf0682-ebc7-4e27-8ad1-0e5005fbc364",
			"timeStampAsString": timeStampAsString,
			"timeStampParts[0]": timeStampParts[0],
		}).Error("Date part has not the correct form, '2006-01-02 15:04:05' or '2006-01-02T15:04:05'")

		err = errors.New(fmt.Sprintf("Date part, '%s' has not the correct form, '2006-01-02'", timeStampParts[0]))

		return "", err
	}

	return parserLayout, err
}
