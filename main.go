package main

import (
	"fmt"
	"github.com/Joko013/adventofcode-go/client"
	"github.com/Joko013/adventofcode-go/solutions"
)

func main() {
	day := 15
	inputData, _ := client.GetInputData(day)
	// fmt.Println(inputData)
	solution := solutions.GetSolution15(inputData)
	fmt.Println(solution)
	part := 2
	response, _ := client.SubmitAnswer(day, solution, part)
	fmt.Println(response)
}