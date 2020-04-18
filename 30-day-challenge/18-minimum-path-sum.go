func minPathSum(grid [][]int) int {
	numRow := len(grid)
	numCol := len(grid[0])

	sums := make([][]int, numRow)
	for i := range sums {
		sums[i] = make([]int, numCol)
	}

	for i := 0; i < numRow; i++ {
		for j := 0; j < numCol; j++ {
			val := grid[i][j]

			if i == 0 && j == 0 {
				sums[i][j] = val
				continue
			}
			if i == 0 {
				sums[i][j] = sums[i][j-1] + val
				continue
			}
			if j == 0 {
				sums[i][j] = sums[i-1][j] + val
				continue
			}
			sums[i][j] = min(sums[i-1][j], sums[i][j-1]) + val
		}
	}

	return sums[numRow-1][numCol-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
