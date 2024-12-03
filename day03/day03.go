package day03

import (
	"adventofcode2024/utils"
	"regexp"
	"strings"
)

func PartA(filename string) int64 {
	input := utils.ReadFileAsSingleLine(filename)
	pattern := `mul\([0-9]{1,3}\,[0-9]{1,3}\)`
	re := regexp.MustCompile(pattern)
	// Find all matches and store them in an array
	matches := re.FindAllString(input, -1)
	var total int64 = 0

	// now that we have matches; multiply
	for _, match := range matches {
		// split on comma
		numStrings := strings.Split(match, ",")
		num01String := numStrings[0]
		num02String := numStrings[1]
		// strip mul(
		num01String = strings.ReplaceAll(num01String, "mul(", "")
		// strip )
		num02String = strings.ReplaceAll(num02String, ")", "")
		// multiply
		num01 := utils.ConvertStringToInt64(num01String)
		num02 := utils.ConvertStringToInt64(num02String)
		var totalToAdd int64 = num01 * num02
		total += totalToAdd
	}
	return total
}
