package solutions

import (
	// "fmt"
	"math"
	"strconv"

	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution15(inputData string) (risk int) {
	m := [][]int{}
	for _, line := range tools.ToArrayWhiteSpaceStr(inputData) {
		l := []int{}
		for _, r := range line {
			r, _ := strconv.Atoi(string(r))
			l = append(l, r)
		}
		m = append(m, l)
	}

	m = expandMap(m)  // part 2

	// djikstra it!
	visited := map[int]map[int]bool{}
	distances := map[int]map[int]int{}
	for i, l := range m {
		for j := range l {
			if distances[i] == nil {
				distances[i] = map[int]int{j: math.MaxInt}
			} else {
				distances[i][j] = math.MaxInt
			}
		}
	}
	distances[0][0] = 0 // starting point

	current := []int{0, 0} // start at the top-left node
	// keep looping over nodes until the end node is reached
	for {
		currentX, currentY := current[0], current[1]
		for _, n := range getNeighbors(currentX, currentY, m) {
			x, y := n[0], n[1]
			if !visited[x][y] { // only consider unvisited neighbors
				newDistance := distances[currentX][currentY] + m[x][y]

				if newDistance < distances[x][y] { // if the path is shorter, use it
					distances[x][y] = newDistance
				}
			}
		}

		// mark the node as visited
		if visited[currentX] == nil {
			visited[currentX] = map[int]bool{currentY: true}
		} else {
			visited[currentX][currentY] = true
		}

		// select the lowest unvisited node as the next node
		next := []int{}
		nextMin := math.MaxInt
		for i, l := range distances {
			for j, d := range l {
				if d < nextMin && !visited[i][j] {
					next = []int{i, j}
					nextMin = d
				}
			}
		}

		// end when the bottom right node is reached
		if next[0] == len(m) - 1 && next[1] == len(m[0]) - 1 {
			return distances[next[0]][next[1]]
		}
		// or just go to next node
		current = next
	}
}

func expandMap(m [][]int) (bigMap [][]int) {
	wideMap := [][]int{}
	for _, line := range m { // expand horizontally
		bigLine := line
		for i := 1; i < 5; i++ {
			bigLine = append(bigLine, increaseRow(line, i)...)
		}
		wideMap = append(wideMap, bigLine)
	}
	
	bigMap = wideMap
	for i := 1; i < 5; i++ {
		for _, bigLine := range wideMap { // expand vertically	
			bigMap = append(bigMap, increaseRow(bigLine, i))
		}
	}
	return bigMap
}

func increaseRow(row []int, i int) (bigRow []int) {
	for _, item := range row {
		value := item + i
		if value > 9 {
			value -= 9
		}
		bigRow = append(bigRow, value)
	}
	return bigRow
}