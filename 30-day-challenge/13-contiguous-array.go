func findMaxLength(nums []int) int {
	var sum, max int

	sums := make(map[int]int)
	sums[0] = -1

	for i, n := range nums {
		if n == 0 {
			sum--
		} else {
			sum++
		}

		if idx, ok := sums[sum]; ok {
			if (i-idx) > max {
				max = i-idx
			}
		} else {
			sums[sum] = i
		}
	}

	return max
}
