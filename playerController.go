package main

type PlayerController struct {
	PlayerCollection []Player
	Curr             Player
}

func (c PlayerController) IteratePlayer() Player {
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
