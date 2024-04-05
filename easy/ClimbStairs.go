package easy

func ClimbStairs(n int) int {
		secondLast, last := 0, 1

		for i:=1; i<=n; i++ {
				secondLast, last = last, secondLast+last
		}

		return last
}
