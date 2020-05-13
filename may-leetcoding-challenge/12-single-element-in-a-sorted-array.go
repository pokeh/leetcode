func singleNonDuplicate(nums []int) int {
	return find(nums, 0, len(nums)-1)
}

func find(nums []int, lo, hi int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2

	if mid > 0 && nums[mid] == nums[mid-1] {
		if res := find(nums, lo, mid-2); res != -1 {
			return res
		}
		if res := find(nums, mid+1, hi); res != -1 {
			return res
		}
		return -1
	}

	if mid < len(nums)-1 && nums[mid] == nums[mid+1] {
		if res := find(nums, lo, mid-1); res != -1 {
			return res
		}
		if res := find(nums, mid+2, hi); res != -1 {
			return res
		}
		return -1
	}

	return nums[mid]
}
