package solutions

import (
	// "fmt"
	"sort"
	"strings"
	"github.com/Joko013/adventofcode-go/tools"
)

func GetSolution041(inputData string) (score int) {
	split := strings.Split(inputData, "\n\n")
	numbersStr, boardsArrStr := split[0], split[1:]

	numbers := tools.ToIntArray(tools.ToArrayCommaStr(numbersStr))

	boards := getBoards(boardsArrStr)

	for _, number := range numbers {
		toKick := []int{}
		for i, board := range boards {
			for j, line := range board {
				for k, value := range line {
					if (value == number) {
						boards[i][j][k] = 999
						if wins(board) {
							if (len(boards) == 1) {
								return getScore(board, number)
							} else {
								toKick = append(toKick, i)
							}
						}
					}
				}
			}
		}

		// order DESC so you can remove boards without changing the order of the previous items
		sort.Ints(toKick)  // sort ASC
		toKick = reverse(toKick)  // then reverse
		for _, i := range toKick {
			boards = append(boards[:i], boards[i+1:]...)
		}
	}
	return
}

func getBoards(boardsArrStr []string) (boards [][][]int) {
	for _, boardStr := range boardsArrStr {
		lines := strings.Split(boardStr, "\n")

		board := [][]int{}
		for _, line := range lines {
			values := tools.ToArrayWhiteSpaceStr(line)
			board = append(board, tools.ToIntArray(values))
		}
		boards = append(boards, board)
	}
	return boards
}

func wins(board [][]int) (wins bool) {
	for _, row := range board {
		if rowWins(row) {
			return true
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if columnWins(board, i) {
			return true
		}
	}
	return false
}

func rowWins(row []int) (wins bool) {
	for _, value := range row {
		if value != 999 {
			return false
		}
	}
	return true
}

func columnWins(board [][]int, i int) (wins bool) {
	for _, row := range board {
		if row[i] != 999 {
			return false
		}
	}
	return true
}

func getScore(board [][]int, number int) (score int) {
	sum := 0

	for _, row := range board {
		for _, value := range row {
			if (value != 999) {
				sum += value
			}
		}
	}
	return sum * number
}

func reverse(s []int) ([] int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}