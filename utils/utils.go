package utils

import (
	"os"
)

func ReadFileStr(filename string) (string, error) {
	bytes, error := os.ReadFile(filename)
	if error != nil {
		return "", error
	}
	return string(bytes), nil
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
