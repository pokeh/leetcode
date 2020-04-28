/*
[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
[["1","1","1","1","1","1","1","1"],["1","1","1","1","1","1","1","0"],["1","1","1","1","1","1","1","0"],["1","1","1","1","1","0","0","0"],["0","1","1","1","1","0","0","0"]]
[["0","0","0","1"],["1","1","0","1"],["1","1","1","1"],["0","1","1","1"],["0","1","1","1"]]
[]
[["0"]]
[["1"]]
[["0","1"]]
*/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	var hasOne bool

	// initialize dp with first row and column copied over from matrix
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		val := convertToInt(matrix[i][0])
		if val == 1 {
			hasOne = true
		}
		dp[i][0] = convertToInt(matrix[i][0])
	}
	for i := range matrix[0] {
		val := convertToInt(matrix[0][i])
		if val == 1 {
			hasOne = true
		}
		dp[0][i] = val
	}

	var maxLen int
	if hasOne {
		maxLen = 1
	}

	// do dp
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			var length int

			if matrix[i][j] == byte('0') {
				length = 0
			} else {
				length = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}

			// fmt.Printf("length at [%d][%d] is %d\n", i, j, length)
			dp[i][j] = length
			maxLen = max(maxLen, length)
		}
	}

	return maxLen * maxLen
}

func convertToInt(b byte) int {
	if b == byte('0') {
		return 0
	}
	return 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
