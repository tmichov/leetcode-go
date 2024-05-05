package easy

func BestTimeToBuyAndSellStock(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		}

		if v - minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}

	return maxProfit
}
