package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileToArrayOfLines(filename string) []string {
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

	return lines
}

func ReadFileAsSingleLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	var result string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Concatenate lines into a single string with no separators
		result += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return result
}

func ReadFileTo2DArray(filename string) [][]string {
	file, _ := os.Open(filename)
	defer file.Close()

	var array [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		array = append(array, row)
	}

	return array
}

func IsValid2DIndex(grid [][]string, row int, col int) bool {
	// Check if row index is valid
	if row < 0 || row >= len(grid) {
		return false
	}

	// Check if column index is valid for the given row
	if col < 0 || col >= len(grid[row]) {
		return false
	}

	return true
}

func ConvertStringToInt(stringToConvert string) int {
	num, err := strconv.Atoi(stringToConvert)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return num
}

func ConvertStringToInt64(stringToConvert string) int64 {
	num, err := strconv.ParseInt(stringToConvert, 10, 64) // Base 10, 64-bit size
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return num
}

func CheckIntAbsValue(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func IndexOf(list []string, element string) int {
	for i, v := range list {
		if v == element {
			return i
		}
	}
	return -1
}
