package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func ConvertStringToInt(stringToConvert string)(int) {
	num, err := strconv.Atoi(stringToConvert)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return num
}

func CheckIntAbsValue(num int)(int) {
	if num < 0 {
		num = -num
	}
	return num
}