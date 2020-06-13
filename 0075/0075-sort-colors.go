// thanks to http://users.monash.edu/~lloyd/tildeAlgDS/Sort/Flag/
func sortColors(nums []int) {
	lo, mid, hi := 0, 0, len(nums)-1

	for mid <= hi {
		switch nums[mid] {
		case 0:
			nums[lo], nums[mid] = nums[mid], nums[lo]
			lo++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[hi] = nums[hi], nums[mid]
			hi--
		}
	}
}

// 0 and 1 only
func sortTwoColors(nums []int) {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		switch nums[lo] {
		case 0:
			lo++
		case 1:
			nums[lo], nums[hi] = nums[hi], nums[lo]
			hi--
		}
	}
}
