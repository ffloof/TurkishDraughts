package ui

import (
	"TurkishDraughts/Board"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"TurkishDraughts/UI/Theme"
)

type DrawTheme interface {
	DrawBoard(*imdraw.IMDraw)
	DrawPieces(*imdraw.IMDraw, *board.BoardState)
	DrawSelected(*imdraw.IMDraw, int)
	DrawMoves(*imdraw.IMDraw, int, map[int][]int)
	DrawChecks(*imdraw.IMDraw, map[int][]int)
	GetMouseData(*pixelgl.Window) (bool, bool, int)
}

var currentTheme DrawTheme = theme.WikipediaTheme{}