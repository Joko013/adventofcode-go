package solutions

import (
	// "fmt"
	"strings"
)

func GetSolution12(inputData string) (count int) {
	neighbours := map[string][]string{}
	for _, c := range strings.Split(inputData, "\n"){
		conn := strings.Split(c, "-")
		neighbours[conn[0]] = append(neighbours[conn[0]], conn[1])
		neighbours[conn[1]] = append(neighbours[conn[1]], conn[0])
	}

	paths := getPaths("start", neighbours, []string{"start"})
	return len(paths)
}

func getPaths(cave string, neighbours map[string][]string, visited []string) (paths []string) {
	anyCaveVisitedTwice := twiceVisitedExists(visited)
	for _, n := range neighbours[cave] {
		if n == "end" {
			paths = append(paths, cave + n)
		} else if isInvalid(n, visited, anyCaveVisitedTwice) {
			continue
		} else {
			nPaths := getPaths(n, neighbours, append(visited, n))
			for _, p := range nPaths {
				paths = append(paths, cave + p)
			}
		}

	}
	return paths
}

func isInvalid(cave string, visited []string, anyCaveVisitedTwice bool) (bool) {
	for _, c := range visited {
		if (cave == c && strings.ToLower(cave) == cave) {
			if anyCaveVisitedTwice || cave == "start" {
				return true
			}
		}
	}
	return false
}

func twiceVisitedExists(visited []string) (bool) {
	m := map[string]int{}
	for _, v := range visited {
		m[v]++
	}

	for k, v := range m {
		if v > 1 && strings.ToLower(k) == k {
			return true
		}
	}
	return false
}
