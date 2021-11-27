package main

import (
	"TurkishDraughts/Board"
	"fmt"
	//"TurkishDraughts/UI"
	//"github.com/faiface/pixel/pixelgl"
)


func main() {
	b := board.BoardFromStr("-------- ---b---- ---b---- -bbW-bb- -------- ---b---- ---b---- --------")
	b.SwapTeam()

	for i:=0;i<64;i++ {
		tile, _ := b.GetBoardTile(i%8,i/8)
		if tile.Full == board.Filled && tile.Team == board.White { fmt.Println(i)}
	}

	for _, v := range b.MaxTakeBoards() {
		v.PrintSingleLine()
	}
	//pixelgl.Run(ui.Init)
}



