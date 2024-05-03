package main

import (
	"fmt"

	"github.com/tmichov/leetcode/medium"
)

func main() {
	balanced := medium.CompareVersionNumbers("1.1.1", "1.00001")

	fmt.Println("balanced: ", balanced)
}

