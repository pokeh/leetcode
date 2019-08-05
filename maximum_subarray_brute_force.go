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
