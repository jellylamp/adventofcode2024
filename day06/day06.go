package day05

import (
	"adventofcode2024/utils"
	// "fmt"
)

var grid [][]string
var visitedMap = make(map[int]map[int]bool)

func PartA(filename string) int {
	grid = utils.ReadFileTo2DArray(filename)
	startingRow, startingCol := utils.FindStartingChar(grid, "^")
	markPositionAndContinue(startingRow, startingCol, utils.N, false, 0)

	totalCount := 0
	for _, innerMap := range visitedMap {
		totalCount += len(innerMap) // Count keys in each inner map
	}
	return totalCount
}

func PartB(filename string) int {
	totalCount := 0
	grid = utils.ReadFileTo2DArray(filename)
	originalGrid := utils.ResetGrid(grid)

	startingRow, startingCol := utils.FindStartingChar(grid, "^")

	// try putting an obstacle in every single spot in the grid
	for rowIndex, row := range originalGrid {
		for col := range row {
			// if rowIndex == 4 && col == 5 {
			// 	debugHere := 0
			// 	debugHere++
			// }

			// set obstacle
			grid[rowIndex][col] = "O"

			// in this case we can do depth first search without overflows because it will go max 4 deep.
			isInfiniteLoop := markPositionAndContinue(startingRow, startingCol, utils.N, false, 0)
			if isInfiniteLoop {

				// utils.PrintGrid(grid)
				// fmt.Println(rowIndex, col)
				totalCount++
			}

			// put grid and map back to original
			grid = utils.ResetGrid(originalGrid)
			visitedMap = make(map[int]map[int]bool)
		}
	}

	return totalCount
}

func markPositionAndContinue(row int, col int, vector utils.Direction, previousWasVisited bool, previousWasVisitedCount int) bool {
	backOutReturn := false

	// start checking if we have two repeats in a row which would imply we have an infinite loop
	wasVisited := utils.WasVisited(row, col, visitedMap)
	if wasVisited && previousWasVisited {
		previousWasVisitedCount++
	}
	if !wasVisited {
		previousWasVisitedCount = 0
	}
	if previousWasVisitedCount > 800 {
		return true
	}

	// first mark the visited spot with an X for ease of debugging; also make a running map for ease of checking/counting
	grid[row][col] = "X"
	visitedMap = utils.MarkGridSpotVisited(row, col, visitedMap)

	// keep checking in the direction of the vector
	offsets := utils.DirectionOffsetList[vector]
	newRow := row + offsets.RowOffset
	newCol := col + offsets.ColOffset

	// Check if the new coordinates are valid
	if utils.IsValid2DIndex(grid, newRow, newCol) {
		nextCell := grid[newRow][newCol]
		// they are valid! check if collision and if so turn right
		if nextCell == "#" || nextCell == "O" {
			// collision, turn right and keep going
			newVector, newRow, newCol := utils.TurnRight(vector, row, col)
			backOutReturn = markPositionAndContinue(newRow, newCol, newVector, wasVisited, previousWasVisitedCount)
		} else {
			// keep going; for ease of debugging marking G, not bothering to turn the carat
			nextCell = "G"
			backOutReturn = markPositionAndContinue(newRow, newCol, vector, wasVisited, previousWasVisitedCount)
		}
	} else {
		return false
	}
	return backOutReturn
}
