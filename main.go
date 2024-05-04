package main

import (
	"fmt"

	"github.com/tmichov/leetcode/medium"
)


func main() {
	people := []int{2,2,2,2,2,2}
	balanced := medium.BoatsToSavePeople(people, 16)

	fmt.Println("balanced: ", balanced)
}

