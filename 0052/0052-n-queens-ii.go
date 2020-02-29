package main

import "fmt"

func totalNQueens(n int) int {
	count := 0

	// columns[{row number}] = {column number}
	columns := make([]int, n)

	solveHelper(0, n, &columns, &count)

	return count
}

func solveHelper(row, n int, cols *[]int, cnt *int) {
	if row == n {
		*cnt++
		return
	}

	for col := 0; col < n; col++ {
		if canPlace(row, col, n, *cols) {
			(*cols)[row] = col
			solveHelper(row+1, n, cols, cnt)
		}
	}
}

func canPlace(row, col, n int, cols []int) bool {
	for i := 0; i < row; i++ {
		if col == cols[i] || (col-row == cols[i]-i) || (col+row == cols[i]+i) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(totalNQueens(4))
}
