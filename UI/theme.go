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
	DrawBoard(*imdraw.IMDraw) //Draws the checkerboard
	DrawPieces(*imdraw.IMDraw, *board.BoardState) //Draws the pieces on it
	DrawSelected(*imdraw.IMDraw, int) //Draws the marker on the currently selected tile
	DrawMoves(*imdraw.IMDraw, int, map[int][]int) //If a tile is selected it draws the moves for that tile
	DrawChecks(*imdraw.IMDraw, map[int][]int) //If there are moves a player has to make highlight the pieces they originate from
	GetMouseData(*pixelgl.Window) (bool, bool, int) //Get whether the player clicked, released, and what tile index they are hovering over
}

//Control panel is the same for all themes
func drawControls(imd *imdraw.IMDraw, win *pixelgl.Window, black bool, white bool, lastEval *PossibleMove){
	basicTxt := text.New(pixel.V(Height+20, Height-30), basicAtlas)
	basicTxt.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	fmt.Fprintln(basicTxt, "[+,-] AI Depth:", board.MaxDepth)
	fmt.Fprintln(basicTxt, "[1] Black AI Moves:", black)
	fmt.Fprintln(basicTxt, "[2] White AI Moves:", white)
	fmt.Fprintln(basicTxt, "[Z] Undo Move")
	fmt.Fprintln(basicTxt, "[<,>] Change Theme")
	if lastEval != nil {
		fmt.Fprintln(basicTxt, "AI Move Evaluation:", lastEval.value)
	}
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2))
}