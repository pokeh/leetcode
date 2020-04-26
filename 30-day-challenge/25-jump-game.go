// Greedy
// Time complexity: O(), Space complexity: O(n)
func canJump(nums []int) bool {
	pos := len(nums) - 1
	for i := len(nums)-1; i >= 0; i-- {
		if i + nums[i] >= pos {
			pos = i
		}
	}
	return pos == 0
}

/*
// DP
// Time complexity: O(n^2), Space complexity: O(n)
func canJump(nums []int) bool {
	// holds whether the end is reachable from that index
	canReachEnd := make([]bool, len(nums))
	canReachEnd[len(nums)-1] = true

Loop:
	for i := len(nums) - 2; i >= 0; i-- {
		n := nums[i]

		if i+n > len(nums) {
			canReachEnd[i] = true
			continue Loop
		}

		for j := 1; j <= n; j++ {
			if canReachEnd[i+j] {
				canReachEnd[i] = true
				continue Loop
			}
		}
	}

	return canReachEnd[0]
}
*/
