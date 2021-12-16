package main

import (
	//"TurkishDraughts/Tournament"
	//"TurkishDraughts/UI"
	//"github.com/faiface/pixel/pixelgl"

	"TurkishDraughts/Board"
	"fmt"
)


func main() {
	//pixelgl.Run(ui.Init)
	//tournament.Run()

	b1 := board.BoardFromStr("-------- bbbbbbbb -bbbbbbb b------- -ww----- w--wwwww wwwwwwww --------")

	b2 := board.BoardFromStr("-------- bbbbbbbb --bbbbbb bb------ -ww----- w--wwwww wwwwwwww --------")
	b2.SwapTeam()
	v1, _ := b1.MinMax(0,-999.0,999.0,board.NewTable())
	v2, _ := b2.MinMax(1,-999.0,0.0,board.NewTable())

	fmt.Println(v1)
	//bs1[1].Print()
	fmt.Println()

	fmt.Println(v2)
	//bs2[0].Print()
	fmt.Println()
}



