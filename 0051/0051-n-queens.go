package main

import "fmt"

// Theme: backtracking
//
// * We know that exactly one queen has to appear on each row
//   => we'll choose to place the i-th queen on the i-th row
//
// * The new problem is to find out the "right" number of columns for each row
//   => we'll have a slice to hold column numbers of the i-th row (columns)
//
// * Tactic
//   loop: Try placing a first queen (row=0) at some column
//     loop: Try placing a second queen (row=1) at some column. Avoid columns that won't work
//       loop: Try placing a third queen (row=2) at some column. Avoid columns that won't work
//         etc...
//
// * What is "columns that won't work"?
//   * same row
//     * due to how we set up the algorithm, this won't occur
//   * same column
//     * columns[row] == columns[i] for some i
//   * same diagonal
//     *  columns[row]-row == columns[i]-i for some i
//     *  columns[row]-row == columns[i]+i for some i
//
func solveNQueens(n int) [][]string {
	solutions := make([][]int, 0)

	// columns[{row number}] = {column number}
	columns := make([]int, n)

	solveHelper(0, n, &columns, &solutions)

	return convertToStr(solutions, n)
}

func solveHelper(row, n int, cols *[]int, slns *[][]int) {
	if row == n {
		// it's a solution, save a copy of it
		*slns = append(*slns, append([]int{}, *cols...))
		return
	}

	for col := 0; col < n; col++ {
		if canPlace(row, col, n, *cols) {
			(*cols)[row] = col
			solveHelper(row+1, n, cols, slns)
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

func convertToStr(solutions [][]int, n int) [][]string {
	res := make([][]string, 0, n)

	for _, s := range solutions {
		sStr := make([]string, 0, n)
		for _, col := range s {
			str := ""
			for i := 0; i < col; i++ {
				str += "."
			}
			str += "Q"
			for i := col + 1; i < n; i++ {
				str += "."
			}
			sStr = append(sStr, str)
		}
		res = append(res, sStr)
	}

	return res
}

func main() {
	fmt.Println(solveNQueens(4))
}
