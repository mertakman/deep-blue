package pathgenerator

import (
	"errors"
)

var ErrInvalidCoordinates = errors.New("given coordinates are in invalid format")

func TextToSquare(loc string) (file, rank int8, err error) {
	if len(loc) != 2 {
		return 0, 0, ErrInvalidCoordinates
	}

	if loc[0] > 64 && loc[0] < 91 { // in case if the file is capital
		file = int8(loc[0]) - 65
	} else if loc[0] > 96 && loc[0] < 123 { // if the capital is lowercase
		file = int8(loc[0]) - 97
	} else {
		return 0, 0, ErrInvalidCoordinates
	}

	if secondChar := int(loc[1]); secondChar > 48 && secondChar < 57 {
		rank = int8(secondChar) - 49 // parses the second character to int8 based from ASCII table .
	} else {
		return 0, 0, ErrInvalidCoordinates
	}

	return file, rank, nil
}
