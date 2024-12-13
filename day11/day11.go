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

type NumberInfo struct {
	IsSingleNumber bool
	Integer1       int
	Integer2       int
}

var lookupMap = make(map[int]NumberInfo)

func getNextSetOfNums(intArray []int) []int {
	updatedArray := []int{}
	for _, rockNumber := range intArray {
		// quick lookup if we've seen this number before
		rockLookup, exists := lookupMap[rockNumber]
		if exists {
			if rockLookup.IsSingleNumber {
				updatedArray = append(updatedArray, rockLookup.Integer1)
			} else {
				updatedArray = append(updatedArray, rockLookup.Integer1)
				updatedArray = append(updatedArray, rockLookup.Integer2)
			}
			continue
		}

		// haven't seen it keep going
		rockString := utils.ConvertIntToString(rockNumber)
		rockStringLength := len(rockString)

		if rockNumber == 0 {
			// rock turns into a 1
			lookupMap[rockNumber] = NumberInfo{IsSingleNumber: true, Integer1: 1, Integer2: -1}
			updatedArray = append(updatedArray, 1)
		} else if rockStringLength%2 == 0 {
			// even number of digits, split in half
			mid := rockStringLength / 2
			// Return the substring from the start to the midpoint
			firstHalf := utils.ConvertStringToInt(rockString[:mid])
			secondHalf := utils.ConvertStringToInt(rockString[mid:])
			updatedArray = append(updatedArray, firstHalf, secondHalf)
			lookupMap[rockNumber] = NumberInfo{IsSingleNumber: false, Integer1: firstHalf, Integer2: secondHalf}
		} else {
			// multiple by 2024
			updatedArray = append(updatedArray, rockNumber*2024)
			lookupMap[rockNumber] = NumberInfo{IsSingleNumber: true, Integer1: rockNumber * 2024, Integer2: -1}
		}
	}
	return updatedArray
}

type EndCount struct {
	depth  int
	number int
}

var endCountMap = make(map[EndCount]int)

func PartB(filename string) int {
	line := utils.ReadFileAsSingleLine(filename)
	lineArr := strings.Split(line, " ")
	intArray, _ := utils.StringArrayToIntArray(lineArr)
	total := 0

	for _, number := range intArray {
		total += getNextSetOfNumsRecursion(number, 75)
	}

	return total
}

func getNextSetOfNumsRecursion(rockNumber int, depth int) int {
	rockString := utils.ConvertIntToString(rockNumber)
	rockStringLength := len(rockString)
	var total int

	if depth == 0 {
		return 1
	}

	// if we know the total, dont do the loop
	rockLookup, exists := endCountMap[EndCount{depth: depth, number: rockNumber}]
	if exists {
		return rockLookup
	} else {
		if rockNumber == 0 {
			// rock turns into a 1
			total = getNextSetOfNumsRecursion(1, depth-1)
		} else if rockStringLength%2 == 0 {
			// even number of digits, split in half
			mid := rockStringLength / 2
			// Return the substring from the start to the midpoint
			firstHalf := utils.ConvertStringToInt(rockString[:mid])
			secondHalf := utils.ConvertStringToInt(rockString[mid:])
			total = getNextSetOfNumsRecursion(firstHalf, depth-1) + getNextSetOfNumsRecursion(secondHalf, depth-1)
		} else {
			// multiple by 2024
			total = getNextSetOfNumsRecursion(rockNumber*2024, depth-1)
		}

		// track totals returned, so we can cut off the whole tree if theres a repeat at the same depth
		endCountMap[EndCount{depth: depth, number: rockNumber}] = total
	}
	return total
}
