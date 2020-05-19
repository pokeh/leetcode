package main

func main() {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solveSudoku(board)
}

func solveSudoku(board [][]int) {
	helper(board)
}

func helper(board [][]int) {
	// TODO: how to call this recursively to achieve backtracking?
	// for row := 0; row < 9; row++ {
	// 	for col := 0; col < 9; col++ {
	// 		for val := 1; val < 10; val++ {
	// 			if valid(val, col, row, board) {
	// 				board[row][col] = val
	// 				helper(board, hogehoge)
	// 			}
	// 		}
	// 	}
	// }
}

func valid(val, row, col int, board [][]int) bool {
	return validRow(val, row, board) &&
		validCol(val, col, board) &&
		validSubgrid(val, row, col, board)
}

func validRow(val, row int, board [][]int) bool {
	for _, v := range board[row] {
		if val == v {
			return false
		}
	}
	return true
}

func validCol(val, col int, board [][]int) bool {
	for i := 0; i < 9; i++ {
		if val == board[i][col] {
			return false
		}
	}
	return true
}

func validSubgrid(val, row, col int, board [][]int) bool {
	subgridRowStart := (row / 3) * 3
	subgridColStart := (col / 3) * 3

	for i := subgridRowStart; i < subgridRowStart+3; i++ {
		for j := subgridColStart; j < subgridColStart+3; j++ {
			if val == board[i][j] {
				return false
			}
		}
	}
	return true
}
