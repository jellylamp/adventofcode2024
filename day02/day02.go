package day02

import (
	"adventofcode2024/utils"
	"strings"
)

func PartA(filename string) int {
	input := utils.ReadFileToArrayOfLines(filename)
	totalSafe := 0

	for _, value := range input {
		levelsArr := strings.Split(value, " ")
		isLineSafe := true
		isIncreasing := true

		// set initial increase / decrease expectation once
		firstNum := utils.ConvertStringToInt(levelsArr[0])
		secondNum := utils.ConvertStringToInt(levelsArr[1])

		if firstNum > secondNum {
			isIncreasing = false
		}

		for index, _ := range levelsArr {
			// avoid index out of bounds
			if index+1 >= len(levelsArr) {
				break
			}

			currentNum := utils.ConvertStringToInt(levelsArr[index])
			nextNum := utils.ConvertStringToInt(levelsArr[index+1])

			isLineSafe = checkValidScore(isIncreasing, currentNum, nextNum)
			if !isLineSafe {
				break
			}
		}

		if isLineSafe {
			totalSafe++
		}
	}
	return totalSafe
}

// clearly spiraled out im not refactoring this
func PartB(filename string) int {
	input := utils.ReadFileToArrayOfLines(filename)
	totalSafe := 0

	for _, value := range input {
		levelsArr := strings.Split(value, " ")
		isLineSafe := true
		isIncreasing := isTrendIncreasing(levelsArr)
		hasDampened := false

		for index, _ := range levelsArr {
			// avoid index out of bounds
			if index+1 >= len(levelsArr) {
				break
			}

			currentNum := utils.ConvertStringToInt(levelsArr[index])
			nextNum := utils.ConvertStringToInt(levelsArr[index+1])

			isLevelSafe := checkValidScore(isIncreasing, currentNum, nextNum)
			if !isLevelSafe {
				if hasDampened {
					isLineSafe = false
					break
				}

				// lets check the next num then just in case
				hasDampened = true

				// try removing current and comparing last and next
				// skip the current number and test the previous as long as its within bounds
				if index-1 >= 0 {
					shouldContinue := checkEdgeCase(levelsArr, index-1, index+1, isIncreasing, isLineSafe)
					if shouldContinue {
						restPassed := checkRestOfLevels(levelsArr, isIncreasing, index+1)
						if restPassed {
							isLineSafe = true
							break
						}
					}
				} else {
					// if first number in array, try removing it
					shouldContinue := checkEdgeCase(levelsArr, index+1, index+2, isIncreasing, isLineSafe)
					if shouldContinue {
						restPassed := checkRestOfLevels(levelsArr, isIncreasing, index+1)
						if restPassed {
							isLineSafe = true
							break
						}
					}
				}

				if index+2 > len(levelsArr)-1 {
					// if last number in array, remove it
					isLineSafe = true
					break
				} else {
					// try removing next and comparing current and next + 1
					// skip the next number and test the next to next number; as long as its within bounds
					shouldContinue := checkEdgeCase(levelsArr, index, index+2, isIncreasing, isLineSafe)
					if shouldContinue {
						restPassed := checkRestOfLevels(levelsArr, isIncreasing, index+2)
						if restPassed {
							isLineSafe = true
							break
						}
					}
				}
				// reached this point no line is safe
				isLineSafe = false
				break
			}
		}

		if isLineSafe {
			totalSafe++
		}
	}
	return totalSafe
}

func checkRestOfLevels(levelsArr []string, isIncreasing bool, startingIndex int) bool {
	for index := startingIndex; index < len(levelsArr)-1; index++ {
		// avoid index out of bounds
		if index+1 >= len(levelsArr) {
			break
		}

		currentNum := utils.ConvertStringToInt(levelsArr[index])
		nextNum := utils.ConvertStringToInt(levelsArr[index+1])

		isLevelSafe := checkValidScore(isIncreasing, currentNum, nextNum)
		if !isLevelSafe {
			return false
		}
	}
	return true
}

func checkEdgeCase(levelsArr []string, indexOne int, indexTwo int, isIncreasing bool, isLineSafe bool) bool {
	prevNum := utils.ConvertStringToInt(levelsArr[indexOne])
	nextNum := utils.ConvertStringToInt(levelsArr[indexTwo])

	isLevelSafe := checkValidScore(isIncreasing, prevNum, nextNum)
	if isLevelSafe {
		return true
	}
	return false
}

func isTrendIncreasing(arrToCheck []string) bool {
	violations := 0

	for i := 1; i < len(arrToCheck); i++ {
		if utils.ConvertStringToInt(arrToCheck[i]) < utils.ConvertStringToInt(arrToCheck[i-1]) {
			violations++
			if violations > 1 {
				return false
			}
		}
	}

	return true
}

func checkValidScore(isIncreasing bool, currentNum int, nextNum int) bool {
	var difference = 0
	if isIncreasing {
		difference = nextNum - currentNum
	} else {
		difference = currentNum - nextNum
	}
	return difference >= 1 && difference <= 3
}
