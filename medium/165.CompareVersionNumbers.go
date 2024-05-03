package medium

import (
	"fmt"
	"strconv"
	"strings"
)

func CompareVersionNumbers(version1, version2 string) int {
	ints1 := parseVersionNumber(version1)
	ints2 := parseVersionNumber(version2)

	fmt.Println(ints1, ints2)

	len1 := len(ints1)
	len2 := len(ints2)

	for i:=0; i < max(len1, len2); i++ {
		v1 := 0
		v2 := 0
		if i < len1 {
			v1 = ints1[i]
		}
		if i < len2 {
			v2 = ints2[i]
		}

		if v1 > v2 {
			return 1
		}

		if v2 > v1 {
			return -1
		}
	}

	return 0
}

func parseVersionNumber(version string) []int {
	segments := strings.Split(version, ".")

	ints := []int{}
	for _, v := range segments {
		for v[0] == '0' && len(v) > 1 {
			v = v[1:];
		}

		val, err := strconv.Atoi(v)

		if err != nil {
			ints = append(ints, 0)
		}

		ints = append(ints, val)
	}

	return ints
}
