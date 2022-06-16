package main

import (
	"fmt"
)

type Iplayer interface {
	makeamove() Move
	getSymbol() byte
	getName() string
}
type Player struct {
	Name   string
	Symbol byte
}

func (p Player) makeamove() Move {
	var i, j int
	fmt.Scan(&i, &j)
	move := Move{i, j}
	return move
}
func (p Player) getSymbol() byte {
	return p.Symbol
}

func (p Player) getName() string {
	return p.Name
}

func (p *Player) ChangeName() {
	var Take_input string
	fmt.Scanln(&Take_input)
	p.Name = Take_input
}
