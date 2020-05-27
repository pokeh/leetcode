// Idea: Keep track of "tallies" which increments when n == 1 and decrements when n == 0
// eg. nums := [0,0,1,1,0,1,1]
// Then tallies := [-1,-2,-1,0,-1,0,1]
//
// Then contiguous subarrays are one of two cases:
//   * from the beginning of the array to where the tally is 0, or
//   * between two values where the tally is the same
//
// In the example, contiguous subarrays are nums[:4], nums[:6], nums[1:3], nums[1:5], and nums[3:5]
//
func findMaxLength(nums []int) int {
	// set {0: -1} to simplify computation for the tally == 0 case
	firstIdxForTally := map[int]int{0: -1}

	var maxLen, tally int

	for i, n := range nums {
		if n == 0 {
			tally--
		} else {
			tally++
		}

		if firstIdx, ok := firstIdxForTally[tally]; !ok {
			firstIdxForTally[tally] = i
		} else {
			// nums[firstIdx+1:i] is a contiguous subarray
			currLen := i - firstIdx
			if currLen > maxLen {
				maxLen = currLen
			}
		}
	}

	return maxLen
}

/*
// Time limit exceeded
func findMaxLength(nums []int) int {
	// stores the tally for each index
	tallies := make([]int, len(nums))

	var tally int

	for i, n := range nums {
		if n == 0 {
			tally--
		} else {
			tally++
		}
		tallies[i] = tally
	}

	// fmt.Println("tallies:", tallies)

	var maxLen int

	for i, iTally := range tallies {
		// if iTally == 0, nums[0:i] is a contiguous subarray
		if iTally == 0 && i > maxLen {
			maxLen = i + 1
		}

		for j, jTally := range tallies {
			// if iTally == jTally, nums[i+1:j] is a contiguous subarray
			if jTally == iTally && j-i > maxLen {
				maxLen = j - i
			}
		}
	}

	return maxLen
}
*/
