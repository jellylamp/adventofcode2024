package day07

import (
	"adventofcode2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func PartA(filename string) int {
	input := utils.ReadFileToArrayOfLines(filename)
	totalCount := 0

	for _, line := range input {
		// split on :
		componentParts := strings.Split(line, ":")
		answer := utils.ConvertStringToInt(componentParts[0])
		// split on " " and get rid of the first one
		numsToTest := strings.Split(componentParts[1], " ")[1:]

		canMakeAnswer := binarySearchTreeDFS(answer, numsToTest, 0, "", 0)
		if canMakeAnswer {
			totalCount += answer
		}
	}

	return totalCount
}

func PartB(filename string) int {
	input := utils.ReadFileToArrayOfLines(filename)
	totalCount := 0

	for _, line := range input {
		// split on :
		componentParts := strings.Split(line, ":")
		answer := utils.ConvertStringToInt(componentParts[0])
		// split on " " and get rid of the first one
		numsToTest := strings.Split(componentParts[1], " ")[1:]

		canMakeAnswer := ternarySearchTreeDFS(answer, numsToTest, 0, "", 0)
		if canMakeAnswer {
			totalCount += answer
		}
	}

	return totalCount
}

func binarySearchTreeDFS(targetAnswer int, numsToTest []string, currentDepth int, operator string, runningNumber int) bool {
	currentNum := utils.ConvertStringToInt(numsToTest[currentDepth])

	// add or multiple based upon operator
	if operator == "+" {
		runningNumber += currentNum
	} else if operator == "*" {
		runningNumber = runningNumber * currentNum
	} else {
		// initial case, set running num to value
		runningNumber = currentNum
	}

	// check if we are at the bottom
	if currentDepth == len(numsToTest)-1 {
		return runningNumber == targetAnswer
	}

	currentDepth++
	if binarySearchTreeDFS(targetAnswer, numsToTest, currentDepth, "+", runningNumber) {
		return true
	}

	if binarySearchTreeDFS(targetAnswer, numsToTest, currentDepth, "*", runningNumber) {
		return true
	}

	return false
}

func ternarySearchTreeDFS(targetAnswer int, numsToTest []string, currentDepth int, operator string, runningNumber int) bool {
	currentNum := utils.ConvertStringToInt(numsToTest[currentDepth])

	// add or multiple based upon operator
	if operator == "+" {
		runningNumber += currentNum
	} else if operator == "*" {
		runningNumber = runningNumber * currentNum
	} else if operator == "||" {
		var err error
		runningNumber, err = strconv.Atoi(fmt.Sprintf("%d%d", runningNumber, currentNum))
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		// initial case, set running num to value
		runningNumber = currentNum
	}

	// check if we are at the bottom
	if currentDepth == len(numsToTest)-1 {
		return runningNumber == targetAnswer
	}

	currentDepth++
	if ternarySearchTreeDFS(targetAnswer, numsToTest, currentDepth, "+", runningNumber) {
		return true
	}

	if ternarySearchTreeDFS(targetAnswer, numsToTest, currentDepth, "*", runningNumber) {
		return true
	}

	if ternarySearchTreeDFS(targetAnswer, numsToTest, currentDepth, "||", runningNumber) {
		return true
	}

	return false
}
