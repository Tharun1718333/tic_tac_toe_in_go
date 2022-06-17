package main

import (
	"math/rand"
)

type RandomPlayer struct {
	Player
}

func (p RandomPlayer) makeamove(b Board) Move {
	var allPossiblePositions []Move
	for l := 0; l < 3; l++ {
		for m := 0; m < 3; m++ {
			if b.board[l][m] == 0 {
				allPossiblePositions = append(allPossiblePositions, Move{l, m})
			}
		}
	}
	return allPossiblePositions[rand.Intn(len(allPossiblePositions))]
}
func (p RandomPlayer) getSymbol() byte {
	return p.Symbol
}
func (p RandomPlayer) getName() string {
	return p.Name
}
