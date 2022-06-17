package main

import "fmt"

type IBoard interface {
	PrintBoard()
	ApplyTheMove()
	ValidateTheMove()
}
type Board struct {
	board [3][3]byte
}

func (b Board) PrintBoard() {
	for i := 0; i < int(len(b.board)); i++ {
		for j := 0; j < int(len(b.board[0])); j++ {
			if string(b.board[i][j]) == "X" || string(b.board[i][j]) == "O" {
				fmt.Print(string(b.board[i][j]))
			} else {
				fmt.Print(" ")
			}
			if j < 2 {
				fmt.Print("|")
			}
		}
		if i < 2 {
			fmt.Println(" ")
			fmt.Println("-----")
		}
	}
	fmt.Println(" ")
	fmt.Println("__________________________")
}
func (b *Board) ApplyTheMove(i int, j int, k byte) {
	// fmt.Println("reached here")
	// fmt.Println(i, j)
	b.board[i][j] = k
}
func (b Board) ValidateTheMove(i int, j int) bool {
	if i < 0 || j < 0 || i > 2 || j > 2 {
		return false
	}
	if b.board[i][j] >= 65 && b.board[i][j] <= 90 {
		return false
	}
	return true
}

// implement a counter
func (b Board) IsFilled() bool {
	count := 0
	for i := 0; i < int(len(b.board)); i++ {
		for j := 0; j < int(len(b.board[0])); j++ {
			if b.board[i][j] >= 65 && b.board[i][j] <= 90 {
				count++
			}
		}
	}
	return count == (int(len(b.board)) * int(len(b.board[0])))
}
