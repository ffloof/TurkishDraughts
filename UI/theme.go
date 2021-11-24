package ui

import (
	"TurkishDraughts/Board"

	"image/color"
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"github.com/faiface/pixel/pixelgl"
)

type DrawTheme interface {
	DrawBoard(*imdraw.IMDraw)
	DrawPieces(*imdraw.IMDraw, *board.BoardState)
	DrawSelected(*imdraw.IMDraw, int)
	DrawMoves(*imdraw.IMDraw, int, map[int][]int)
	DrawChecks(*imdraw.IMDraw, map[int][]int)
	GetMouseData(*pixelgl.Window) (bool, bool, int)
}

func drawControls(imd *imdraw.IMDraw, win *pixelgl.Window, black bool, white bool){
	basicTxt := text.New(pixel.V(Height+20, Height-30), basicAtlas)
	basicTxt.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	fmt.Fprintln(basicTxt, "[+,-] AI Depth:", board.MaxDepth)
	fmt.Fprintln(basicTxt, "[1] Black AI Moves:", black)
	fmt.Fprintln(basicTxt, "[2] White AI Moves:", white)
	fmt.Fprintln(basicTxt, "[Z] Undo Move")
	fmt.Fprintln(basicTxt, "[<,>] Change Theme")
	//fmt.Fprintln(basicTxt, "[Z] Undo")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2))
}