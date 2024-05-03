package easy

//1				 1
//2       1  1             
//3      1  2  1 
//4     1  3  3  1        
//5    1  4  6  4  1
//6   1  5 10 10 5  1
//7  1  6 15 20 15 6 1

func PascalsTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := range triangle {
		triangle[i] = make([]int, i+1)	
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

