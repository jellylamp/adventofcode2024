package day09

import (
	"adventofcode2024/utils"
	"strconv"
)

func PartA(filename string) int {
	line := utils.ReadFileAsSingleLine(filename)

	idVal := 0
	uncompactedList := []string{}
	periodLocationList := []int{}

	for index, characterAsRune := range line {
		character := string(characterAsRune)
		// loop through character by character
		isOdd := index%2 != 0

		if isOdd {
			// blank space X times
			xTimes := utils.ConvertStringToInt(character)
			for i := 0; i < xTimes; i++ {
				uncompactedList = append(uncompactedList, ".")
				// keep track of where periods are so we can more easily traverse the list again
				periodLocationList = append(periodLocationList, len(uncompactedList)-1)
			}
		} else {
			// add id in character num times
			xTimes := utils.ConvertStringToInt(character)
			for i := 0; i < xTimes; i++ {
				uncompactedList = append(uncompactedList, strconv.Itoa(idVal))
			}

			// increment after
			idVal++
		}
	}

	// make a copy so we have uncompacted list for debugging
	compactedList := make([]string, len(uncompactedList))
	copy(compactedList, uncompactedList)
	compactedListLength := len(compactedList)
	stopLooping := false

	// loop through period list so we can compact and replace
	for index := 0; index < len(periodLocationList); index++ {
		indexToReplace := periodLocationList[index]

		// loop through the end to find the last digit in the list
		for endIndex := compactedListLength - 1; endIndex >= 0; endIndex-- {
			var locationOfLastDigit int
			if compactedList[endIndex] != "." {
				// wooo a digit
				locationOfLastDigit = endIndex
			} else {
				// keep searchin
				continue
			}

			if locationOfLastDigit > periodLocationList[index] {
				// first index of a period, replace it with the last number of the array
				compactedList[indexToReplace] = compactedList[locationOfLastDigit]
				compactedList[locationOfLastDigit] = "."
				break
			} else {
				// we have fully compressed
				stopLooping = true
				break
			}
		}
		if stopLooping {
			break
		}
	}

	checksum := 0
	// loop through compacted list and make a checksum!
	for index, value := range compactedList {
		if value != "." {
			bitChecksum := index * utils.ConvertStringToInt(value)
			checksum += bitChecksum
		}
	}
	return checksum
}

type Fragment struct {
	length                 int
	id                     string
	startingLocation       int
	hasBeenCheckedOrFilled bool
}

func PartB(filename string) int {
	line := utils.ReadFileAsSingleLine(filename)

	idVal := 0
	uncompactedList := []Fragment{}
	uncompactedLength := 0
	periodLocationList := []Fragment{}
	idLocationList := []Fragment{}

	listLength := 0
	for index, characterAsRune := range line {
		character := string(characterAsRune)
		// loop through character by character
		isOdd := index%2 != 0

		if isOdd {
			// blank space X times
			xTimes := utils.ConvertStringToInt(character)
			if xTimes != 0 {
				uncompactedList = append(uncompactedList, Fragment{xTimes, ".", listLength, false})

				// keep track of where periods are so we can more easily traverse the list again
				periodLocationList = append(periodLocationList, Fragment{xTimes, ".", listLength, false})
				listLength += xTimes
				uncompactedLength++
			}
		} else {
			// add id in character num times
			xTimes := utils.ConvertStringToInt(character)
			if xTimes != 0 {
				uncompactedList = append(uncompactedList, Fragment{xTimes, strconv.Itoa(idVal), listLength, false})
				idLocationList = append(idLocationList, Fragment{xTimes, strconv.Itoa(idVal), listLength, false})

				// increment after
				listLength += xTimes
				uncompactedLength++
				idVal++
			}
		}
	}

	// make a copy so we have uncompacted list for debugging
	compactedList := make([]Fragment, len(uncompactedList))
	copy(compactedList, uncompactedList)
	periodFragmentListCopy := make([]Fragment, len(periodLocationList))
	copy(periodFragmentListCopy, periodLocationList)

	idListLength := len(idLocationList)

	for idListIndex := idListLength - 1; idListIndex >= 0; idListIndex-- {
		// starting with the end of the list (highest id)
		//see if there is a section in the period list that will contain it!
		currentFragment := idLocationList[idListIndex]
		currentFragmentIndex := findFragmentIndex(compactedList, currentFragment)

		for periodIndex := range periodLocationList {
			// update the copy, not what we are looping over
			// oh shit we using pointers baby
			periodFragment := &periodFragmentListCopy[periodIndex]
			if periodFragment.length >= currentFragment.length && !periodFragment.hasBeenCheckedOrFilled {
				periodFragmentIndex := findFragmentIndex(compactedList, *periodFragment)

				// only move things if they move back
				if periodFragmentIndex < currentFragmentIndex {
					// replace current fragments location with periods
					compactedList[currentFragmentIndex] = Fragment{currentFragment.length, ".", currentFragment.startingLocation, true}

					// replace period fragment with current fragment
					compactedList[periodFragmentIndex] = currentFragment

					// If there is any space left.... update period fragments index and length
					lengthLeftover := periodFragment.length - currentFragment.length
					if lengthLeftover > 0 {
						periodFragment.length = lengthLeftover
						periodFragment.startingLocation += lengthLeftover

						// ugh fuck this again

						// need to make sure compacted list still has this spot saved...
						compactedList = append(compactedList[:periodFragmentIndex+1], append([]Fragment{*periodFragment}, compactedList[periodFragmentIndex+1:]...)...)
					} else {
						// remove period fragment from list
						periodFragment.hasBeenCheckedOrFilled = true
					}
					break
				} else {
					// we can break here... if at this point its smaller it wont get better
					break
				}
			}
		}
	}

	flattenedFragments := flattenFragments(compactedList)
	checksum := 0
	// loop through compacted list and make a checksum!
	for index, fragment := range flattenedFragments {
		if fragment.id != "." {
			bitChecksum := index * utils.ConvertStringToInt(fragment.id)
			checksum += bitChecksum
		}
	}
	return checksum
}

func areFragmentsEqual(a, b Fragment) bool {
	return a.length == b.length &&
		a.id == b.id &&
		a.startingLocation == b.startingLocation &&
		a.hasBeenCheckedOrFilled == b.hasBeenCheckedOrFilled
}

func findFragmentIndex(list []Fragment, target Fragment) int {
	for i, fragment := range list {
		if areFragmentsEqual(fragment, target) {
			return i
		}
	}
	return -1
}

func flattenFragments(fragments []Fragment) []Fragment {
	var flattened []Fragment
	for _, fragment := range fragments {
		for i := 0; i < fragment.length; i++ {
			flattened = append(flattened, Fragment{
				length:                 1,
				id:                     fragment.id,
				startingLocation:       fragment.startingLocation + i,
				hasBeenCheckedOrFilled: fragment.hasBeenCheckedOrFilled,
			})
		}
	}
	return flattened
}
