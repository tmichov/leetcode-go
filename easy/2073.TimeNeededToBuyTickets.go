package easy

func TimeNeededToBuyTickets(tickets []int, k int) int {
	target := tickets[k]
	seconds := 0

	for i, v := range tickets {
		if i <= k {
			seconds += min(v, target)
		} else {
			seconds += min(v, target-1)
		}
	}

	return seconds
}
