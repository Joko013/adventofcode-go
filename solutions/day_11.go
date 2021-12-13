package solutions

import (
	"fmt"
	"strings"

	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution11(inputData string) (step int) {
	lines := tools.ToArrayWhiteSpaceStr(inputData)
	
	octopuses := [][]int{}
	for _, line := range lines {
		octopuses = append(octopuses, tools.ToIntArray(strings.Split(line, "")))
	}

	fmt.Println(octopuses)
	for i := 0; i < 1000; i++ {
		flashes := makeStep(&octopuses)
		if flashes == 100 {
			return i + 1
		}
	}

	return
}

func makeStep(octopuses *[][]int) (count int) {
	for i, line := range *octopuses {
		for j := range line {
			iterateOctopus(i, j, octopuses)
		}
	}

	for i, line := range *octopuses {
		for j, energy := range line {
			if energy > 9 {
				(*octopuses)[i][j] = 0
				count++
			}
		}
	}
	return count
}

func iterateOctopus(i int, j int, octopuses *[][]int) (count int) {
	(*octopuses)[i][j]++
	if (*octopuses)[i][j] == 10 {
		count ++

		neighbors := getNeighborsDiag(i, j, *octopuses)
		for _, coordinates := range neighbors {
			k, l := coordinates[0], coordinates[1]
			count += iterateOctopus(k, l, octopuses)		
		}
	}
	return count
}

func getNeighborsDiag(i int, j int, octopuses [][]int) (neighbors [][]int) {
	neighbors = getNeighbors(i, j, octopuses)
	if i > 0 && j > 0 {
		neighbors = append(neighbors, []int{i-1, j-1})
	}
	if i < len(octopuses) - 1 && j < len(octopuses[i]) - 1 {
		neighbors = append(neighbors, []int{i+1, j+1})
	}
	if j > 0 && i < len(octopuses) - 1{
		neighbors = append(neighbors, []int{i+1, j-1})
	}
	if j < len(octopuses[i]) - 1 && i > 0 {
		neighbors = append(neighbors, []int{i-1, j+1})
	}
	return neighbors
}