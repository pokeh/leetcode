func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	n, m := dim[0], dim[1]

	row := 0
	col := m - 1

	for row < n {
		if binaryMatrix.Get(row, col) == 0 {
			// move down
			row++
		} else {
			// move left
			col--

			// reached the first column
			if col < 0 {
				return 0
			}
		}
	}

	// not found
	if col == m-1 {
		return -1
	}

	return col + 1
}
