package day02

import (
	"adventofcode2024/utils"
	"strings"
)



func PartA(filename string)(int) {
	input := utils.ReadFileToArrayOfLines(filename)
	totalSafe := 0

	for _, value := range input {
		levelsArr := strings.Split(value, " ")
		isLineSafe := true
		isIncreasing := true

		// set initial increase / decrease expectation once
		firstNum := utils.ConvertStringToInt(levelsArr[0])
		secondNum := utils.ConvertStringToInt(levelsArr[1])

		if (firstNum > secondNum) {
			isIncreasing = false
		}

		for index, _ := range levelsArr {
			// avoid index out of bounds
			if index + 1 >= len(levelsArr) {
				break
			}

			currentNum := utils.ConvertStringToInt(levelsArr[index])
			nextNum := utils.ConvertStringToInt(levelsArr[index + 1])		
			
			isLineSafe = checkValidScore(isIncreasing, currentNum, nextNum)
			if (!isLineSafe) {
				break
			}
		}

		if (isLineSafe) {
			totalSafe++
		}
    }
	return totalSafe
}

func checkValidScore(isIncreasing bool, currentNum int, nextNum int)(bool) {
	var difference = 0
	if (isIncreasing) {
		difference = nextNum - currentNum
	} else {
		difference = currentNum - nextNum
	}
	return difference >= 1 && difference <= 3
}