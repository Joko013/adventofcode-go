package solutions

import (
	"testing"

	"github.com/Joko013/adventofcode-go/solutions"
)

func TestGetSolution12(t *testing.T) {
	inputData := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	result := 36

	got := solutions.GetSolution12(inputData)
	if got != result {
		t.Errorf("GetSolution(...) = %d; want: %d", got, result)
	}
}
