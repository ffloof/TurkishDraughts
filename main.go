package main

import (
	"TurkishDraughts/Network"
	"TurkishDraughts/Board"
)



func main() {
	//network.Init()

	
	//Ai plays against itself for my amusement
	//b := board.CreateStartingBoard()
	b := board.BoardFromStr("-------- -------- -------- W------- -------- --w----- --www--- -----B--")
	b.Turn = board.White
	for true {
		b = *(network.Analyze(b, 10))
		b.Print()
		//winner, _ := b.PlayerHasWon()
		//if winner { break }
		if b == board.BoardFromStr("-------- -------- -------- -------- -------- -------- -------- --------") { break }
	}
}