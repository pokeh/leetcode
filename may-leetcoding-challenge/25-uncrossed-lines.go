func maxUncrossedLines(A []int, B []int) int {
	// dp is a memo table to store the max num up to that point
	// add a zero-padded top row and left column to simplify computations
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}

	// fill dp
	for i, a := range A {
		for j, b := range B {
			if a == b {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])

			}
		}
	}

	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return dp[len(A)][len(B)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
