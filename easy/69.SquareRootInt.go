package easy

func SquareRootInt(x int) int {
	z := 1.0
	i := 1
	for {
		oldZ := z
		z -= (z*z - float64(x)) / (2 * z)

		if oldZ-z < 0.00000001 && oldZ-z > 0 {
			break
		}

		if oldZ == z && i != 1 {
			break
		}

		i++
	}

	return int(z)
}
