package day14

import (
	"adventofcode2024/utils"
	"regexp"
)

type Coordinate struct {
	X int
	Y int
}

func PartA(filename string, gridWidth int, gridHeight int, numSeconds int) int {
	lineArr := utils.ReadFileToArrayOfLines(filename)

	// p=0,4 v=3,-3
	// p=6,3 v=-1,-3
	allCoordinates := make([]Coordinate, 0)
	midCoordinates := make([]Coordinate, 0)

	// Regular expression to match numbers (including negatives)
	re := regexp.MustCompile(`-?\d+`)

	midXLine := gridWidth / 2
	midYLine := gridHeight / 2

	for _, line := range lineArr {
		lineNums := re.FindAllString(line, -1)
		originalX := utils.ConvertStringToInt(lineNums[0])
		originalY := utils.ConvertStringToInt(lineNums[1])
		velocityX := utils.ConvertStringToInt(lineNums[2])
		velocityY := utils.ConvertStringToInt(lineNums[3])

		// do something with mod and grid size and num seconds
		// xt​=(x0​+vx​⋅t)mod w
		// yt​=(y0​+vy​⋅t)mod h
		finalX := (originalX + (velocityX * numSeconds)) % gridWidth
		finalY := (originalY + (velocityY * numSeconds)) % gridHeight

		// Handle negative modulo results by adding w (for x) or h (for y). (not the same as absolute value)
		if finalX < 0 {
			finalX += gridWidth
		}
		if finalY < 0 {
			finalY += gridHeight
		}
		// add final location to a grid list (not a map they can share a square)
		// DO NOT add if it is on the exact middle; we will add to another list for debugging
		if finalX == midXLine {
			midCoordinates = append(midCoordinates, Coordinate{finalX, finalY})
		} else if finalY == midYLine {
			midCoordinates = append(midCoordinates, Coordinate{finalX, finalY})
		} else {
			allCoordinates = append(allCoordinates, Coordinate{finalX, finalY})
		}
	}

	NWQuadrant := 0
	NEQuadrant := 0
	SWQuadrant := 0
	SEQuadrant := 0

	// loop over all final locations and calculate quadrants
	for _, coordinate := range allCoordinates {
		if coordinate.X < midXLine && coordinate.Y < midYLine {
			NWQuadrant++
		} else if coordinate.X > midXLine && coordinate.Y < midYLine {
			NEQuadrant++
		} else if coordinate.X < midXLine && coordinate.Y > midYLine {
			SWQuadrant++
		} else if coordinate.X > midXLine && coordinate.Y > midYLine {
			SEQuadrant++
		}
	}

	return NWQuadrant * NEQuadrant * SWQuadrant * SEQuadrant
}
