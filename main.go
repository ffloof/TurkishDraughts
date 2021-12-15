package main

import (
	"TurkishDraughts/Board"
	//"TurkishDraughts/UI"
	//"github.com/faiface/pixel/pixelgl"
)


func main() {
	b := board.CreateStartingBoard()
	board.MCTS(b, 10000)
	//pixelgl.Run(ui.Init)
}



