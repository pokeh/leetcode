package main

import "fmt"

func main() {
	// fmt.Println(lcs("abc", "aba"))
	fmt.Println(lcs("hello", "olleh"))
}

func lcs(s1, s2 string) string {
	dp := make([][]int, len(s1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return decodeDpToString(s1, s2, dp)
}

func decodeDpToString(s1, s2 string, dp [][]int) string {
	res := ""

loop:
	for i, j := len(s1), len(s2); i > 0 && j > 0; {
		switch dp[i][j] {
		case dp[i-1][j]:
			i--
		case dp[i][j-1]:
			j--
		case dp[i-1][j-1] + 1:
			res += string(s1[i-1])
			i--
			j--
		case 0:
			break loop
		}
	}

	return reverse(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(s string) string {
	res := ""

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}
