package day08

import (
	"adventofcode2024/utils"
)

type Position struct {
	Row    int
	Column int
}

func PartA(filename string) int {
	grid := utils.ReadFileTo2DArray(filename)
	antennaMap := make(map[string][]Position)
	positionMap := make(map[Position]bool)

	for rowIndex, row := range grid {
		for colIndex := range row {
			character := grid[rowIndex][colIndex]
			if character != "." {
				antennaMap[character] = append(antennaMap[character], Position{Row: rowIndex, Column: colIndex})
			}
		}
	}

	// we now have a map of all antenna and their locations
	for _, antenna := range antennaMap {
		// grab all indices and lets compare them
		// edge cases! antenna at a spot one already is; antenna would be placed but off the map
		for positionListIndex, position := range antenna {
			// look at every pair AFTER the current item (dont compare to yourself or previously compared items)
			for positionListSearchIndex := positionListIndex + 1; positionListSearchIndex < len(antenna); positionListSearchIndex++ {
				positionToCheck := antenna[positionListSearchIndex]

				// get the difference in indices between position and positionToCheck
				manhattanRowDistance := utils.CheckIntAbsValue(position.Row - positionToCheck.Row)
				manhattanColDistance := utils.CheckIntAbsValue(position.Column - positionToCheck.Column)

				// Determine the direction
				topAntennaRowDir, topAntennaColDir := getDirectionForPosition(position, positionToCheck)
				bottomAntennaRowDir, bottomAntennaColDir := getDirectionForPosition(positionToCheck, position)

				// take that distance and grab the top and bottom indexes in the grid that the antenna would go
				topAntenna := Position{Row: position.Row + (topAntennaRowDir * manhattanRowDistance), Column: position.Column + (topAntennaColDir * manhattanColDistance)}
				if utils.IsValid2DIndex(grid, topAntenna.Row, topAntenna.Column) {
					// add it to a map of indices we've found
					positionMap[topAntenna] = true
				}

				bottomAntenna := Position{Row: positionToCheck.Row + (bottomAntennaRowDir * manhattanRowDistance), Column: positionToCheck.Column + (bottomAntennaColDir * manhattanColDistance)}
				if utils.IsValid2DIndex(grid, bottomAntenna.Row, bottomAntenna.Column) {
					// add it to a map of indices we've found
					positionMap[bottomAntenna] = true
				}
			}
		}
	}

	return len(positionMap)
}

func getDirectionForPosition(position1 Position, position2 Position) (int, int) {
	rowDirection := 1
	if position1.Row < position2.Row {
		rowDirection = -1
	}

	colDirection := 1
	if position1.Column < position2.Column {
		colDirection = -1
	}
	return rowDirection, colDirection
}

func PartB(filename string) int {
	grid := utils.ReadFileTo2DArray(filename)
	antennaMap := make(map[string][]Position)
	positionMap := make(map[Position]bool)

	for rowIndex, row := range grid {
		for colIndex := range row {
			character := grid[rowIndex][colIndex]
			if character != "." {
				antennaMap[character] = append(antennaMap[character], Position{Row: rowIndex, Column: colIndex})
			}
		}
	}

	// we now have a map of all antenna and their locations
	for _, antenna := range antennaMap {
		totalCountForLetter := 0
		// grab all indices and lets compare them
		// edge cases! antenna at a spot one already is; antenna would be placed but off the map
		for positionListIndex, topPosition := range antenna {
			// look at every pair AFTER the current item (dont compare to yourself or previously compared items)
			for positionListSearchIndex := positionListIndex + 1; positionListSearchIndex < len(antenna); positionListSearchIndex++ {
				bottomPosition := antenna[positionListSearchIndex]

				// get the difference in indices between position and positionToCheck
				manhattanRowDistance := utils.CheckIntAbsValue(topPosition.Row - bottomPosition.Row)
				manhattanColDistance := utils.CheckIntAbsValue(topPosition.Column - bottomPosition.Column)

				// Determine the direction
				topAntennaRowDir, topAntennaColDir := getDirectionForPosition(topPosition, bottomPosition)
				bottomAntennaRowDir, bottomAntennaColDir := getDirectionForPosition(bottomPosition, topPosition)

				topIsValid, bottomIsValid := true, true
				topPositionCopy := topPosition
				bottomPositionCopy := bottomPosition

				// 4, 2, 5,5 & 8,9
				// take that distance and grab the top and bottom indexes in the grid that the antenna would go all the way until we hit the end
				for topIsValid {
					topAntenna := Position{Row: topPositionCopy.Row + (topAntennaRowDir * manhattanRowDistance), Column: topPositionCopy.Column + (topAntennaColDir * manhattanColDistance)}
					if utils.IsValid2DIndex(grid, topAntenna.Row, topAntenna.Column) {
						// add it to a map of indices we've found
						positionMap[topAntenna] = true
						totalCountForLetter++
					} else {
						topIsValid = false
					}
					topPositionCopy = topAntenna
				}

				for bottomIsValid {
					bottomAntenna := Position{Row: bottomPositionCopy.Row + (bottomAntennaRowDir * manhattanRowDistance), Column: bottomPositionCopy.Column + (bottomAntennaColDir * manhattanColDistance)}
					if utils.IsValid2DIndex(grid, bottomAntenna.Row, bottomAntenna.Column) {
						// add it to a map of indices we've found
						positionMap[bottomAntenna] = true
						totalCountForLetter++
					} else {
						bottomIsValid = false
					}
					bottomPositionCopy = bottomAntenna
				}
			}
		}

		if totalCountForLetter > 2 {
			// add all original antenna into position map also
			for _, antenna := range antennaMap {
				for _, posToAdd := range antenna {
					positionMap[posToAdd] = true
				}
			}
		}
	}

	return len(positionMap)
}
