func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dfs(&image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(imgPtr *[][]int, row, col, oldColor, newColor int) {
	if oldColor == newColor {
		return
	}

	img := *imgPtr

	if img[row][col] != oldColor {
		return
	}

	img[row][col] = newColor

	if row > 0 {
		dfs(imgPtr, row-1, col, oldColor, newColor)
	}
	if col > 0 {
		dfs(imgPtr, row, col-1, oldColor, newColor)
	}
	if row < len(img)-1 {
		dfs(imgPtr, row+1, col, oldColor, newColor)
	}
	if col < len(img[0])-1 {
		dfs(imgPtr, row, col+1, oldColor, newColor)
	}
}
