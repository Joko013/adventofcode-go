package solutions

import (
	// "fmt"
	"strconv"
	"strings"
)

func getFishes(inputData string) (map[int]int) {
	fishStrs := strings.Split(inputData, ",")
	fishes := make(map[int]int)

	// make counts of each timer value
	for _, str := range fishStrs {
		value, _ := strconv.Atoi(strings.TrimSpace(str))
		fishes[value] ++
	}

	return fishes
}

func iterate(old map[int]int) (map[int]int) {
	new := make(map[int]int)
	for timer, count := range old {
		if timer == 0 {
			// create new fish and reset timer
			new[8] = count
			new[6] += count
		} else {
			// decrement timer
			new[timer - 1] += count
		}
	}

	return new
}

func GetSolution06(inputData string, iterations int) (count int) {
	fishes := getFishes(inputData)

	for i := 0; i < iterations; i++ {
		fishes = iterate(fishes)
	}
	
	return fishes[0] + fishes[1] + fishes[2] + fishes[3] + fishes[4] + fishes[5] + fishes[6] + fishes[7] + fishes[8]
}
