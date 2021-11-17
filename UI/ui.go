package ui

import (
	"TurkishDraughts/Board"

	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	Width = 1600
	Height = 900
)


func Init() {
	b := board.CreateStartingBoard()
	//b := board.BoardFromStr("-------- bbbb-bbb -------- bbbbbbbb wwwwwwww -------- wwwwwwww --------")
	//b.SwapTeam()

	cfg := pixelgl.WindowConfig{
		Title:  "Turkish Draughts Engine",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	/*
	//Search variables
	searching := false
	totalMoves := 0
	possibleMoves := []PossibleMove{} 

	output := make(chan PossibleMove)
	quit := make(chan bool) */

	//Interacting variables
	lastTileIndex := -1
	var takeMap map[int][]int
	var moveMap map[int][]int 

	for !win.Closed() {
		imd := imdraw.New(nil)
		if takeMap == nil { takeMap = ValidUiTakes(&b)}
		if moveMap == nil { moveMap = ValidUiMoves(&b)}

		//Drawing logic
		win.Clear(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})

		drawBoard(imd)
		drawPieces(&b, imd)
		drawChecks(imd, takeMap)
		drawSelected(imd, lastTileIndex)

		clicked, tileIndex := drawHover(win, imd)

		if clicked {
			if lastTileIndex != -1 {
				if tileIndex == lastTileIndex {
					lastTileIndex = -1
				} else {
					//Move if clicked on a valid tile
				}
			} else if t, _ := b.GetBoardTile(tileIndex%8, tileIndex/8); t.Full == board.Filled && t.Team == b.Turn {
				_, takeExists := takeMap[tileIndex]
				moves, _ := moveMap[tileIndex]

				if takeExists || (len(takeMap) == 0 && len(moves) != 0) {
					lastTileIndex = tileIndex
				}
			}
		}

		imd.Draw(win)
		win.Update()

		winner, _ := b.PlayerHasWon()
		if winner { continue }

		/*
		//Engine logic
		if totalMoves != len(possibleMoves) {
			//Check if theres a result
			select {
			case pMove := <-output:
				possibleMoves = append(possibleMoves, pMove)
			}
		} else if !searching {
			searching = true
			possibleMoves = []PossibleMove{}
			totalMoves = Search(b, 9, quit, output)
			//Start searching board states
		}

		//Add auto pick move logic
		if totalMoves == len(possibleMoves) {
			searching = false

			var bestMove PossibleMove 

			for i, checkMove := range possibleMoves {
				//TODO: add cool arrows with numbers
				if i == 0 || (b.Turn == board.White && checkMove.value > bestMove.value) || (b.Turn == board.Black && checkMove.value < bestMove.value) {
					bestMove = checkMove
				}
			}
			b = bestMove.board
		}*/


	}
}
