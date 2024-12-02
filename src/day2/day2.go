package main

import (
	"AdventOfCode2024/util"
	"fmt"
	"strconv"
	"strings"
)

func safetyTest(line []int) bool{
	decreasing := false
	increasing := false
	for i := range line {
		if i != 0 {
			difference := util.GetDifference(line[i],line[i-1])
			if difference < 1 || difference > 3 {
				return false
			}
			if line[i-1] < line[i] {
				increasing = true
			}
			if line[i-1] > line[i] {
				decreasing = true
			}
			if increasing && decreasing {
				return false
			}
		}
	}
	return true
}

func safetyTestAllowOneError(line []int) bool {
    passed := safetyTest(line)
    if passed {
        return true
    }
    for i := len(line) - 1; i >= 0; i-- {
        newLine, err := util.RemoveByIndex(line, i)
        if err != nil {
            panic(err)
        }
        newLinePassed := safetyTest(newLine)
        if newLinePassed {
            return true
        }
    }
    return false
}


func part1(lines []string) {
	safeCount := 0
	for _, line := range lines {
		linepart := strings.Fields(line)
		var tempArray []int
		for _, num := range linepart {
			numAsInt, err := strconv.Atoi(num)
			if err != nil {
				panic(fmt.Errorf("failed to parse right integer: %w", err))
			}
			tempArray = append(tempArray, numAsInt)
		}
		if safetyTest(tempArray) {
			safeCount = safeCount +1
		}
	}
	fmt.Printf("Part 1: safeCount: %d\n", safeCount)
}

func part2(lines []string) {
	safeCount := 0
	for _, line := range lines {
		linepart := strings.Fields(line)
		var tempArray []int
		for _, num := range linepart {
			numAsInt, err := strconv.Atoi(num)
			if err != nil {
				panic(fmt.Errorf("failed to parse right integer: %w", err))
			}
			tempArray = append(tempArray, numAsInt)
		}
		if safetyTestAllowOneError(tempArray) {
			safeCount = safeCount +1
		}
	}
	fmt.Printf("Part 2: safeCount: %d\n", safeCount)
}

func main(){
	lines, err := util.ReadFileLines("./day2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	part1(lines)
	part2(lines)
}
