func maxProfit(prices []int) int {
	var sum int

	for i := 1; i < len(prices); i++ {
		diff := prices[i]-prices[i-1]

		if diff > 0 {
			sum += diff
		}
	}

	return sum
}
