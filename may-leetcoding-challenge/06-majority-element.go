func majorityElement(nums []int) int {
	maj := len(nums) / 2
	cnts := make(map[int]int)

	for _, n := range nums {
		cnts[n]++
		if cnts[n] > maj {
			return n
		}
	}

	// not idiomatic Go, but since we can't return an error
	panic("majority element doesn't exist")
}
