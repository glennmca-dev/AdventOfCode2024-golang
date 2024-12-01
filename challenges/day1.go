package main

import (
	"AdventOfCode2024/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var leftList []int
var rightList []int
var lines []string
var differences []int
var totalDifference int 

func getDifference(first, second int) int {
	if(first > second) {
		return first - second
	} else if(second > first) {
		return second - first
	} else {
		return 0
	}
}

func part1(lines []string) {
	for _, line := range lines {
	// for i, line := range lines {
		// fmt.Printf("Line %d: %s\n", i+1, line)
		lineParts := strings.Fields(line)
    leftInt, err := strconv.Atoi(lineParts[0])
		if err != nil {
			panic(fmt.Errorf("failed to parse left integer: %w", err))
		}

		rightInt, err := strconv.Atoi(lineParts[1])
		if err != nil {
			panic(fmt.Errorf("failed to parse right integer: %w", err))
		}
		leftList	=	append(leftList, leftInt)
		rightList	=	append(rightList, rightInt)
	}
	sort.Ints(leftList)
	sort.Ints(rightList)

  for i  := range leftList	{

		leftInt := leftList[i]
		rightInt := rightList[i]
		totalDifference = totalDifference + getDifference(leftInt,rightInt)
	} 

	fmt.Printf("totalDifference: %d\n", totalDifference)
}

func part2() {
	leftListNoDupes := util.RemoveDuplicates(leftList) 
  var similarities []int
	for _, item := range leftListNoDupes {
		var similarityScore int
		numAppearances := 0
		for _, checkItem := range rightList {
			if checkItem == item {
				numAppearances = numAppearances + 1
			}	
		}
		similarityScore =	item * numAppearances
		similarities = append(similarities, similarityScore)
	}
	fmt.Printf("similarityScore: %d\n", util.SumArray(similarities))
}

func main() {
	lines, err := util.ReadFileLines("/home/glenn/repos/AdventOfCode2024-golang/challenges/day1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	part1(lines)
	part2()


	
}





















