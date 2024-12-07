package day06

import (
	"adventofcode2024/utils"
)

var grid [][]string
var visitedMap = make(map[string]bool)

func PartA(filename string) int {
	grid = utils.ReadFileTo2DArray(filename)
	startingRow, startingCol := utils.FindStartingChar(grid, "^")
	markPositionAndContinue(startingRow, startingCol, utils.N)

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
			// no need to test for a spot thats already blocked
			if grid[rowIndex][col] == "#" || grid[rowIndex][col] == "^" {
				continue
			}

			// set obstacle
			grid[rowIndex][col] = "0"

			isInfiniteLoop := markPositionAndContinue(startingRow, startingCol, utils.N)
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

func markPositionAndContinue(row int, col int, vector utils.Direction) bool {
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

		if nextCell == "#" || nextCell == "0" {
			newVector, newRow, newCol := utils.TurnRight(vector, row, col)
			nextCell = grid[newRow][newCol]

			// we turned right... into a wall
			if nextCell == "#" || nextCell == "0" {
				newVector, newRow, newCol = utils.TurnRight(newVector, row, col)
				nextCell = grid[newRow][newCol]

				// we turned right... into ANOTHER wall; just exit back the way you came
				if nextCell == "#" || nextCell == "0" {
					newVector, newRow, newCol = utils.TurnRight(newVector, row, col)
				}

			}

			return markPositionAndContinue(newRow, newCol, newVector)
		}

		return markPositionAndContinue(newRow, newCol, vector)
	}

	return false
}
