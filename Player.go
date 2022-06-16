package main

import "fmt"

type Player struct {
	Name   string
	Symbol byte
}
type CurrentPlayer struct {
	PlayerCollection []Player
	Curr             Player
}

func (p *Player) ChangeName() {
	var Take_input string
	fmt.Scanln(&Take_input)
	p.Name = Take_input
}
func (c CurrentPlayer) IteratePlayer() Player {
	var i int
	for ; i < len(c.PlayerCollection); i++ {
		if c.PlayerCollection[i] == c.Curr {
			break
		}
	}
	if i == len(c.PlayerCollection)-1 {
		return c.PlayerCollection[0]
	}
	return c.PlayerCollection[i+1]
}
