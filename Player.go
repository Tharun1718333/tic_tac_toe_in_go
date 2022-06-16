package main

import "fmt"

type Player struct {
	Name   string
	Symbol byte
}

func (p *Player) ChangeName() {
	var Take_input string
	fmt.Scanln(&Take_input)
	p.Name = Take_input
}
