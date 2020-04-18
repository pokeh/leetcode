func numIslands(grid [][]byte) int {
	var cnt int

	for i, row := range grid {
		for j, val := range row {
			if val == '1' {
				cnt++
				removeSameIsland(grid, i, j)
			}
		}
	}

	return cnt
}

func removeSameIsland(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
		return
	}

	if grid[row][col] != '1' {
		return
	}

	grid[row][col] = '0'

	removeSameIsland(grid, row-1, col)
	removeSameIsland(grid, row+1, col)
	removeSameIsland(grid, row, col-1)
	removeSameIsland(grid, row, col+1)
}
