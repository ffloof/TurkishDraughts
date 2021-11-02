package main

import (
	"TurkishDraughts/Network"
	"TurkishDraughts/Board"
	"fmt"
)

func main() {
	b := board.CreateStartingBoard()
	fmt.Println(b.RawBoardValue())
	network.Init()
}