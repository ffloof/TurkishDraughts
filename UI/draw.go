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

var (
	Border = 20.0
	Gaps = 4.0
	
	BoardBg = color.RGBA{0xED, 0xEB, 0xE9, 0xFF}
	TileDark = color.RGBA{0xB5, 0x88, 0x63, 0xFF}
	TileLight = color.RGBA{0xF0, 0xD9, 0xB5, 0xFF}
	
	TeamWhite = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	TeamBlack = color.RGBA{0x00, 0x00, 0x00, 0xFF}

	TakeColor = color.RGBA{0xFF, 0x00, 0x00, 0x7F}
	MoveColor = color.RGBA{0x00, 0x3F, 0x00, 0x7F}

	TextColor = color.RGBA{0x00, 0x00, 0x00, 0xFF}

	CornerSize = 20.0

	HoverSize = 2.0

	InternalGap = 4.0
)

func getTileSpace() float64 {
	return (float64(Height) - (2*Border) + Gaps)/ 8.0
}

func getTileSize() float64{
	return (getTileSpace() - Gaps)
}

func getTilePosBL(x int, y int) (float64, float64) {
	return float64(x)*getTileSpace() + Border, float64(7-y)*getTileSpace() + Border
	
}

func getTilePosCenter(x int, y int) (float64, float64) {
	return (float64(x)+0.5)*getTileSpace() + Border, (float64(7-y)+0.5)*getTileSpace() + Border
}


func drawBoard(imd *imdraw.IMDraw){
	currentTheme.DrawBoard(imd)
}

func drawPieces(imd *imdraw.IMDraw, b *board.BoardState){
	currentTheme.DrawPieces(imd, b)
}

func drawControls(imd *imdraw.IMDraw, win *pixelgl.Window, black bool, white bool){
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(Height+20, Height-30), basicAtlas)
	basicTxt.Color = TextColor
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

func corners(imd *imdraw.IMDraw, x1, y1 float64){
	size := getTileSize()
	x2 := x1 + size
	y2 := y1 + size

	//Bottom left corner
	imd.Push(pixel.V(x1+CornerSize,y1), pixel.V(x1,y1), pixel.V(x1,y1+CornerSize))
	imd.Polygon(0)

	//Bottom right corner
	imd.Push(pixel.V(x2-CornerSize,y1), pixel.V(x2,y1), pixel.V(x2,y1+CornerSize))
	imd.Polygon(0)

	//Top left corner
	imd.Push(pixel.V(x1+CornerSize,y2), pixel.V(x1,y2), pixel.V(x1,y2-CornerSize))
	imd.Polygon(0)

	//Top right corner
	imd.Push(pixel.V(x2-CornerSize,y2), pixel.V(x2,y2), pixel.V(x2,y2-CornerSize))
	imd.Polygon(0)
}