package day05

import (
	"adventofcode2024/utils"
	"strings"
)

func PartA(filename string) int {
	input := utils.ReadFileToArrayOfLines(filename)
	orderMap := make(map[string][]string)
	updateList := []string{}

	inOrderingRules := true
	middleSums := 0
	correctRules := [][]string{}

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
			updateList = append(updateList, line)
			splitArr := strings.Split(line, ",")
			if isInstructionCorrect(splitArr, orderMap) {
				correctRules = append(correctRules, splitArr)
				middleSums += findMiddleAndConvertToNum(splitArr)
			}
		}
	}

	return middleSums
}

func isInstructionCorrect(splitArr []string, orderMap map[string][]string) bool {
	// loop over splitArr, grabbing each character and making sure that the items behind it are allowed

	// The fifth update, 61,13,29, is also not in the correct order, since it breaks the rule 29|13.
	// The last update, 97,13,75,29,47, is not in the correct order due to breaking several rules.

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
