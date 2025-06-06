package main

import "fmt"

func main() {
	isValid := isValidSudoku([][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	})

	fmt.Println(isValid)
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := board[i]
		col := []byte{board[0][i], board[1][i], board[2][i], board[3][i], board[4][i], board[5][i], board[6][i], board[7][i], board[8][i]}

		if !(valid(row) && valid(col)) {
			return false
		}

		if (i+1)%3 == 0 {
			grid := []byte{board[0][i-2], board[0][i-1], board[0][i], board[1][i-2], board[1][i-1], board[1][i], board[2][i-2], board[2][i-1], board[2][i]}
			if !valid(grid) {
				return false
			}
			grid = []byte{board[3][i-2], board[3][i-1], board[3][i], board[4][i-2], board[4][i-1], board[4][i], board[5][i-2], board[5][i-1], board[5][i]}
			if !valid(grid) {
				return false
			}
			grid = []byte{board[6][i-2], board[6][i-1], board[6][i], board[7][i-2], board[7][i-1], board[7][i], board[8][i-2], board[8][i-1], board[8][i]}
			if !valid(grid) {
				return false
			}
		}
	}

	return true
}

func valid(row []byte) bool {
	for i := 0; i < 9; i++ {
		bytes := make(map[byte]struct{}, 9)
		for _, elem := range row {
			if elem == '.' {
				continue
			}

			if _, ok := bytes[elem]; ok {
				return false
			}

			bytes[elem] = struct{}{}
		}
	}

	return true
}
