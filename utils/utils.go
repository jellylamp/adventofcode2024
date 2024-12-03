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
