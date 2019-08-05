// application of Kadane's algorithm
func maxProfit(prices []int) int {
    // local max
    lMax := 0
    // global max
    gMax := 0
    // a temp variable to store the diff between two consecutive prices
    diff := 0
    
    for i := 1; i < len(prices); i++ {
        diff = prices[i] - prices[i-1]
        lMax = max(lMax+diff, diff)
        gMax = max(gMax, lMax)
    }
    
    return max(gMax, 0)
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}
