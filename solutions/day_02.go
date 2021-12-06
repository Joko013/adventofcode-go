package solutions

import (
	// "fmt"
	"strconv"
	"strings"
)

func GetSolution02(inputData string) (result int) {
	depth, horizontal := 0, 0

	commands := strings.Split(inputData, "\n")

	for _, command := range commands {
		slice := strings.Split(command, " ")
		command := slice[0]
		amount, _ := strconv.Atoi(slice[1])

		if command == "up" {
			depth -= amount
		}
		if command == "down" {
			depth += amount
		}
		if command == "forward" {
			horizontal += amount
		}

	}

	return depth * horizontal
}

func GetSolution022(inputData string) (result int) {
	depth, horizontal, aim := 0, 0, 0

	commands := strings.Split(inputData, "\n")

	for _, command := range commands {
		slice := strings.Split(command, " ")
		command := slice[0]
		amount, _ := strconv.Atoi(slice[1])

		if command == "up" {
			aim -= amount
		}
		if command == "down" {
			aim += amount
		}
		if command == "forward" {
			horizontal += amount
			depth += amount * aim
		}
	}
	return depth * horizontal
}