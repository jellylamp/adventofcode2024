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

			// b one too many perimeter

			// recursively search for complete areas; mark locations on the visited map.
			// if visited move on; once we have found the full maps; grab the area and perimeter
			findFullAreaPartB(grid, character, rowIndex, colIndex, utils.N, -1, -1, rowIndex, colIndex)
			runningWoodPrice += area * perimeter
			area = 0
			perimeter = 0
		}
	}

	return runningWoodPrice
}

type Fence struct {
	fenceDirection utils.Direction
	row            int
	col            int
	letter         string
}

var fenceMap = make(map[Fence]bool)

func findFullAreaPartB(grid [][]string, letter string, currentRow int, currentCol int, currentDirection utils.Direction, previousRow int, previousCol int, startingRow int, startingCol int) {
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
	if checkNorthValue != letter {
		// when I would normally add a perimeter; I should instead first check if I already have one in the previous location and vector
		_, previousFenceExists := fenceMap[Fence{fenceDirection: utils.N, row: previousRow, col: previousCol, letter: letter}]
		loopedBackToStarter := false
		if northRow == startingRow && northCol == startingCol {
			loopedBackToStarter = true
		}

		if !previousFenceExists && !loopedBackToStarter {
			perimeter++
		}
		fenceMap[Fence{fenceDirection: utils.N, row: currentRow, col: currentCol, letter: letter}] = true
	}

	checkEastValue, eastRow, eastCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.E)
	if checkEastValue != letter {
		// when I would normally add a perimeter; I should instead first check if I already have one in the previous location and vector
		_, previousFenceExists := fenceMap[Fence{fenceDirection: utils.E, row: previousRow, col: previousCol, letter: letter}]
		loopedBackToStarter := false
		if eastRow == startingRow && eastCol == startingCol {
			loopedBackToStarter = true
		}

		if !previousFenceExists && !loopedBackToStarter {
			perimeter++
		}
		fenceMap[Fence{fenceDirection: utils.E, row: currentRow, col: currentCol, letter: letter}] = true
	}

	checkSouthValue, southRow, southCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.S)
	if checkSouthValue != letter {
		// when I would normally add a perimeter; I should instead first check if I already have one in the previous location and vector
		_, previousFenceExists := fenceMap[Fence{fenceDirection: utils.S, row: previousRow, col: previousCol, letter: letter}]
		loopedBackToStarter := false
		if southRow == startingRow && southCol == startingCol {
			loopedBackToStarter = true
		}

		if !previousFenceExists && !loopedBackToStarter {
			perimeter++
		}
		fenceMap[Fence{fenceDirection: utils.S, row: currentRow, col: currentCol, letter: letter}] = true
	}

	checkWestValue, westRow, westCol := utils.GetValidStringByVector(grid, currentRow, currentCol, utils.W)
	if checkWestValue != letter {
		// when I would normally add a perimeter; I should instead first check if I already have one in the previous location and vector
		_, previousFenceExists := fenceMap[Fence{fenceDirection: utils.W, row: previousRow, col: previousCol, letter: letter}]
		loopedBackToStarter := false
		if westRow == startingRow && westCol == startingCol {
			loopedBackToStarter = true
		}
		// b is broken here because west row == 2  and west col == -1
		// compare starting row/column/direction? no of course that will always exist it has to match what WE have
		if !previousFenceExists && !loopedBackToStarter {
			perimeter++
		}
		fenceMap[Fence{fenceDirection: utils.W, row: currentRow, col: currentCol, letter: letter}] = true
	}

	/////////////
	if checkNorthValue == letter {
		// don't go backwards down the same route
		// if previousDirection != utils.S {
		// keep going!
		findFullAreaPartB(grid, letter, northRow, northCol, utils.N, currentRow, currentCol, startingRow, startingCol)
		// }
	}

	if checkEastValue == letter {
		// if previousDirection != utils.W {
		// keep going!
		findFullAreaPartB(grid, letter, eastRow, eastCol, utils.E, currentRow, currentCol, startingRow, startingCol)
		// }
	}

	if checkSouthValue == letter {
		// if previousDirection != utils.N {
		// keep going!
		findFullAreaPartB(grid, letter, southRow, southCol, utils.S, currentRow, currentCol, startingRow, startingCol)
		// }
	}

	if checkWestValue == letter {
		// if previousDirection != utils.E {
		// keep going!
		findFullAreaPartB(grid, letter, westRow, westCol, utils.W, currentRow, currentCol, startingRow, startingCol)
		// }
	}
}
