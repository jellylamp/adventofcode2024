package day01

import (
	"adventofcode2024/utils"
	"fmt"
	"strings"
	"sort"
)

func Run(filename string)(int) {
    input := utils.ReadFileToArrayOfLines(filename)
	var firstDigitList []int
	var secondDigitList []int
	for index, value := range input {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
		digitsArr := strings.Split(value, "   ")
		firstDigit := utils.ConvertStringToInt(digitsArr[0])
		secondDigit := utils.ConvertStringToInt(digitsArr[1])
	    firstDigitList = append(firstDigitList, firstDigit)
		secondDigitList = append(secondDigitList, secondDigit)
    }

	// sort both lists by smallest
    sort.Ints(firstDigitList)
	sort.Ints(secondDigitList)
	var totalDistanceCalc = 0

	for index := range firstDigitList {
		distanceCalc := utils.CheckIntAbsValue(firstDigitList[index] - secondDigitList[index])
		totalDistanceCalc += distanceCalc
	}
	return totalDistanceCalc
}
