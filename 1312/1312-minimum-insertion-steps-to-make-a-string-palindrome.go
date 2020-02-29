package main

func minInsertions(s string) int {
	if len(s) < 2 {
		return 0
	}

	if reverse(s) == s {
		return 0
	}

	return len(s) - lcs(s, reverse(s))
}

func reverse(s string) string {
	var r string

	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}

	return r
}

// longest common subsequence
// solved by bottom-up DP
func lcs(s1, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)

	dp := make([][]int, l1+1)

	for i, _ := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[l1][l2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
