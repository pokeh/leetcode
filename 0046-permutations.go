func permute(nums []int) [][]int {
	res := make([][]int, 0)

	var p func(i int)
	p = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}

		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]
			p(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	p(0)

	return res
}
