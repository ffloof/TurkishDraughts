package ui

import (
	"TurkishDraughts/Board"
	"TurkishDraughts/UI/Theme"

	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

var currentTheme DrawTheme =  theme.LichessTheme{}

func drawBoard(imd *imdraw.IMDraw){
	currentTheme.DrawBoard(imd)
}

func drawPieces(imd *imdraw.IMDraw, b *board.BoardState){
	currentTheme.DrawPieces(imd, b)
}

func drawControls(imd *imdraw.IMDraw, win *pixelgl.Window, black bool, white bool){
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(Height+20, Height-30), basicAtlas)
	basicTxt.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	fmt.Fprintln(basicTxt, "[+,-] AI Depth:", board.MaxDepth)
	fmt.Fprintln(basicTxt, "[1] Black AI Moves:", black)
	fmt.Fprintln(basicTxt, "[2] White AI Moves:", white)
	fmt.Fprintln(basicTxt, "[Z] Undo Move")
	//fmt.Fprintln(basicTxt, "[Z] Undo")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2))
}

func getMouseData(win *pixelgl.Window) (bool, bool, int) {
	return currentTheme.GetMouseData(win)
}

func drawSelected(imd *imdraw.IMDraw, index int) {
	currentTheme.DrawSelected(imd, index)
}

func drawMoves(imd *imdraw.IMDraw, index int, moveMap map[int][]int){
	currentTheme.DrawMoves(imd, index, moveMap)
}

func drawChecks(imd *imdraw.IMDraw, moveMap map[int][]int) {
	currentTheme.DrawChecks(imd, moveMap)
}