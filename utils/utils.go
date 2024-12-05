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

func IsCharNum(c byte) bool {
	if c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
		return true
	}
	return false
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func SliceContains[T comparable](s []T, o T) (int, bool) {
	for i, v := range s {
		if v == o {
			return i, true
		}
	}

	return 0, false
}
