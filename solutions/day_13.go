package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution13(inputData string) (result int) {
	split := strings.Split(inputData, "\n\n")
	dotsStr, foldsStr := tools.ToArrayWhiteSpaceStr(split[0]), strings.Split(split[1],  "\n")

	dots := [][]int{}
	for _, line := range dotsStr {
		dots = append(dots, tools.ToIntArray(tools.ToArrayCommaStr(line)))
	}

	foldInstructions := []FoldInstruction{}
	for _, line := range foldsStr {
		splitLine := strings.Split(line, "=")
		axis := string(splitLine[0][len(splitLine[0])-1])
		index, _ := strconv.Atoi(splitLine[1])
		fold := FoldInstruction{
			axis: axis,
			index: index,
		}
		foldInstructions = append(foldInstructions, fold)
	}

	for _, fi := range foldInstructions {
		dots = fold(dots, fi)
	}

	printCode(dots)
	return len(dots)
}

type FoldInstruction struct {
	axis string
	index int
}

func fold(dots [][]int, fi FoldInstruction) (folded [][]int) {
	for _, dot := range dots {
		var newDot []int

		if fi.axis == "x" {
			if dot[0] > fi.index {
				newDot = []int{fi.index - (dot[0] - fi.index), dot[1]}
			} else {
				newDot = dot
			}
		} else {
			if dot[1] > fi.index {
				newDot = []int{dot[0], fi.index - (dot[1] - fi.index)}
			} else {
				newDot = dot
			}
		}
		folded = addDot(folded, newDot)
	}

	return folded
}

func addDot(dots [][]int, dot []int) (updated [][]int) {
	if inDots(dot, dots) {
		return dots
	}
	return append(dots, dot)
}

func inDots(d []int, dots [][]int) (is bool) {
	for _, dot := range dots {
		if d[0] == dot[0] && d[1] == dot[1] {
			return true
		}
	}
	return false
}

func printCode(dots[][]int) {
	for i := 0; i < 40; i++ {
		line := ""
		for j := 0; j < 40; j++ {
			dot := []int{i, j}
			if inDots(dot, dots) {
				line = line + "#"
			} else {
				line = line + " "
			}
		}
		fmt.Println(line)
	}
}
