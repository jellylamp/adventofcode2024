package utils

import "fmt"

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
	RowOffset, ColOffset int
}

var DirectionOffsetList = map[Direction]DirectionOffsets{
	NW: {RowOffset: -1, ColOffset: -1},
	N:  {RowOffset: -1, ColOffset: 0},
	NE: {RowOffset: -1, ColOffset: 1},
	E:  {RowOffset: 0, ColOffset: 1},
	SE: {RowOffset: 1, ColOffset: 1},
	S:  {RowOffset: 1, ColOffset: 0},
	SW: {RowOffset: 1, ColOffset: -1},
	W:  {RowOffset: 0, ColOffset: -1},
}

func FindStartingChar(grid [][]string, startingChar string) (int, int) {
	for rowIndex, row := range grid {
		for colIndex := range row {
			if grid[rowIndex][colIndex] == startingChar {
				return rowIndex, colIndex
			}
		}
	}

	return 0, 0
}

func MarkGridSpotVisited(x, y int, visitedMap map[int]map[int]bool) map[int]map[int]bool {
	if _, exists := visitedMap[x]; !exists {
		visitedMap[x] = make(map[int]bool)
	}
	visitedMap[x][y] = true
	// return bc pass by value
	return visitedMap
}

func WasVisited(x, y int, visitedMap map[int]map[int]bool) bool {
	if _, exists := visitedMap[x]; !exists {
		return false
	}
	return visitedMap[x][y]
}

func GetValidSurroundingIndices(grid [][]string, rowIndex int, colIndex int) []Direction {
	vectorsToCheck := []Direction{}

	// check everything given boundaries
	for dir, offsets := range DirectionOffsetList {
		newRow := rowIndex + offsets.RowOffset
		newCol := colIndex + offsets.ColOffset

		// Check if the new coordinates are valid
		if IsValid2DIndex(grid, newRow, newCol) {
			vectorsToCheck = append(vectorsToCheck, dir)
		}
	}
	return vectorsToCheck
}

func GetValidIndicesByVector(grid [][]string, rowIndex int, colIndex int, vector Direction) []Direction {
	vectorsToCheck := []Direction{}

	// keep checking in the direction of the vector
	offsets := DirectionOffsetList[vector]
	newRow := rowIndex + offsets.RowOffset
	newCol := colIndex + offsets.ColOffset

	// Check if the new coordinates are valid
	if IsValid2DIndex(grid, newRow, newCol) {
		vectorsToCheck = append(vectorsToCheck, vector)
	}
	return vectorsToCheck
}

func IsValid2DIndex(grid [][]string, row int, col int) bool {
	// Check if row index is valid
	if row < 0 || row >= len(grid) {
		return false
	}

	// Check if column index is valid for the given row
	if col < 0 || col >= len(grid[row]) {
		return false
	}

	return true
}

func TurnRight(vector Direction, row int, col int) (Direction, int, int) {
	// 8 is total number of directions - need this to wrap around
	// enums are just an int so we can add 2 (since we dont want diagonals)
	newVector := (vector + 2) % 8
	offsets := DirectionOffsetList[newVector]
	newRow := row + offsets.RowOffset
	newCol := col + offsets.ColOffset

	return newVector, newRow, newCol
}

func PrintGrid(grid [][]string) {
	fmt.Printf("----------------------------------------------\n")
	for _, row := range grid {
		fmt.Println(row)
	}
}
