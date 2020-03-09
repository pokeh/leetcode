package main

func singleNumber(nums []int) int {
	res := 0

	for _, num := range nums {
		// since n^n=0 for all n, this will cancel out numbers that appear twice
		res ^= num
	}

	return res
}
