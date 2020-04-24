func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		// fmt.Printf("nums[%v]: %v, nums[%v]: %v\n", left, nums[left], right, nums[right])
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// in left half (no pivot)
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
				continue
			}

			// in right half (with pivot)
			left = mid + 1
			continue
		}

		// in right half (no pivot)
		if nums[mid] < target && target <= nums[right] {
			left = mid + 1
			continue
		}

		// in left half (with pivot)
		right = mid - 1
	}

	return -1
}
