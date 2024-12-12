package day11

import (
	"adventofcode2024/utils"
	"strings"
)

func PartA(filename string) int {
	line := utils.ReadFileAsSingleLine(filename)
	lineArr := strings.Split(line, " ")
	intArray, _ := utils.StringArrayToIntArray(lineArr)
	for index := 0; index < 25; index++ {
		intArray = getNextSetOfNums(intArray)
	}

	return len(intArray)
}

func PartB(filename string) int {
	line := utils.ReadFileAsSingleLine(filename)
	lineArr := strings.Split(line, " ")
	intArray, _ := utils.StringArrayToIntArray(lineArr)
	for index := 0; index < 75; index++ {
		intArray = getNextSetOfNums(intArray)
	}

	return len(intArray)
}

func getNextSetOfNums(intArray []int) []int {
	updatedArray := []int{}
	for _, rockNumber := range intArray {
		rockString := utils.ConvertIntToString(rockNumber)
		rockStringLength := len(rockString)

		if rockNumber == 0 {
			// rock turns into a 1
			updatedArray = append(updatedArray, 1)
		} else if rockStringLength%2 == 0 {
			// even number of digits, split in half
			mid := rockStringLength / 2
			// Return the substring from the start to the midpoint
			firstHalf := rockString[:mid]
			secondHalf := rockString[mid:]
			updatedArray = append(updatedArray, utils.ConvertStringToInt(firstHalf), utils.ConvertStringToInt(secondHalf))
		} else {
			// multiple by 2024
			updatedArray = append(updatedArray, rockNumber*2024)
		}
	}
	return updatedArray
}
