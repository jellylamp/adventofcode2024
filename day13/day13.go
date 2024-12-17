package day13

import (
	"adventofcode2024/utils"
	"regexp"
)

type XYModifiers struct {
	XModifier int
	YModifier int
}

var countMap = make(map[XYModifiers]bool)

func PartA(filename string) int {
	lineArr := utils.ReadFileToArrayOfLines(filename)
	// Regular expression to match numbers
	re := regexp.MustCompile(`\d+`)
	totalTokensForAllPrizes := 0

	for index := 0; index < len(lineArr); index = index + 4 {
		// index 0 is button a
		// index 1 is button b
		// index 2 is prize string
		// index 3 is space
		buttonAString := lineArr[index]
		buttonANums := re.FindAllString(buttonAString, -1)
		buttonA := XYModifiers{XModifier: utils.ConvertStringToInt(buttonANums[0]), YModifier: utils.ConvertStringToInt(buttonANums[1])}

		buttonBString := lineArr[index+1]
		buttonBNums := re.FindAllString(buttonBString, -1)
		buttonB := XYModifiers{XModifier: utils.ConvertStringToInt(buttonBNums[0]), YModifier: utils.ConvertStringToInt(buttonBNums[1])}

		prizeString := lineArr[index+2]
		prizeNums := re.FindAllString(prizeString, -1)
		prizeX := utils.ConvertStringToInt(prizeNums[0])
		prizeY := utils.ConvertStringToInt(prizeNums[1])

		// feed values into BST - once we have equaled or gone past X and Y values, then we should break that chain.
		// If we do find an equal, add to map and we will calculate smallest amount of tokens later
		binarySearchTreeDFS(prizeX, prizeY, 100, 100, 0, 0, buttonA, buttonB)

		// calculate which count is most efficient
		lowestTokenTotal := 0
		for key := range countMap {
			// a costs 3
			// b costs 1
			tokenTotal := (100-key.XModifier)*3 + (100-key.YModifier)*1
			if tokenTotal < lowestTokenTotal || lowestTokenTotal == 0 {
				lowestTokenTotal = tokenTotal
			}
		}
		totalTokensForAllPrizes += lowestTokenTotal
		// clear count map
		countMap = make(map[XYModifiers]bool)
	}

	return totalTokensForAllPrizes
}

func binarySearchTreeDFS(targetX int, targetY int, currentADepth int, currentBDepth int, runningX int, runningY int, buttonA XYModifiers, buttonB XYModifiers) bool {
	// Immediately exit if either depth is zero
	if currentADepth == 0 || currentBDepth == 0 {
		return false
	}

	// Exit if we've exceeded the target
	if runningX > targetX || runningY > targetY {
		return false
	}

	// Success condition
	if runningX == targetX && runningY == targetY {
		countMap[XYModifiers{XModifier: currentADepth, YModifier: currentBDepth}] = true
		return true
	}

	// Only attempt recursive calls if depths allow
	if currentADepth > 0 && binarySearchTreeDFS(targetX, targetY, currentADepth-1, currentBDepth, runningX+buttonA.XModifier, runningY+buttonA.YModifier, buttonA, buttonB) {
		return true
	}

	if currentBDepth > 0 && binarySearchTreeDFS(targetX, targetY, currentADepth, currentBDepth-1, runningX+buttonB.XModifier, runningY+buttonB.YModifier, buttonA, buttonB) {
		return true
	}

	return false
}
