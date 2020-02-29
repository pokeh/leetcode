package main

func rob(nums []int) int {
	memo := [2]int{0, 0}

	for _, n := range nums {
		memo[0], memo[1] = memo[1], max(n+memo[0], memo[1])
	}

	return memo[1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
// Time limit exceeded
func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
		return max(nums[0]+rob(nums[2:]), rob(nums[1:]))
	}
}
*/
