package solutions

import (
	"testing"

	"github.com/Joko013/adventofcode-go/solutions"
)

func TestGetSolution10(t *testing.T) {
	inputData := `[({(<(())[]>[[{[]{<()<>>
		[(()[<>])]({[<{<<[]>>(
		{([(<{}[<>[]}>{[]{[(<()>
		(((({<>}<{<{<>}{[]{[]{}
		[[<[([]))<([[{}[[()]]]
		[{[{({}]{}}([{[{{{}}([]
		{<[[]]>}<{[{[{[]{()[[[]
		[<(<(<(<{}))><([]([]()
		<{([([[(<>()){}]>(<<{{
		<{([{{}}[<[[[<>{}]]]>[]]`
	result := 288957

	got := solutions.GetSolution10(inputData)
	if got != result {
		t.Errorf("GetSolution(...) = %d; want: %d", got, result)
	}
}