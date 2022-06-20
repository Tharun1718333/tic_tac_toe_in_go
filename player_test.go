package main

import "testing"

func TestPlayerGetName(t *testing.T) {
	p1 := Player{"p1", 'X'}
	if p1.getName() != "p1" {
		t.Errorf("failed to get name")
	}
}
func TestPlayerGetSymoblo(t *testing.T) {
	p1 := Player{"p1", 'X'}
	if p1.getSymbol() != 'X' {
		t.Errorf("failed to get symbol")
	}
}
