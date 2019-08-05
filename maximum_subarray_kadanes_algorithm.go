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
