package easy

import (
	"sort"
	"strconv"
)

func RelativeRanks(score []int) []string {
	var sorted []int
	sorted = append(sorted, score...)
	sort.Ints(sorted)

	pos := len(sorted)
	values := make(map[int]int)

	for _, v := range sorted {
		values[v] = pos
		pos--
	}

	var result []string

	for _, s := range score {
		place := values[s]
		if place == 1 {
			result = append(result, "Gold Medal")
		} else if place == 2 {
			result = append(result, "Silver Medal")
		} else if place == 3 {
			result = append(result, "Bronze Medal")
		} else {
			result = append(result, strconv.Itoa(place))
		}
	}

	return result
}
