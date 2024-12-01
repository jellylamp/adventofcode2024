package day01

import (
	"adventofcode2024/utils"
	"fmt"
	"strings"
	"sort"
)

var firstDigitList []int
var secondDigitList []int

func parseInput(filename string) {
    input := utils.ReadFileToArrayOfLines(filename)

	for index, value := range input {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
		digitsArr := strings.Split(value, "   ")
		firstDigit := utils.ConvertStringToInt(digitsArr[0])
		secondDigit := utils.ConvertStringToInt(digitsArr[1])
	    firstDigitList = append(firstDigitList, firstDigit)
		secondDigitList = append(secondDigitList, secondDigit)
    }
}

func PartA(filename string)(int) {
	parseInput(filename)

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

func PartB(filename string)(int) {
	parseInput(filename)

	// sort both lists by smallest
    sort.Ints(firstDigitList)
	sort.Ints(secondDigitList)

	countMap := make(map[int]int)

	var totalSimilarityScore = 0

	// add second list to a map
	for _, num := range secondDigitList {
		countMap[num]++
	}

	for index := range firstDigitList {
		similarity := firstDigitList[index] * countMap[firstDigitList[index]]
		totalSimilarityScore += similarity
	}
	return totalSimilarityScore
}

