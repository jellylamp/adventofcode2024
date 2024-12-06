package day05

import (
	"adventofcode2024/utils"
)

var grid [][]string
var visitedMap = make(map[int]map[int]bool)

func PartA(filename string) int {
	grid = utils.ReadFileTo2DArray(filename)
	startingRow, startingCol := utils.FindStartingChar(grid, "^")
	markPositionAndContinue(startingRow, startingCol, utils.N)

	totalCount := 0
	for _, innerMap := range visitedMap {
		totalCount += len(innerMap) // Count keys in each inner map
	}
	return totalCount
}

func markPositionAndContinue(row int, col int, vector utils.Direction) {
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
		if nextCell == "#" {
			// colliion, turn right and keep going
			newVector, newRow, newCol := utils.TurnRight(vector, row, col)
			// utils.PrintGrid(grid)
			markPositionAndContinue(newRow, newCol, newVector)
		} else {
			// keep going; for ease of debugging marking G, not bothering to turn the carat
			nextCell = "G"
			markPositionAndContinue(newRow, newCol, vector)
		}
	} else {
		// out of bounds! Solved the puzzle
		// utils.PrintGrid(grid)
		return
	}
}
