func productExceptSelf(nums []int) []int {
	var res []int

	left := []int{1}
	right := []int{1}

	for i := 1; i < len(nums); i++ {
		left = append(left, left[i-1] * nums[i-1])
		right = append(right, right[i-1] * nums[len(nums)-i])
	}

	for i := 0; i < len(nums); i++ {
		res = append(res, left[i] * right[len(nums)-i-1])
	}

	return res
}

/*
// not O(n) :(
func productExceptSelf(nums []int) []int {
	var res []int

	for i, _ := range nums {
		prod := 1
		for j, n := range nums {
			if i != j {
				prod *= n
			}
		}
		res = append(res, prod)
	}

	return res
}
*/

/*
// uses division :(
func productExceptSelf(nums []int) []int {
	prod := 1
	zeroCnt := 0

	for _, n := range nums {
		if n == 0 {
			zeroCnt++
		} else {
			prod *= n
		}
	}

	if zeroCnt > 1 {
		return make([]int, len(nums))
	}

	var res []int
	for _, n := range nums {
		var val int
		if n == 0 {
			val = prod
		} else if zeroCnt > 0 {
			val = 0
		} else {
			val = prod/n
		}

		res = append(res, val)
	}
	return res
}
*/
