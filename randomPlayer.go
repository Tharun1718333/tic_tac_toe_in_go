package main

import (
	"math/rand"
)

type RandomPlayer struct {
	Player
}

func (p RandomPlayer) makeamove() Move {
	var i, j int
	i = rand.Intn(3)
	j = rand.Intn(3)
	move := Move{i, j}
	return move
}
func (p RandomPlayer) getSymbol() byte {
	return p.Symbol
}
func (p RandomPlayer) getName() string {
	return p.Name
}
