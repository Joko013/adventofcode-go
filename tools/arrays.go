package tools

import (
	"strconv"
	"strings"
)

func ToArrayCommaStr (inputStr string) (arr []string) {
	return strings.Split(inputStr, ",")
}

func ToArrayWhiteSpaceStr (inputStr string) (arr []string) {
	return strings.Fields(inputStr)
}

func ToIntArray (strArr []string) (arr []int) {
	for _, x := range strArr {
		value, _ := strconv.Atoi(x)
		arr = append(arr, value)
	}
	return arr
}
