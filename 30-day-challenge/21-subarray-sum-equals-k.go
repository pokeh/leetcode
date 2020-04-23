func subarraySum(nums []int, k int) int {
	var cnt int

	for i := range nums {
		var sum int

		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				cnt++
			}
		}
	}

	return cnt
}
