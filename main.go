package main

import (
	"TurkishDraughts/UI"
	"github.com/faiface/pixel/pixelgl"
)


func main() {
	pixelgl.Run(ui.Init)
	//network.Init()
	
	//Ai plays against itself for my amusement
	//b := board.CreateStartingBoard()
	//for true {
	//	b = *(network.Analyze(b, 10))
	//	b.Print()
	//	winner, _ := b.PlayerHasWon()
	//	if winner { break }
	//}

	//b := network.ParseHistory("a3-a4 h6-h5 g3-g4 b6-b5 b3-a3 b5-a5")
	//b.Print()
}



