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
	inX                  bool
}

var directionOffsets = map[Direction]DirectionOffsets{
	NW: {rowOffset: -1, colOffset: -1, inX: true},
	N:  {rowOffset: -1, colOffset: 0, inX: false},
	NE: {rowOffset: -1, colOffset: 1, inX: true},
	E:  {rowOffset: 0, colOffset: 1, inX: false},
	SE: {rowOffset: 1, colOffset: 1, inX: true},
	S:  {rowOffset: 1, colOffset: 0, inX: false},
	SW: {rowOffset: 1, colOffset: -1, inX: true},
	W:  {rowOffset: 0, colOffset: -1, inX: false},
}

var partACount = 0
var partBCount = 0
var partBChecked = make(map[int]map[int]bool)

func PartA(filename string) int {
	grid := utils.ReadFileTo2DArray(filename)

	for rowIndex, row := range grid {
		for colIndex := range row {
			// in this case we can do depth first search without overflows because it will go max 4 deep.
			checkCellXMASLine(grid, rowIndex, colIndex, 0, N)
		}
	}

	return partACount
}

func checkCellXMASLine(grid [][]string, rowIndex int, colIndex int, levelsDeep int, vector Direction) bool {
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
	vectorPathList := getVectorPathList(vector, levelsDeep, grid, rowIndex, colIndex, false)

	for _, vectorValue := range vectorPathList {
		offsets := directionOffsets[vectorValue]
		newRow := rowIndex + offsets.rowOffset
		newCol := colIndex + offsets.colOffset

		// here is where we go through a list of locations in a for loop
		if levelsDeep == 0 {
			// keep digging for the next one; no vector restraint; we know the active letter is X
			checkCellXMASLine(grid, newRow, newCol, 1, vectorValue)
		}
		if currentLetter == "M" && levelsDeep == 1 {
			// keep digging for next one, but follow passed in vector
			checkCellXMASLine(grid, newRow, newCol, 2, vectorValue)
		}
		if currentLetter == "A" && levelsDeep == 2 {
			// keep digging for next one, but follow passed in vector
			checkCellXMASLine(grid, newRow, newCol, 3, vectorValue)
		}
	}

	return false
}

func getVectorPathList(vector Direction, levelsDeep int, grid [][]string, rowIndex int, colIndex int, xOnly bool) []Direction {
	vectorsToCheck := []Direction{}

	if levelsDeep == 0 {
		// check everything given boundaries
		for dir, offsets := range directionOffsets {
			newRow := rowIndex + offsets.rowOffset
			newCol := colIndex + offsets.colOffset

			// Check if the new coordinates are valid
			if utils.IsValid2DIndex(grid, newRow, newCol) {
				if (xOnly && offsets.inX) || !xOnly {
					vectorsToCheck = append(vectorsToCheck, dir)
				}
			}
		}
	} else {
		// keep checking in the direction of the vector (i.e. don't zig zag to complete an answer)
		offsets := directionOffsets[vector]
		newRow := rowIndex + offsets.rowOffset
		newCol := colIndex + offsets.colOffset

		// Check if the new coordinates are valid
		// doesnt need in X already following the vector
		if utils.IsValid2DIndex(grid, newRow, newCol) {
			vectorsToCheck = append(vectorsToCheck, vector)
		}
	}
	return vectorsToCheck
}

func PartB(filename string) int {
	grid := utils.ReadFileTo2DArray(filename)

	for rowIndex, row := range grid {
		for colIndex := range row {
			// in this case we can do depth first search without overflows because it will go max 4 deep.
			checkCellMasX(grid, rowIndex, colIndex, 0, N, "", 0, 0)
		}
	}

	return partBCount
}

func checkCellMasX(grid [][]string, rowIndex int, colIndex int, levelsDeep int, vector Direction, firstLetter string, aRowIndex int, aColIndex int) bool {
	currentLetter := grid[rowIndex][colIndex]

	// quick bail if first letter isn't S or M
	if (currentLetter != "S" && currentLetter != "M") && levelsDeep == 0 {
		return false
	} else if currentLetter != "A" && levelsDeep == 1 {
		return false
	}
	if currentLetter == "S" && firstLetter == "M" && levelsDeep == 2 {
		// thats XMAS! check cross group based on vector
		if checkCrossPath(grid, aRowIndex, aColIndex, vector) {
			partBCount++
			return true
		}
	}
	if currentLetter == "M" && firstLetter == "S" && levelsDeep == 2 {
		// thats XMAS! check cross group based on vector
		if checkCrossPath(grid, aRowIndex, aColIndex, vector) {
			partBCount++
			return true
		}
	}

	// go find the next letter based on valid indices
	vectorPathList := getVectorPathList(vector, levelsDeep, grid, rowIndex, colIndex, true)

	for _, vectorValue := range vectorPathList {
		offsets := directionOffsets[vectorValue]
		newRow := rowIndex + offsets.rowOffset
		newCol := colIndex + offsets.colOffset

		// here is where we go through a list of locations in a for loop
		if levelsDeep == 0 {
			// keep digging for the next one; no vector restraint; we know the active letter is X
			checkCellMasX(grid, newRow, newCol, 1, vectorValue, currentLetter, aRowIndex, aColIndex)
		}
		if currentLetter == "A" && levelsDeep == 1 && !isChecked(rowIndex, colIndex) {
			// mark this A as checked so we don't reuse it on the reverse path
			markChecked(rowIndex, colIndex)

			// keep digging for next one, but follow passed in vector
			checkCellMasX(grid, newRow, newCol, 2, vectorValue, firstLetter, rowIndex, colIndex)
		}
	}

	return false
}

func checkCrossPath(grid [][]string, aRowIndex int, aColIndex int, foundVector Direction) bool {
	if foundVector == NW || foundVector == SE {
		// check NE and SW for MAS or SAM
		neOffsets := directionOffsets[NE]
		neRow := aRowIndex + neOffsets.rowOffset
		neCol := aColIndex + neOffsets.colOffset
		swOffsets := directionOffsets[SW]
		swRow := aRowIndex + swOffsets.rowOffset
		swCol := aColIndex + swOffsets.colOffset

		if grid[neRow][neCol] == "S" && grid[swRow][swCol] == "M" {
			return true
		} else if grid[neRow][neCol] == "M" && grid[swRow][swCol] == "S" {
			return true
		}
	} else if foundVector == SW || foundVector == NE {
		// check NW and SE for MAS or SAM
		nwOffsets := directionOffsets[NW]
		nwRow := aRowIndex + nwOffsets.rowOffset
		nwCol := aColIndex + nwOffsets.colOffset
		seOffsets := directionOffsets[SE]
		seRow := aRowIndex + seOffsets.rowOffset
		seCol := aColIndex + seOffsets.colOffset

		if grid[nwRow][nwCol] == "S" && grid[seRow][seCol] == "M" {
			return true
		} else if grid[nwRow][nwCol] == "M" && grid[seRow][seCol] == "S" {
			return true
		}
	}
	return false
}

func markChecked(x, y int) {
	if _, exists := partBChecked[x]; !exists {
		partBChecked[x] = make(map[int]bool)
	}
	partBChecked[x][y] = true
}

func isChecked(x, y int) bool {
	if _, exists := partBChecked[x]; !exists {
		return false
	}
	return partBChecked[x][y]
}
