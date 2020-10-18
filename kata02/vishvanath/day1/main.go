package main

func chop(searchTarget int, listOfIntegers []int) int {
	if len(listOfIntegers) == 0 {
		return -1
	}
	var startingIndex = 0
	var endIndex = len(listOfIntegers) - 1
	for startingIndex <= endIndex {
		var middleIndex = (startingIndex + endIndex) / 2
		var value = listOfIntegers[middleIndex]
		if value == searchTarget {
			return middleIndex
		} else if value < searchTarget {
			startingIndex = middleIndex + 1
		} else {
			endIndex = middleIndex - 1
		}
	}
	return -1
}
