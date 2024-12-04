package main

import (
	"AdventOfCode2024/util"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func findMatches(line string) [][]string {
	pattern := `mul\(\d{1,},\d{1,}\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(line,-1)

	return matches
}

func doMultiplications(matches [][]string) int {
	total := 0
	for _, match := range matches {
		for _, submatch := range match {
			numsString := strings.ReplaceAll(submatch,"mul(","")
			numsString = strings.ReplaceAll(numsString,")","")
			numsArr := strings.Split(numsString, ",")
			firstNumAsInt, err := strconv.Atoi(numsArr[0])
			if err != nil {
				panic(err)
			}
			secondNumAsInt, err := strconv.Atoi(numsArr[1])
			if err != nil {
				panic(err)
			}
			total += firstNumAsInt * secondNumAsInt
		}
	}
	return total
}

func part1(lines []string){
	total:= 0 

	for _, item := range lines {
		matches := findMatches(item)
		result:= doMultiplications(matches)
		total += result
	}
	fmt.Printf("Part 1: %d\n", total)
}

func removeDontBlocks(input string) string {
	// Define the regex pattern to match "don't()" and the first "do()" inclusively
	pattern := `don't\(\).*?do\(\)`
	re := regexp.MustCompile(pattern)

	// Replace all matches in a loop until no more matches are found
	for re.MatchString(input) {
		input = re.ReplaceAllString(input, "")
	}

	return input
}

func removeIncompleteDontBlocks(input string) string {
	// Define the regex pattern to match "don't()" blocks without a "do()" afterwards
	pattern := `don't\(\).*?(?:(do\(\))|$)`
	re := regexp.MustCompile(pattern)

	// Replace all matches where "do()" is not found
	return re.ReplaceAllStringFunc(input, func(match string) string {
		if !regexp.MustCompile(`do\(\)`).MatchString(match) {
			return ""
		}
		return match
	})
}

func part2(lines []string) {
	// no point parsing, just remove anything between don't() and do()
	line := strings.Join(lines, "")
	newLine := removeDontBlocks(line)
	newLine = removeIncompleteDontBlocks(newLine)
	fmt.Println(newLine)
	matches := findMatches(newLine)
	result := doMultiplications(matches)
	fmt.Printf("Part 2: %d\n", result)
}

func main() {
	lines, err := util.ReadFileLines("./day3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	part1(lines)
	part2(lines)
}



