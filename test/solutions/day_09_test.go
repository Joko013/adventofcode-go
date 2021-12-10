package solutions

import (
	"testing"

	"github.com/Joko013/adventofcode-go/solutions"
)

func TestGetSolution09(t *testing.T) {
	inputData := `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`
	result := 1134

	got := solutions.GetSolution09(inputData)
	if got != result {
		t.Errorf("GetSolution(...) = %d; want: %d", got, result)
	}
}
