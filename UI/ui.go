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
	selectedTileIndex := -1
	var moveMap map[int][]int 
	isTakeMap := false

	for !win.Closed() {
		imd := imdraw.New(nil)
		if moveMap == nil { 
			moveMap = ValidUiTakes(&b)
			isTakeMap = true
		}
		if len(moveMap) == 0 { 
			moveMap = ValidUiMoves(&b)
			isTakeMap = false
		}


		//Drawing logic
		win.Clear(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})

		drawBoard(imd)
		drawSelected(imd, selectedTileIndex)

		if isTakeMap { drawChecks(imd, moveMap) }
		
		clicked, released, tileIndex := getMouseData(win)
		if selectedTileIndex != -1 { 
			drawMoves(imd, selectedTileIndex, moveMap) 
		}

		drawPieces(imd, &b)

		imd.Draw(win)
		win.Update()

		winner, _ := b.PlayerHasWon()
		if winner { continue }

		if clicked {
			if selectedTileIndex != -1 {
				if tileIndex == selectedTileIndex {
					selectedTileIndex = -1
				} else {
					tryMove(&b, tileIndex)
				}
			} else if _, moveExists := moveMap[tileIndex]; moveExists {
				selectedTileIndex = tileIndex
			}
		}

		if released && tileIndex != selectedTileIndex {
			tryMove(&b, tileIndex)
		}


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

func tryMove(b *board.BoardState, tileIndex int){

}