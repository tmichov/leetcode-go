package easy

// 1700. Number of students unable to eat lunch, daily
func NumberOfStudentsUnableToEatLunch(students []int, sandwiches []int) int {
	countPreferences := make(map[int]int)

	for _, preference := range students {
		countPreferences[preference]++
	}

	for _, sandwich := range sandwiches {
		if countPreferences[sandwich] == 0 {
			break
		}
		countPreferences[sandwich]--
	}

	return countPreferences[0] + countPreferences[1]
}
