package main

import (
	"TurkishDraughts/Network"
	"TurkishDraughts/Board"
)



func main() {
	//network.Init()

	
	//Ai plays against itself for my amusement
	b := board.CreateStartingBoard()
	for true {
		b = *(network.Analyze(b, 10))
		b.Print()
		winner, _ := b.PlayerHasWon()
		if winner { break }
	}
}