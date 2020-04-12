func maxSubArray(nums []int) int {
	localMax := nums[0]
	globalMax := nums[0]

	for i := 1; i < len(nums); i++ {
		localMax = max(localMax+nums[i], nums[i])
		globalMax = max(localMax, globalMax)
	}

	return globalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
