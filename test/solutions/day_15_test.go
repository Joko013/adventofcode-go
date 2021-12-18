package solutions

import (
	"testing"

	"github.com/Joko013/adventofcode-go/solutions"
)

func TestGetSolution15(t *testing.T) {
	inputData := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`
	result := 315

	got := solutions.GetSolution15(inputData)
	if got != result {
		t.Errorf("GetSolution(...) = %d; want: %d", got, result)
	}
}