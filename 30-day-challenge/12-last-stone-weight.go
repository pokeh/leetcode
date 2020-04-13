func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Slice(stones, func(a, b int) bool { return stones[a] > stones[b] })

		diff := stones[0]-stones[1]
		stones = stones[2:]

		if diff != 0 {
			stones = append(stones, diff)
		}
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
