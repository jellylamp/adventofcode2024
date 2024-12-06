package day05

import (
	"adventofcode2024/utils"
	// "fmt"
)

var grid [][]string
var visitedMap = make(map[string]bool)

func PartA(filename string) int {
	grid = utils.ReadFileTo2DArray(filename)
	startingRow, startingCol := utils.FindStartingChar(grid, "^")
	markPositionAndContinue(startingRow, startingCol, utils.N, false, 0)

	totalCount := utils.CountUniqueCoordinates(visitedMap)
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
				totalCount++
			}

			// put grid and map back to original
			grid = utils.ResetGrid(originalGrid)
			visitedMap = make(map[string]bool)
		}
	}

	return totalCount
}

func markPositionAndContinue(row int, col int, vector utils.Direction, previousWasVisited bool, previousWasVisitedCount int) bool {
	wasVisited := utils.WasVisited(row, col, visitedMap, vector)

	if wasVisited {
		return true
	}

	// Mark visited spot
	grid[row][col] = "X"
	visitedMap = utils.MarkGridSpotVisited(row, col, visitedMap, vector)

	offsets := utils.DirectionOffsetList[vector]
	newRow := row + offsets.RowOffset
	newCol := col + offsets.ColOffset

	if utils.IsValid2DIndex(grid, newRow, newCol) {
		nextCell := grid[newRow][newCol]

		if nextCell == "#" || nextCell == "O" {
			newVector, newRow, newCol := utils.TurnRight(vector, row, col)
			return markPositionAndContinue(newRow, newCol, newVector, wasVisited, previousWasVisitedCount)
		}

		return markPositionAndContinue(newRow, newCol, vector, wasVisited, previousWasVisitedCount)
	}

	return false
}
