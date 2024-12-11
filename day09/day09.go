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
	length           int
	id               string
	startingLocation int
	indexInList      int
}

func PartB(filename string) int {
	line := utils.ReadFileAsSingleLine(filename)

	idVal := 0
	uncompactedList := []Fragment{}
	uncompactedLength := 0
	periodLocationList := []Fragment{}

	listLength := 0
	for index, characterAsRune := range line {
		character := string(characterAsRune)
		// loop through character by character
		isOdd := index%2 != 0

		if isOdd {
			// blank space X times
			xTimes := utils.ConvertStringToInt(character)

			// maybe we only add periods in x times but not the normal ones, so we dont get messed up?
			for i := 0; i < xTimes; i++ {
				uncompactedList = append(uncompactedList, Fragment{xTimes, ".", listLength, uncompactedLength})
				// keep track of where periods are so we can more easily traverse the list again
				periodLocationList = append(periodLocationList, Fragment{xTimes, ".", listLength, uncompactedLength})
			}
			listLength += xTimes
			uncompactedLength++
		} else {
			// add id in character num times
			xTimes := utils.ConvertStringToInt(character)
			uncompactedList = append(uncompactedList, Fragment{xTimes, strconv.Itoa(idVal), listLength, uncompactedLength})

			// increment after
			listLength += xTimes
			uncompactedLength++
			idVal++
		}
	}

	// make a copy so we have uncompacted list for debugging
	compactedList := make([]Fragment, len(uncompactedList))
	copy(compactedList, uncompactedList)

	// make a copy of this list so our index's stay relevant after insertion
	compactedListUnaltered := make([]Fragment, len(uncompactedList))
	copy(compactedListUnaltered, uncompactedList)

	compactedListLength := len(compactedList)
	// lastFragment := compactedList[compactedListLength -1]
	// highestFragmentId := lastFragment.id
	var lastFragment Fragment
	highestId := -1

	// loop through period blocks so we can compact and replace
	// 00992111777.44.333....5555.6666.....8888..
	for index := 0; index < len(periodLocationList); index++ {
		indexToReplace := periodLocationList[index].indexInList
		blockLength := periodLocationList[index].length

		for endIndex := compactedListLength - 1; endIndex >= 0; endIndex-- {
			lastFragment = compactedList[endIndex]
			if lastFragment.id != "." {
				lastFragmentIdInInt := utils.ConvertStringToInt(lastFragment.id)
				if highestId > lastFragmentIdInInt || highestId == -1 {
					// we have not tried this fragments id!
					highestId = lastFragmentIdInInt
					if lastFragment.length <= blockLength {

						// fresh index proceed like normal
						if compactedList[indexToReplace].id == "." {
							// replace our block with a "." fragment our blocks long
							compactedList[lastFragment.indexInList] = Fragment{lastFragment.length, ".", lastFragment.startingLocation, lastFragment.indexInList}
							smallFragment := lastFragment
							smallFragment.length = 1
							for blocksFilled := lastFragment.length; blocksFilled > 0; blocksFilled-- {
								// replace indexToReplaceWithOurBlock - but size it down to 1 so we don't miscalculate later
								compactedList[indexToReplace] = smallFragment
								indexToReplace++
								blockLength--
							}

							if blockLength > 0 {
								// we gotta keep filling
								continue
							} else {
								// we have filled it completely, break
								break
							}
						}
					}
				} else {
					// we have tried this fragment, move along
					continue
				}
			} else {
				// we found a blank keep going
				continue
			}
		}
	}

	checksum := 0
	// loop through compacted list and make a checksum!
	for index, fragment := range compactedList {
		if fragment.id != "." {
			bitChecksum := index * utils.ConvertStringToInt(fragment.id)
			checksum += bitChecksum
		}
	}
	return checksum
}
