package day10

import (
	"adventofcode2024/utils"
)

func PartA(filename string) int {
	grid := utils.ReadFileTo2DArray(filename)

	// step one find all zeros; then we search!
	startingList := utils.FindStartingPositionList(grid, "0")
	trailheadCount := 0
	for _, startingPosition := range startingList {
		nineMap := make(map[utils.Position]bool)
		// target answer will increase
		OrthagonalSearchTreeDFS(grid, startingPosition.Row, startingPosition.Column, 0, 1, -1, nineMap)
		trailheadCount += len(nineMap)
	}

	return trailheadCount
}

func OrthagonalSearchTreeDFS(grid [][]string, row int, col int, currentNumber int, lookingFor int, previousNumber int, nineMap map[utils.Position]bool) bool {
	shouldReturn := false

	// if 9 and previous was 8 ; add to the map
	if currentNumber == 9 && previousNumber == 8 {
		// found one! return
		nineMap[utils.Position{Row: row, Column: col}] = true
		return true
	}

	// check NESW; if 1, continue down and increment number; else return
	checkNorthValue, northRow, northCol := utils.GetValidIntByVector(grid, row, col, utils.N)
	if checkNorthValue == lookingFor {
		// keep going!
		if OrthagonalSearchTreeDFS(grid, northRow, northCol, lookingFor, lookingFor+1, currentNumber, nineMap) {
			shouldReturn = true
		}
	}

	checkEastValue, eastRow, eastCol := utils.GetValidIntByVector(grid, row, col, utils.E)
	if checkEastValue == lookingFor {
		// keep going!
		if OrthagonalSearchTreeDFS(grid, eastRow, eastCol, lookingFor, lookingFor+1, currentNumber, nineMap) {
			shouldReturn = true
		}
	}

	checkSouthValue, southRow, southCol := utils.GetValidIntByVector(grid, row, col, utils.S)
	if checkSouthValue == lookingFor {
		// keep going!
		if OrthagonalSearchTreeDFS(grid, southRow, southCol, lookingFor, lookingFor+1, currentNumber, nineMap) {
			shouldReturn = true
		}
	}

	checkWestValue, westRow, westCol := utils.GetValidIntByVector(grid, row, col, utils.W)
	if checkWestValue == lookingFor {
		// keep going!
		if OrthagonalSearchTreeDFS(grid, westRow, westCol, lookingFor, lookingFor+1, currentNumber, nineMap) {
			shouldReturn = true
		}
	}

	return shouldReturn
}
