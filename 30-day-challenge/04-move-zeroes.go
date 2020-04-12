func moveZeroes(nums []int)  {
	ptr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[ptr] = nums[i]
			ptr++
		}
	}
	for i := ptr; i < len(nums); i++ {
		nums[i] = 0
	}
}
