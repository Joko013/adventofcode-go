package solutions

import (
	"fmt"
	"math"
	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution07(inputData string) (totalFuel int) {
	positions := tools.ToIntArray(tools.ToArrayCommaStr(inputData))
	fmt.Println(positions)
	min, max := getMinMax(positions)
	
	totalFuel = math.MaxInt
	for n := min; n <= max; n++ {
		consumption := 0
		for _, position := range positions {
			difference := int(math.Abs(float64(n - position)))
			// consumption += difference
			consumption += difference * (difference + 1) / 2  // sum of N + N-1 + ... + 1
		}
		
		if consumption < totalFuel {
			totalFuel = consumption
		}
	}

	return totalFuel
}

func getMinMax(s []int) (min int, max int) {
	min, max = s[0], s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	return min, max
}