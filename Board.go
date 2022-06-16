package main

import "fmt"

type BoardFunctions interface {
	PrintBoard()
	ApplyTheMove()
	ValidateTheMove()
}
type Board struct {
	GameBoard [3][3]byte
}

func (b Board) PrintBoard() {
	for i := 0; i < int(len(b.GameBoard)); i++ {
		for j := 0; j < int(len(b.GameBoard[0])); j++ {
			if string(b.GameBoard[i][j]) == "X" || string(b.GameBoard[i][j]) == "O" {
				fmt.Print(string(b.GameBoard[i][j]))
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
	b.GameBoard[i][j] = k
}
func (b Board) ValidateTheMove(i int, j int) bool {
	if i < 0 || j < 0 || i > 2 || j > 2 {
		return false
	}
	if b.GameBoard[i][j] >= 65 && b.GameBoard[i][j] <= 90 {
		return false
	}
	return true
}

// implement a counter
func (b Board) IsFilled() bool {
	count := 0
	for i := 0; i < int(len(b.GameBoard)); i++ {
		for j := 0; j < int(len(b.GameBoard[0])); j++ {
			if b.GameBoard[i][j] >= 65 && b.GameBoard[i][j] <= 90 {
				count++
			}
		}
	}
	return count == (int(len(b.GameBoard)) * int(len(b.GameBoard[0])))
}
