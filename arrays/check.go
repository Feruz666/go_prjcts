package main

import "fmt"


func main() {
	var board [8][8]string
	var sudoku [8][8]int 

	board[0][0] = "r"
	board[0][7] = "r"

	for column := range board[1]{
		board[1][column] = "p"
	}
	fmt.Print(board)
	fmt.Print(sudoku)
}