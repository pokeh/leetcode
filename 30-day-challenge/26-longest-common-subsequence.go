func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, 2)
	numRow := len(text2) + 1
	dp[0], dp[1] = make([]int, numRow), make([]int, numRow)
​
	for _, r1 := range text1 {
		for j, r2 := range text2 {
			row := j + 1
​
			if r1 == r2 {
				dp[1][row] = max(dp[1][row-1], dp[0][row-1]+1)
			} else {
				dp[1][row] = max(dp[1][row-1], dp[0][row])
			}
		}
		copy(dp[0], dp[1])
		dp[1][0] = 0
	}
​
	return dp[0][numRow-1]
}
​
func max(a, b int) int {
	if a > b {
		return a
	}
​
	return b
}
