package util

import (
	"bufio"
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
