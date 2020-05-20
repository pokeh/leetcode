package main

import "fmt"

const emptyVal = 0

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

	solution := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	if equal(board, solution) {
		fmt.Println("yay!")
	} else {
		fmt.Println("boo!")
	}
}

func solveSudoku(board [][]int) {
	helper(board)
}

func helper(board [][]int) bool {
	rowNum, colNum := findNextSpace(board)

	if rowNum == -1 && colNum == -1 {
		return true
	}

	for val := 1; val <= 9; val++ {
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

func findNextSpace(board [][]int) (int, int) {
	for i, row := range board {
		for j, val := range row {
			if val == emptyVal {
				return i, j
			}
		}
	}

	return -1, -1
}

func valid(val, rowNum, colNum int, board [][]int) bool {
	return validRow(val, rowNum, board) &&
		validCol(val, colNum, board) &&
		validSubgrid(val, rowNum, colNum, board)
}

func validRow(val, rowNum int, board [][]int) bool {
	for _, v := range board[rowNum] {
		if val == v {
			return false
		}
	}
	return true
}

func validCol(val, colNum int, board [][]int) bool {
	for i := 0; i < 9; i++ {
		if val == board[i][colNum] {
			return false
		}
	}
	return true
}

func validSubgrid(val, rowNum, colNum int, board [][]int) bool {
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

func equal(b1, b2 [][]int) bool {
	for i, row := range b1 {
		for j := range row {
			if b1[i][j] != b2[i][j] {
				return false
			}
		}
	}
	return true
}
