package solutions

import (
	"fmt"
	"math"
	"strings"
)

func GetSolution14(inputData string) (result int){
	split := strings.Split(inputData, "\n\n")
	str, rest := split[0], split[1]
	rulesStr := strings.Split(rest, "\n")

	rules := map[string]string{}
	for _, r := range rulesStr {
		s := strings.Split(r, " -> ")
		rules[s[0]] = s[1]
	}

	pairs := map[string]int{}
	for i := 0; i < len(str) - 1; i++ {
		pair := str[i:i+2]
		if pairs[pair] == 0 {
			pairs[pair] = 1
		} else {
			pairs[pair]++
		}
	}

	for i := 0; i < 40; i++ {
		pairs = makeInsertionStep(pairs, rules)
	}
	
	counts := map[string]int{}
	for pair, count := range pairs {
		for _, rune := range pair {
			char := string(rune)
			if counts[char] == 0 {
				counts[char] = count
			} else {
				counts[char] += count
			}
		}
	}
	
	min, max := math.MaxInt, 0

	for k, v := range counts {
		if k == string(str[0]) || k == string(str[len(str)-1]) {
			// starting and ending letters of the string
			v++
		}
		fmt.Println(k, v/2)
		vMin, vMax := v / 2, v / 2
		if vMin < min {
			min = vMin
		} else if vMax > max {
			max = vMax
		}
	}
	return max - min
}

func makeInsertionStep(pairs map[string]int, rules map[string]string) (newPairs map[string]int) {
	newPairs = map[string]int{}
	for pair, count := range pairs {
		newLetter := rules[pair]
		if newLetter == "" {
			continue
		}
		items := []string{string(pair[0]) + newLetter, newLetter + string(pair[1])}

		for _, item := range items {
			if newPairs[item] == 0 {
				newPairs[item] = count
			} else {
				newPairs[item] += count
			}
		}
	}
	return newPairs
}