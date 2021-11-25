package main

import (
	"TurkishDraughts/Board"
	"fmt"
	//"TurkishDraughts/UI"
	//"github.com/faiface/pixel/pixelgl"
)


func main() {
	b := board.BoardFromStr("-------- B------- -------- ---w---- -------- -------- -----b-- -bbb----")

	for i:=0;i<64;i++ {
		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Filled && tile.Team == board.White { fmt.Println(i)}
	}

	for _, v := range b.AllMoveBoards() {
		v.PrintSingleLine()
	}
	//pixelgl.Run(ui.Init)
}



