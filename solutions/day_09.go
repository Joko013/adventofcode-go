package solutions

import (
	// "fmt"
	"sort"
	"strings"

	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution09(inputData string) (total int) {
	lines := tools.ToArrayWhiteSpaceStr(inputData)

	heightmap := [][]int{}
	for _, line := range lines {
		heightmap = append(heightmap, tools.ToIntArray(strings.Split(line, "")))
	}

	lengths := []int{}
	for i, line := range heightmap {
		for j := range line {
			if isLowest(i, j, heightmap) {
				coordinates := []int{i, j}
				basin := findBasin(i, j, heightmap, [][]int{coordinates})
				lengths = append(lengths, len(basin))
			}
		}
	}

	sort.Ints(lengths)
	l := len(lengths)
	return lengths[l - 1] * lengths[l - 2] * lengths[l - 3]
}

func getNeighbors(i int, j int, heightmap [][]int) (neighbors [][]int) {
	if i > 0 {
		neighbors = append(neighbors, []int{i-1, j})
	}
	if i < len(heightmap) - 1 {
		neighbors = append(neighbors, []int{i+1, j})
	}
	if j > 0 {
		neighbors = append(neighbors, []int{i, j-1})
	}
	if j < len(heightmap[i]) - 1 {
		neighbors = append(neighbors, []int{i, j+1})
	}
	return neighbors
}

func isLowest(i int, j int, heightmap [][]int) (bool) {
	for _, coordinates := range getNeighbors(i, j, heightmap) {
		k, l := coordinates[0], coordinates[1]
		if heightmap[i][j] >= heightmap[k][l] {
			return false
		}
	}
	return true
}

func findBasin(i int, j int, heightmap [][]int, basin [][]int) ([][]int) {
	for _, coordinates := range getNeighbors(i, j, heightmap) {
		k, l := coordinates[0], coordinates[1]
		if isNewBasinPart(k, l, heightmap, basin) {
			basin = findBasin(k, l, heightmap, append(basin, []int{k, l}))
		}
	}
	return basin
}

func isNewBasinPart(i int, j int, heightmap [][]int, basin [][]int ) (bool) {
	if alreadyInBasin([]int{i, j}, basin) || heightmap[i][j] > 8 {
		return false
	}
	return true
}

func alreadyInBasin(item []int, basin [][]int) (bool) {
	for _, v := range basin {
		if item[0] == v[0] && item[1] == v[1] {
			return true
		}
	}
	return false
}