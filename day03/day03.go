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

// this is probably doable in regex; but go doesn't support negative lookaheads/behinds and it seemed hard.
func PartB(filename string) int64 {
	input := utils.ReadFileAsSingleLine(filename)
	pattern := `mul\([0-9]{1,3}\,[0-9]{1,3}\)`
	re := regexp.MustCompile(pattern)
	doInputs := strings.Split(input, "do()")
	var allMatches []string

	for _, doStatement := range doInputs {
		// we have already split on do() so we can split on don't() and ignore everything but the first object.
		dontStatements := strings.Split(doStatement, "don't()")
		doWithoutDont := dontStatements[0]

		// Find all matches and store them in an array
		matches := re.FindAllString(doWithoutDont, -1)
		allMatches = append(allMatches, matches...)
	}

	var total int64 = 0

	// now that we have matches; multiply
	for _, match := range allMatches {
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
