package main

import "fmt"

func main() {
	board := make([][]rune, 8)
	for i := range board {
		board[i] = make([]rune, 8)
		for j := range board[i] {
			board[i][j] = '#'
		}
	}
	Read("https://lichess.org/txLb0xmp", board)
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println()
	}
}
