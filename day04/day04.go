package day04

import (
	"adventofcode2024/utils"
)

type Direction int

const (
	NW Direction = iota
	N
	NE
	E
	SE
	S
	SW
	W
)

type DirectionOffsets struct {
	rowOffset, colOffset int
}

var directionOffsets = map[Direction]DirectionOffsets{
	NW: {rowOffset: -1, colOffset: -1},
	N:  {rowOffset: -1, colOffset: 0},
	NE: {rowOffset: -1, colOffset: 1},
	E:  {rowOffset: 0, colOffset: 1},
	SE: {rowOffset: 1, colOffset: 1},
	S:  {rowOffset: 1, colOffset: 0},
	SW: {rowOffset: 1, colOffset: -1},
	W:  {rowOffset: 0, colOffset: -1},
}

var partACount = 0

func PartA(filename string) int {
	grid := utils.ReadFileTo2DArray(filename)

	for rowIndex, row := range grid {
		for colIndex, _ := range row {
			// in this case we can do depth first search without overflows because it will go max 4 deep.
			checkCell(grid, rowIndex, colIndex, 0, N)
		}
	}

	return partACount
}

func checkCell(grid [][]string, rowIndex int, colIndex int, levelsDeep int, vector Direction) bool {
	currentLetter := grid[rowIndex][colIndex]
	// quick bail if letter isn't even X
	if currentLetter != "X" && levelsDeep == 0 {
		return false
	} else if currentLetter != "M" && levelsDeep == 1 {
		return false
	} else if currentLetter != "A" && levelsDeep == 2 {
		return false
	} else if currentLetter != "S" && levelsDeep == 3 {
		return false
	}
	if currentLetter == "S" && levelsDeep == 3 {
		// thats XMAS! Return true
		partACount++
		return true
	}

	// go find the next letter based on valid indices
	vectorPathList := getVectorPathList(vector, levelsDeep, grid, rowIndex, colIndex)

	for _, vectorValue := range vectorPathList {
		offsets := directionOffsets[vectorValue]
		newRow := rowIndex + offsets.rowOffset
		newCol := colIndex + offsets.colOffset

		// nextLetter := grid[newRow][newCol]

		// here is where we go through a list of locations in a for loop
		if levelsDeep == 0 {
			// keep digging for the next one; no vector restraint; we know the active letter is X
			checkCell(grid, newRow, newCol, 1, vectorValue)
		}
		if currentLetter == "M" && levelsDeep == 1 {
			// keep digging for next one, but follow passed in vector
			checkCell(grid, newRow, newCol, 2, vectorValue)
		}
		if currentLetter == "A" && levelsDeep == 2 {
			// keep digging for next one, but follow passed in vector
			checkCell(grid, newRow, newCol, 3, vectorValue)
		}
	}

	return false
}

func getVectorPathList(vector Direction, levelsDeep int, grid [][]string, rowIndex int, colIndex int) []Direction {
	vectorsToCheck := []Direction{}

	if levelsDeep == 0 {
		// check everything given boundaries
		for dir, offsets := range directionOffsets {
			newRow := rowIndex + offsets.rowOffset
			newCol := colIndex + offsets.colOffset

			// Check if the new coordinates are valid
			if utils.IsValid2DIndex(grid, newRow, newCol) {
				vectorsToCheck = append(vectorsToCheck, dir)
			}
		}
	} else {
		// keep checking in the direction of the vector (i.e. don't zig zag to complete an answer)
		offsets := directionOffsets[vector]
		newRow := rowIndex + offsets.rowOffset
		newCol := colIndex + offsets.colOffset

		// Check if the new coordinates are valid
		if utils.IsValid2DIndex(grid, newRow, newCol) {
			vectorsToCheck = append(vectorsToCheck, vector)
		}
	}
	return vectorsToCheck
}
