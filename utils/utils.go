package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileToArrayOfLines(filename string) ([]string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(lines)
	
	return lines
}
