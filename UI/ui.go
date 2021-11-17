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
	forcedIndex := false

	for !win.Closed() {
		//Pre drawing logic
		imd := imdraw.New(nil)
		moving := false

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
		drawMoves(imd, selectedTileIndex, moveMap)

		drawPieces(imd, &b)

		imd.Draw(win)
		win.Update()

		winner, _ := b.PlayerHasWon()
		if winner { continue }

		//User input 
		if clicked { //Clicking
			if selectedTileIndex != -1 {
				if tileIndex != selectedTileIndex {
					moving = true
				}
			} else if _, moveExists := moveMap[tileIndex]; moveExists {
				if (!forcedIndex) { selectedTileIndex = tileIndex }
			}
		}
		if released && tileIndex != selectedTileIndex { //Dragging
			moving = true
		}

		if moving {
			if contains(moveMap[selectedTileIndex], tileIndex) {
				swapTeams := tryMove(&b, selectedTileIndex, tileIndex)
				moveMap = ValidUiTakes(&b)
				if swapTeams || len(moveMap) == 0 {
					forcedIndex = false
					moveMap = nil
					b.SwapTeam()
				} else {
					forcedIndex = true
				}
			}
			selectedTileIndex = tileIndex			
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

func contains(a []int, b int) bool {
	for _, v := range a {
		if b == v { return true }
	}
	return false
}

func tryMove(b *board.BoardState, fromIndex, toIndex int) bool {
	tile, _ := b.GetBoardTile(fromIndex%8, fromIndex/8)
	b.SetBoardTile(toIndex%8, toIndex/8, tile)
	b.SetBoardTile(fromIndex%8, fromIndex/8, board.Tile{})
	
	changeIndex := toIndex - fromIndex
	change := 0

	if changeIndex >= 8 || changeIndex <= -8 {
		change = 8
	} else {
		change = 1
	}
	if changeIndex < 0 {
		change *= -1
	}

	swapTeam := true
	for i:=fromIndex; i!=toIndex; i+=change{
		tile, _ := b.GetBoardTile(i%8, i/8)
		if tile.Full == board.Filled {
			swapTeam = false
		}
		b.SetBoardTile(i%8, i/8, board.Tile{})
	}
	return swapTeam
}

