package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	// args := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}

	// args := []string{"1........", ".2.......", "..3......", "...4.....", "....5....", ".....6...", "......7..", ".......8.", "........9"}

	// args := []string{"8........", "..36.....", ".7..9.2..", ".5...7...", "....457..", "...1...3.", "..1....68", "..85...1.", ".9....4.."}

	if !isValid(args) {
		fmt.Println("Error")
		return
	}

	board := strToNumSlice(args)

	if !boardValid(board) {
		fmt.Println("Error")
		return
	}

	// fmt.Println(board)

	if solveSudoku(board) {
		print(board)
	} else {
		fmt.Println("Error")
	}

}

func isSafe(board [][]int, col, row, num int) bool {
	for i := 0; i < 9; i++ {
		if num == board[row][i] {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if num == board[j][col] {
			return false
		}
	}

	boxRowStart := row - row%3
	boxColStart := col - col%3

	for r := boxRowStart; r < boxRowStart+3; r++ {
		for d := boxColStart; d < boxColStart+3; d++ {
			if board[r][d] == num {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]int) bool {
	// check if not full
	isFull := true
	col := -1
	row := -1
	for i, arr := range board {
		for j, nb := range arr {
			if nb == 0 {
				isFull = false
				row = i
				col = j
				break
			}
		}
	}

	if isFull {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isSafe(board, col, row, num) {
			board[row][col] = num

			if solveSudoku(board) {
				return true
			}

			board[row][col] = 0
		}
	}

	return false
}

// print [][]int
func print(board [][]int) {
	for _, arr := range board {
		for i, nb := range arr {
			fmt.Printf("%v", nb)

			if i != 8 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

// strToNumSlice - copy args to 2d int slice
func strToNumSlice(args []string) [][]int {
	result := make([][]int, 9)
	for i := range result {
		result[i] = make([]int, 9)
	}

	for i, str := range args {
		for j, c := range str {
			if c == '.' {
				result[i][j] = 0
				continue
			}

			result[i][j] = int(c - 48)
		}
	}

	return result
}

// check if args is valid
func isValid(args []string) bool {
	// check for args length
	if len(args) != 9 {
		return false
	}

	for _, str := range args {
		// check for each string length
		if len(str) != 9 {
			return false
		}

		// check if each character is valids
		for _, c := range str {
			if (c < '0' || c > '9') && c != '.' {
				return false
			}
		}
	}

	return true
}

func boardValid(board [][]int) bool {
	countX := 0
	countY := 0

	count := 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				countX++
			} else {
				count++
			}

			if board[j][i] == 0 {
				countY++
			}
		}
		if countX == 9 || countY == 9 {
			return false
		}

		countX = 0
		countY = 0
	}

	if count < 17 {
		return false
	}

	countB := 0
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j = j + 3 {
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					if board[x][y] == 0 {
						countB++
					}
				}
			}

			if countB == 9 {
				return false
			}

			countB = 0
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 8; j++ {
			for x := j + 1; x < 9; x++ {
				if board[i][j] == board[i][x] && board[i][j] != 0 {
					return false
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 8; j++ {
			for x := j + 1; x < 9; x++ {
				if board[j][i] == board[x][i] && board[j][i] != 0 {
					return false
				}
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j = j + 3 {
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					// fmt.Printf("x - %v , y- %v \n", board[x][y], board[y][x])
					for a := i; a < i+3; a++ {
						for b := j; b < j+3; b++ {
							if (board[x][y] == board[a][b] && board[x][y] != 0) && (x != a && y != b) {
								return false
							}
						}
					}
				}
			}
		}
	}

	return true
}
