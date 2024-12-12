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
