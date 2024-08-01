package medium

import (
	"fmt"
	"strconv"
)

func ValidSudoku(board [][]byte) bool {
	columnMap := make([][]bool, 10)
	groupMap := make([][]bool, 10)

	for i:=0; i<9; i++ {
		columnMap[i] = make([]bool, 10)
		groupMap[i] = make([]bool, 10)
	}

	for row := 0; row < 9; row++ {
		rowMap := make([]bool, 10)

		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue;
			}

			value, _ := strconv.Atoi(fmt.Sprintf("%c", board[row][col]))

			if rowMap[value] {
				return false
			} else {
				rowMap[value] = true
			}

			if columnMap[col][value] {
				return false;
			} else {
				columnMap[col][value] = true
			}

			groupId := getGroupId(row, col)
			if groupMap[groupId][value] {
				return false;
			} else {
				groupMap[groupId][value] = true
			}
		}
	}

	return true;
}

func getGroupId(row, col int) int {
	return (row/3) + ((col/3) * 3)
}
