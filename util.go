package main

func uniq(nums []int) []int {
	m := make(map[int]bool)

	res := make([]int, 0, len(nums))

	for _, n := range nums {
		if !m[n] {
			res = append(res, n)
			m[n] = true
		}
	}

	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -1 * a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
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
