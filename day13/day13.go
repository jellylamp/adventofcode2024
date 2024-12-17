package day13

import (
	"adventofcode2024/utils"
	"fmt"
	"regexp"
)

type XYModifiers struct {
	XModifier int
	YModifier int
}

var countMap = make(map[XYModifiers]bool)
var memo = make(map[string]bool)

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
		// clear count map and cache
		countMap = make(map[XYModifiers]bool)
		memo = make(map[string]bool)
	}

	return totalTokensForAllPrizes
}

func binarySearchTreeDFS(targetX, targetY, currentADepth, currentBDepth, runningX, runningY int, buttonA, buttonB XYModifiers) bool {
	// Generate a unique key for the current state
	stateKey := fmt.Sprintf("%d,%d,%d,%d", runningX, runningY, currentADepth, currentBDepth)

	// Check if we have already visited this state
	if result, found := memo[stateKey]; found {
		return result
	}

	// Exit if either depth reaches zero
	if currentADepth == 0 || currentBDepth == 0 {
		memo[stateKey] = false
		return false
	}

	// Exit if we exceed the target
	if runningX > targetX || runningY > targetY {
		memo[stateKey] = false
		return false
	}

	// Success condition: we hit the target
	if runningX == targetX && runningY == targetY {
		countMap[XYModifiers{XModifier: currentADepth, YModifier: currentBDepth}] = true
		memo[stateKey] = true
		return true
	}

	// Recursive calls: press button A or button B
	pressA := binarySearchTreeDFS(targetX, targetY, currentADepth-1, currentBDepth, runningX+buttonA.XModifier, runningY+buttonA.YModifier, buttonA, buttonB)
	pressB := binarySearchTreeDFS(targetX, targetY, currentADepth, currentBDepth-1, runningX+buttonB.XModifier, runningY+buttonB.YModifier, buttonA, buttonB)

	// Cache the result for this state
	memo[stateKey] = pressA || pressB
	return memo[stateKey]
}
