// thanks to https://www.youtube.com/watch?v=ojz8xZc8pog
func countSquares(matrix [][]int) int {
	// initialize dp with zeros
	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	var res int

	// fill dp with number of squares with the lower-right edge on that point
	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == 1 {
				dp[i+1][j+1] = min(min(dp[i][j], dp[i+1][j]), dp[i][j+1]) + 1

				// also sum these counts up as we go
				res += dp[i+1][j+1]
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
