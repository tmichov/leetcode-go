package main

import (
	"fmt"

	"github.com/tmichov/leetcode/easy"
)


func main() {
	prices := []int{7,6,4,3,1}
	test := easy.BestTimeToBuyAndSellStock(prices)

	fmt.Println(test)
}

