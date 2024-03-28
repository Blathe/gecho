package utils

import (
	"fmt"
	"strconv"
)

func ToBool(value string) bool {
	v, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Println("error converting form data to bool...")
	}
	return v
}

func ToInt(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("error converting form data to integer...")
	}
	return v
}
