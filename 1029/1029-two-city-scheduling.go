func twoCitySchedCost(costs [][]int) int {
	var totalCost int
	bGain := make([]int, len(costs))

	for i, c := range costs {
		totalCost += c[0] // choose all A at first
		bGain[i] = c[1] - c[0]
	}

	sort.Ints(bGain)

	for i := 0; i < len(costs)/2; i++ {
		// change half to the most economizing Bs
		totalCost += bGain[i]
	}

	return totalCost
}
