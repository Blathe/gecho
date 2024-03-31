package utils

import (
	"strconv"
)

// Convert a form value (string) to bool.
func StringToBool(value string) (bool, error) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return v, nil
}

// Convert a form value (string) to int.
func StringToInt(value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return v, nil
}
