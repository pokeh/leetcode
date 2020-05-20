package main

const emptyVal = '.'

func solveSudoku(board [][]byte) {
	helper(board)
}

func helper(board [][]byte) bool {
	rowNum, colNum := findNextSpace(board)

	if rowNum == -1 && colNum == -1 {
		return true
	}

	for val := byte('1'); val <= byte('9'); val++ {
		if valid(val, rowNum, colNum, board) {
			board[rowNum][colNum] = val
			if helper(board) {
				return true
			}
			board[rowNum][colNum] = emptyVal
		}
	}

	return false
}

func findNextSpace(board [][]byte) (int, int) {
	for i, row := range board {
		for j, val := range row {
			if val == emptyVal {
				return i, j
			}
		}
	}
	return -1, -1
}

func valid(val byte, rowNum, colNum int, board [][]byte) bool {
	return validRow(val, rowNum, board) &&
		validCol(val, colNum, board) &&
		validSubgrid(val, rowNum, colNum, board)
}

func validRow(val byte, rowNum int, board [][]byte) bool {
	for _, v := range board[rowNum] {
		if val == v {
			return false
		}
	}
	return true
}

func validCol(val byte, colNum int, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if val == board[i][colNum] {
			return false
		}
	}
	return true
}

func validSubgrid(val byte, rowNum, colNum int, board [][]byte) bool {
	subgridRowStart := (rowNum / 3) * 3
	subgridColStart := (colNum / 3) * 3

	for i := subgridRowStart; i < subgridRowStart+3; i++ {
		for j := subgridColStart; j < subgridColStart+3; j++ {
			if val == board[i][j] {
				return false
			}
		}
	}
	return true
}
