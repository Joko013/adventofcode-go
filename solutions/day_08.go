package solutions

import (
	// "fmt"
	"strconv"
	"strings"

	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution08(inputData string) (sum int) {
	// fmt.Println(inputData)
	newLineSplit := strings.Split(inputData, "\n")

	for _, line := range newLineSplit {
		pipeSplit := strings.Split(line, " | ")
		patterns := tools.ToArrayWhiteSpaceStr(pipeSplit[0])
		outputValues := tools.ToArrayWhiteSpaceStr(pipeSplit[1])

		mapping := getMapping(patterns)
		
		strValue := ""
		for _, v := range outputValues {
			for key, element := range mapping {
				if len(v) == len(element) && containsAll(element, v) {
					strValue += strconv.Itoa(key)
				}
			}
		}
		value, _ := strconv.Atoi(strValue)
		sum += value
	}

	return sum
}

func getMapping(patterns []string) (mapping map[int]string) {
	mapping = map[int]string{}

	toMap := []string{}
	// get initial mapping from the unique values
	for _, v := range patterns {
		length := len(v)
		if length == 2 {
			mapping[1] = v
		} else if length == 3 {
			mapping[7] = v
		} else if length == 4 {
			mapping[4] = v
		} else if length == 7 {
			mapping[8] = v
		} else {
			toMap = append(toMap, v)
		}
	}

	for len(toMap) > 0 {
		remains := []string{}
		// deduce from the shape of the number
		for _, v := range toMap {
			length := len(v)
			if length == 5 {
				if containsAll(v, mapping[1]) {
					// 3 is the only 5-long digit that contains the same segments as 1
					mapping[3] = v
				} else if mapping[9] != "" && containsAll(mapping[9], v) {
					// 5 looks like a 9 with one segment cut off
					mapping[5] = v
				} else if mapping[3] != "" && mapping[5] != "" {
					// and 2 remains
					mapping[2] = v
				} else {
					remains = append(remains, v)
				}
			} else {
				// only 6-long segments remains
				if containsAll(v, mapping[4]) {
					// 9 is the only 6-long digit that contains the same segments as 4
					mapping[9] = v
				} else if mapping[9] != "" && containsAll(v, mapping[1]) {
					// then 0 remains as the only one to contain the same segments as 1
					mapping[0] = v
				} else if mapping[9] != "" && mapping[0] != "" {
					// and 6 remains
					mapping[6] = v
				} else {
					remains = append(remains, v)
				}
			}
		}
		toMap = remains
	}

	return mapping
}

func containsAll(s string, chars string) (bool) {
	for _, c := range chars {
		if !strings.Contains(s, string(c)) {
			return false
		}
	}
	return true
}
