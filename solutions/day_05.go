package solutions

import (
	"fmt"
	"strings"

	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution05(inputData string) (countOverTwo int) {
	linesStr := strings.Split(inputData, "\n")
	fmt.Println(linesStr[len(linesStr)-1])

	grid := make(map[int]map[int]int)
	for _, lineStr := range linesStr {
		startEnd := strings.Split(lineStr, " -> ")
		startStr, endStr := startEnd[0], startEnd[1]
		start := tools.ToIntArray(tools.ToArrayCommaStr(startStr))
		end := tools.ToIntArray(tools.ToArrayCommaStr(endStr))
		
		if start[0] == end[0] || start[1] == end[1] {
			// horizontal and vertical lines
			if start[0] > end[0] || start[1] > end[1] {
				// order the coordinates so the value always increases from 1 to 2
				start, end = end, start
			}
	
			if start[0] != end[0] {
				// vertical line
				for i := start[0]; i <= end[0]; i++ {
					if grid[i] == nil {
						grid[i] = map[int]int{start[1]: 1}
					} else {
						grid[i][start[1]]++
					}
				}
			} else {
				// horizontal line
				for j := start[1]; j <= end[1]; j++ {
					if grid[start[0]] == nil {
						grid[start[0]] = map[int]int{j: 1}
					} else {
						grid[start[0]][j]++
					}
				}
			}
		} else {
			// diagonal line
			if start[0] > end[0] {
				// to make sure the X value increases during the iteration
				start, end = end, start
			}

			if start[1] < end[1] {
				for i := 0; i <= end[0] - start[0]; i++ {
					if grid[start[0] + i] == nil {
						grid[start[0] + i] = map[int]int{start[1] + i: 1}
					} else {
						grid[start[0] + i][start[1] + i]++
					}
				}
			} else {
				for i := 0; i <= end[0] - start[0]; i++ {
					if grid[start[0] + i] == nil {
						grid[start[0] + i] = map[int]int{start[1] - i: 1}
					} else {
						grid[start[0] + i][start[1] - i]++
					}
				}
			}
		}
	}
	// fmt.Println(grid)

	for _, x := range grid {
		for _, y := range x {
			if y > 1 {
				countOverTwo++
			}
		}
	}

	return countOverTwo
}
