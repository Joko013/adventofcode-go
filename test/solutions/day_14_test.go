package solutions

import (
	"testing"

	"github.com/Joko013/adventofcode-go/solutions"
)

func TestGetSolution14(t *testing.T) {
	inputData := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
	result := 1588

	got := solutions.GetSolution14(inputData)
	if got != result {
		t.Errorf("GetSolution(...) = %d; want: %d", got, result)
	}
}