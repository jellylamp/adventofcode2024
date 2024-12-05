package day05

import (
	"adventofcode2024/utils"
	"fmt"
	"strings"
)

func PartA(filename string) int {
	input := utils.ReadFileToArrayOfLines(filename)
	orderMap := make(map[string][]string)
	inOrderingRules := true
	middleSums := 0

	for _, line := range input {
		// loop through input and add to ordering rules
		// split the line before | after
		// add to map, key is the number to look up
		if line == "" {
			inOrderingRules = false
			continue
		}
		if inOrderingRules {
			splitList := strings.Split(line, "|")
			beforeNum := splitList[0]
			afterNum := splitList[1]
			orderMap[beforeNum] = append(orderMap[beforeNum], afterNum)
		} else {
			// put update instructions in its own list; in case we need it
			splitArr := strings.Split(line, ",")
			if isInstructionCorrect(splitArr, orderMap) {
				middleSums += findMiddleAndConvertToNum(splitArr)
			}
		}
	}

	return middleSums
}

func PartB(filename string) int {
	input := utils.ReadFileToArrayOfLines(filename)
	orderMap := make(map[string][]string)
	inOrderingRules := true
	middleSums := 0

	for _, line := range input {
		// loop through input and add to ordering rules
		// split the line before | after
		// add to map, key is the number to look up
		if line == "" {
			inOrderingRules = false
			continue
		}
		if inOrderingRules {
			splitList := strings.Split(line, "|")
			beforeNum := splitList[0]
			afterNum := splitList[1]
			orderMap[beforeNum] = append(orderMap[beforeNum], afterNum)
		} else {
			// put update instructions in its own list; in case we need it
			splitArr := strings.Split(line, ",")
			if !isInstructionCorrect(splitArr, orderMap) {
				reorderedArr := reorderInstructions(splitArr, orderMap)
				if isInstructionCorrect(reorderedArr, orderMap) {
					middleSums += findMiddleAndConvertToNum(reorderedArr)
				} else {
					fmt.Println("ruh roh")
				}
			}
		}
	}

	return middleSums
}

func reorderInstructions(splitArr []string, orderMap map[string][]string) []string {
	// ok so we have a map of correct instructions

	updatedOrderList := make([]string, len(splitArr))
	copy(updatedOrderList, splitArr)

	// loop over list and make a secondary list running
	for _, runningSplitArrValue := range splitArr {
		// take first item, check its correctness. if wrong, move index
		orderRules := orderMap[runningSplitArrValue]
		for _, orderRule := range orderRules {
			// if ordered rule is in the splitArr, make sure its after value index
			orderedRuleIndexInSplitArr := utils.IndexOf(updatedOrderList, orderRule)
			// get the latest order, in case it moved
			runningSplitIndexInOrderList := utils.IndexOf(updatedOrderList, runningSplitArrValue)

			// if orderedRuleIndexInSplitArr is -1, then this rule isnt in split arr
			if orderedRuleIndexInSplitArr != -1 && orderedRuleIndexInSplitArr <= runningSplitIndexInOrderList {

				// Remove the element from the "from" index
				updatedOrderList = append(updatedOrderList[:runningSplitIndexInOrderList], updatedOrderList[runningSplitIndexInOrderList+1:]...)

				orderedRuleIndexInOrderList := utils.IndexOf(updatedOrderList, orderRule)
				// Insert the element at the correct position
				updatedOrderList = append(updatedOrderList[:orderedRuleIndexInOrderList], append([]string{runningSplitArrValue}, updatedOrderList[orderedRuleIndexInOrderList:]...)...)

			}
		}
	}
	return updatedOrderList
}

func isInstructionCorrect(splitArr []string, orderMap map[string][]string) bool {
	// loop over splitArr, grabbing each character and making sure that the items behind it are allowed
	for valueIndex, value := range splitArr {
		valueRules := orderMap[value]
		for _, valueRule := range valueRules {
			// if valueRule is in the splitArr, make sure its after value index
			updateOrder := utils.IndexOf(splitArr, valueRule)
			if updateOrder != -1 && updateOrder <= valueIndex {
				// incorrect
				return false
			}
		}
	}
	return true
}

// im assuming odd length lines only
func findMiddleAndConvertToNum(array []string) int {
	length := len(array)

	middleVal := array[length/2]
	return utils.ConvertStringToInt(middleVal)
}
