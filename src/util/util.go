package util

import (
	"bufio"
	"regexp"
	"fmt"
	"os"
)

// ReadFileLines takes a file path and returns a slice of strings containing each line of the file,
// or an error if something goes wrong.
func ReadFileLines(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Initialize a slice to hold the lines
	var fileLines []string

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return fileLines, nil
}

func RemoveDuplicates(input []int) []int {
	set := make(map[int]struct{}) // Simulated set
	unique := []int{}

	for _, val := range input {
		if _, exists := set[val]; !exists {
			set[val] = struct{}{}
			unique = append(unique, val)
		}
	}

	return unique
}

func SumArray(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// This one was a real pain in the hole, you cant modify the existing slice
// for the purposes of assigning it to a new var, you need to create an
// entirely new slice, then the correct slice, with the element removed 
// will be returned
func RemoveByIndex(input []int, index int) ([]int, error) {
    if index < 0 || index >= len(input) {
        return nil, fmt.Errorf("index out of range")
    }
    // Create a new slice by appending parts of the original
    result := append([]int{}, input[:index]...)
    result = append(result, input[index+1:]...)
    return result, nil
}

func GetDifference(first, second int) int {
		if first > second {
		return first - second
} else if second > first {
		return second- first
	} else {
		return 0
	}
}

// RemoveBetween removes content between two delimiters (inclusive).
func RemoveBetween(input, start, end string) (string, error) {
	// Create a regex pattern to match everything between the start and end strings
	pattern := fmt.Sprintf(`%s.*?%s`, regexp.QuoteMeta(start), regexp.QuoteMeta(end))

	// Compile the regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile regex: %w", err)
	}

	// Replace all matches with an empty string
	result := re.ReplaceAllString(input, "")
	return result, nil
}









