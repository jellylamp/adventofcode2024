package day12

import (
	"adventofcode2024/utils"
)

var visitedMap = make(map[utils.Position]bool)
var area = 0
var perimeter = 0

func PartA(filename string) int {
	grid := utils.ReadFileTo2DArray(filename)
	runningWoodPrice := 0

	for rowIndex, row := range grid {
		for colIndex := range row {
			character := grid[rowIndex][colIndex]

			location := utils.Position{Row: rowIndex, Column: colIndex}
			_, exists := visitedMap[location]
			if exists {
				continue
			}

			// recursively search for complete areas; mark locations on the visited map.
			// if visited move on; once we have found the full maps; grab the area and perimeter
			findFullArea(grid, character, rowIndex, colIndex)
			runningWoodPrice += area * perimeter
			area = 0
			perimeter = 0
		}
	}

	return runningWoodPrice
}

func findFullArea(grid [][]string, letter string, currentRow int, currentCol int) {
	// check visited!
	location := utils.Position{Row: currentRow, Column: currentCol}
	_, exists := visitedMap[location]
	if exists {
		return
	} else {
		// mark that we've been here keep checking
		visitedMap[location] = true
		area++
	}

	// check NESW if matching, keep going; else add to perimeter
	checkNorthValue, northRow, northCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.N)
	if checkNorthValue == letter {
		// keep going!
		findFullArea(grid, letter, northRow, northCol)
	} else {
		// this doesn't match but we gain a perimeter
		perimeter++
	}

	checkEastValue, eastRow, eastCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.E)
	if checkEastValue == letter {
		// keep going!
		findFullArea(grid, letter, eastRow, eastCol)
	} else {
		// this doesn't match but we gain a perimeter
		perimeter++
	}

	checkSouthValue, southRow, southCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.S)
	if checkSouthValue == letter {
		// keep going!
		findFullArea(grid, letter, southRow, southCol)
	} else {
		// this doesn't match but we gain a perimeter
		perimeter++
	}

	checkWestValue, westRow, westCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.W)
	if checkWestValue == letter {
		// keep going!
		findFullArea(grid, letter, westRow, westCol)
	} else {
		// this doesn't match but we gain a perimeter
		perimeter++
	}
}

func PartB(filename string) int {
	grid := utils.ReadFileTo2DArray(filename)
	runningWoodPrice := 0

	for rowIndex, row := range grid {
		for colIndex := range row {
			character := grid[rowIndex][colIndex]

			location := utils.Position{Row: rowIndex, Column: colIndex}
			_, exists := visitedMap[location]
			if exists {
				continue
			}

			// recursively search for complete areas; mark locations on the visited map.
			// if visited move on; once we have found the full maps; grab the area and perimeter
			findFullAreaPartB(grid, character, rowIndex, colIndex, utils.N, utils.E)
			runningWoodPrice += area * perimeter
			area = 0
			perimeter = 0
		}
	}

	return runningWoodPrice
}

func findFullAreaPartB(grid [][]string, letter string, currentRow int, currentCol int, previousDirection utils.Direction, currentDirection utils.Direction) {
	// check visited!
	location := utils.Position{Row: currentRow, Column: currentCol}
	_, exists := visitedMap[location]
	if exists {
		return
	} else {
		// mark that we've been here keep checking
		visitedMap[location] = true
		area++
	}

	// check NESW if matching, keep going; else add to perimeter
	checkNorthValue, northRow, northCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.N)
	if checkNorthValue == letter {
		if previousDirection != utils.S {
			// keep going!
			findFullAreaPartB(grid, letter, northRow, northCol, currentDirection, utils.N)
		}
	} else {
		// this doesn't match but we gain a perimeter; but don't want to add if we already have; OR add if we are starting to move back
		if previousDirection != currentDirection || currentDirection == utils.N && previousDirection == utils.N {
			perimeter++
		}
	}

	checkEastValue, eastRow, eastCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.E)
	if checkEastValue == letter {
		if previousDirection != utils.W {
			// keep going!
			findFullAreaPartB(grid, letter, eastRow, eastCol, currentDirection, utils.E)
		}
	} else {
		// this doesn't match but we gain a perimeter; but don't want to add if we already have
		if previousDirection != currentDirection || currentDirection == utils.E && previousDirection == utils.E {
			perimeter++
		}
	}

	checkSouthValue, southRow, southCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.S)
	if checkSouthValue == letter {
		if previousDirection != utils.N {
			// keep going!
			findFullAreaPartB(grid, letter, southRow, southCol, currentDirection, utils.S)
		}
	} else {
		// this doesn't match but we gain a perimeter; but don't want to add if we already have
		if previousDirection != currentDirection || currentDirection == utils.S && previousDirection == utils.S {
			perimeter++
		}
	}

	checkWestValue, westRow, westCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.W)
	if checkWestValue == letter {
		if previousDirection != utils.E {
			// keep going!
			findFullAreaPartB(grid, letter, westRow, westCol, currentDirection, utils.W)
		}
	} else {
		// this doesn't match but we gain a perimeter; but don't want to add if we already have
		if previousDirection != currentDirection || currentDirection == utils.W && previousDirection == utils.W {
			perimeter++
		}
	}
}
