package solutions

import (
	// "fmt"
	"strconv"
	"strings"
)

func GetSolution031(inputData string) (powerConsumption int) {
	reports := strings.Fields(inputData)
	gammaBin, epsilonBin := "", ""
	
	for i := 0; i < len(reports[0]); i++ {
		counts := map[string]int{
			"0": 0,
			"1": 0,
		}

		for _, report := range reports {
			counts[string(report[i])]++
		}

		if counts["0"] > counts["1"] {
			gammaBin += "0"
			epsilonBin += "1"
		} else {
			gammaBin += "1"
			epsilonBin += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gammaBin, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonBin, 2, 64)
	return int(gamma * epsilon)
}


func GetSolution032(inputData string) (lifeSupportRating int) {
	reports := strings.Fields(inputData)
	oxyBin := filter(reports, true)
	co2Bin := filter(reports, false)

	oxy, _ := strconv.ParseInt(oxyBin, 2, 64)
	co2, _ := strconv.ParseInt(co2Bin, 2, 64)
	return int(oxy * co2)
}

func filter(reports []string, mostCommon bool) (report string) {
	candidates := reports

	for i := 0; i < len(reports[0]); i++ {
		// count the bit occurence
		counts := map[string]int{
			"0": 0,
			"1": 0,
		}

		for _, report := range candidates {
			counts[string(report[i])]++
		}

		// select the bit to keep
		keep := "0"
		if counts["1"] >= counts["0"] && mostCommon {
			keep = "1"
		} else if counts["1"] < counts["0"] && !mostCommon {
			keep = "1"
		}

		// filter out the reports with the wrong bit on the current position i
		newCandites := []string{}
		for _, report := range candidates {
			if string(report[i]) == keep {
				newCandites = append(newCandites, report)
			}
		}
		
		// if only one report remains, that's the one
		if len(newCandites) == 1 {
			return newCandites[0]
		} else {
			candidates = newCandites
		}
	}
	return
}
