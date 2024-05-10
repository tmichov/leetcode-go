package medium

import "sort"

type Fraction struct {
	Numerator   int
	Denominator int
	Value       float64
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	length := len(arr)

	fractions := make([]Fraction, 0)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			fractions = append(fractions, Fraction{
				Numerator:   arr[i],
				Denominator: arr[j],
				Value:       float64(arr[i]) / float64(arr[j]),
			})
		}
	}

	sort.Slice(fractions, func(i, j int) bool {
		return fractions[i].Value < fractions[j].Value
	})

	return []int{fractions[k-1].Numerator, fractions[k-1].Denominator}
}
