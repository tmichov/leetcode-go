package medium

import (
	"slices"
)

func BoatsToSavePeople(people []int, limit int) int {
	slices.Sort(people)

	numBoats := 0

	start := 0
	end := len(people) - 1

	for start < end {
		if people[end]+people[start] <= limit {
			start++
		}

		numBoats++
		end--

		if start == end {
			numBoats++
		}
	}

	return numBoats
}
