package day09

import (
	"adventofcode2024/utils"
	"strconv"
)

func PartA(filename string) int {
	line := utils.ReadFileAsSingleLine(filename)

	idVal := 0
	uncompactedList := []string{}
	periodLocationList := []int{}

	for index, characterAsRune := range line {
		character := string(characterAsRune)
		// loop through character by character
		isOdd := index%2 != 0

		if isOdd {
			// blank space X times
			xTimes := utils.ConvertStringToInt(character)
			for i := 0; i < xTimes; i++ {
				uncompactedList = append(uncompactedList, ".")
				// keep track of where periods are so we can more easily traverse the list again
				periodLocationList = append(periodLocationList, len(uncompactedList)-1)
			}
		} else {
			// add id in character num times
			xTimes := utils.ConvertStringToInt(character)
			for i := 0; i < xTimes; i++ {
				uncompactedList = append(uncompactedList, strconv.Itoa(idVal))
			}

			// increment after
			idVal++
		}
	}

	// make a copy so we have uncompacted list for debugging
	compactedList := make([]string, len(uncompactedList))
	copy(compactedList, uncompactedList)
	compactedListLength := len(compactedList)
	stopLooping := false

	// loop through period list so we can compact and replace
	for index := 0; index < len(periodLocationList); index++ {
		indexToReplace := periodLocationList[index]

		// loop through the end to find the last digit in the list
		for endIndex := compactedListLength - 1; endIndex >= 0; endIndex-- {
			var locationOfLastDigit int
			if compactedList[endIndex] != "." {
				// wooo a digit
				locationOfLastDigit = endIndex
			} else {
				// keep searchin
				continue
			}

			if locationOfLastDigit > periodLocationList[index] {
				// first index of a period, replace it with the last number of the array
				compactedList[indexToReplace] = compactedList[locationOfLastDigit]
				compactedList[locationOfLastDigit] = "."
				break
			} else {
				// we have fully compressed
				stopLooping = true
				break
			}
		}
		if stopLooping {
			break
		}
	}

	checksum := 0
	// 0 * 1 + 11 * 1 + 11 * 2 + 11 * 3 + 11 * 4 + 11 * 5 + 1 * 6 + 2 * 7 + 3 * 8 + 4 * 9 + 5 * 10 + 6 * 11 + 7 * 12 + 8 * 13 + 9 * 14 + 10 * 15
	// loop through compacted list and make a checksum!
	for index, value := range compactedList {
		if value != "." {
			bitChecksum := index * utils.ConvertStringToInt(value)
			checksum += bitChecksum
		}
	}
	return checksum
}
