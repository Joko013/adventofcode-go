package solutions

import (
	// "fmt"
	"strings"
	"strconv"
)

func GetSolution011 (inputData string) (count int) {
	measurements := strings.Fields(inputData)

	lastValue, _ := strconv.Atoi(measurements[0])

	for _, v := range measurements[1:] {
		newValue, _ := strconv.Atoi(v)

		if (newValue > lastValue) {
			count ++
		}
		lastValue = newValue
	}

	return count
}

func GetSolution012 (inputData string) (count int) {
	measurementsStrs := strings.Fields(inputData)

	var measurements []int
	for _, x := range measurementsStrs {
		value, _ := strconv.Atoi(x)
		measurements = append(measurements, value)
	}

	lastSum := measurements[0] + measurements[1] + measurements[2]

	for i := 0; i < len(measurements) - 2; i++ {
		newSum := measurements[i] + measurements[i + 1] + measurements[i + 2]
		if (newSum > lastSum) {
			count ++
		}

		lastSum = newSum
	}

	return count
}