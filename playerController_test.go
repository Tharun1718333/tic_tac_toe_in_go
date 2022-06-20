package main

import "testing"

func TestIteratePlayer(t *testing.T) {
	p1 := Player{"p1", 'X'}
	p2 := Player{"p2", '0'}
	playerStack := []Iplayer{p1, p2}
	playercontroller := PlayerController{playerStack, p1}
	p3 := playercontroller.IteratePlayer()
	if p3 != p2 {
		t.Errorf("failed to iterate the player")
	}
}
