package main

// Kadane's algorithm -----------------------------
// runtime complexity = O(n)
// memory = O(1)
func maxSubArray(nums []int) int {
	// local max
	lMax := nums[0]
	// global max
	gMax := nums[0]

	for i := 1; i < len(nums); i++ {
		lMax = max(lMax+nums[i], nums[i])
		gMax = max(lMax, gMax)
	}

	return gMax
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// Kadane's algorithm -----------------------------

/*
// brute force ------------------------------------
// runtime complexity = O(n^2)
// memory = O(1)
func maxSubArray(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    prevMax := maxSubArray(nums[:len(nums)-1])
    currMax := maxContainLastEl(nums)

    if currMax > prevMax {
        return currMax
    } else {
        return prevMax
    }
}

func maxContainLastEl(nums []int) int {
    sum := nums[len(nums)-1]
    max := sum

    for i := len(nums)-2; i >= 0; i-- {
        newSum := sum + nums[i]
        if max < newSum {
            max = newSum
        }
        sum = newSum
    }

    return max
}
// brute force ------------------------------------
*/
