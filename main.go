package main

import (
	"fmt"
	"github.com/Joko013/adventofcode-go/client"
	"github.com/Joko013/adventofcode-go/solutions"
)

func main() {
	day := 6
	part := 2
	inputData, _ := client.GetInputData(day)
	// fmt.Println(inputData)
	solution := solutions.GetSolution06(inputData, 256)
	fmt.Println(solution)
	response, _ := client.SubmitAnswer(day, solution, part)
	fmt.Println(response)
}