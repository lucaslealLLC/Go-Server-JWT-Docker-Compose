package common

import (
	"errors"
	"regexp"
)

func AlphaWithSpaces(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("not a string")
	}
	re := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if re.MatchString(str) {
		return nil
	}
	return errors.New("string passed is not alpha")
}

func IdNumeric(value interface{}) error {

	switch t := value.(type) {

	case string:
		re := regexp.MustCompile(`^[0-9 ]*$`)
		if re.MatchString(t) {
			return nil
		}
		return errors.New("not numerical")

	case int, uint, float64:
		return nil

	default:
		return errors.New("not numerical")
	}
}
