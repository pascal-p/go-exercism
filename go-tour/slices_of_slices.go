package main

/*
 3.14 Slices of slices

 Slices can contain any type, including other slices.

*/

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for ix := 0; ix < len(board); ix++ {
		fmt.Printf("%s\n", strings.Join(board[ix], " "))
	}
}

// go run slices_of_slices.go
// X _ X
// O _ X
// _ _ O
