package solutions

import (
	"sort"

	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution10(inputData string) (score int) {
	lines := tools.ToArrayWhiteSpaceStr(inputData)
	
	points := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	scores := []int{}
	for _, line := range lines {
		score := 0
		illegalChar, openingChars := checkLine(line)
		if illegalChar == "" {
			for i := len(openingChars) - 1; i >= 0; i -- {
				score *= 5
				score += points[openingChars[i]]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[(len(scores)/2)]
}

func checkLine(line string) (illegalChar string, openingChars []string) {
	pairs := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	for _, rune := range line {
		c := string(rune)
		if isOpeningChar(c, pairs) {
			openingChars = append(openingChars, c)
		} else if pairs[openingChars[len(openingChars) - 1]] == c {
			openingChars = openingChars[:len(openingChars) - 1]
		} else {
			return c, openingChars
		}
			
	}

	return "", openingChars
}


func isOpeningChar(c string, pairs map[string]string) (bool) {
	for key := range pairs {
		if c == key {
			return true
		}
	}
	return false
}
