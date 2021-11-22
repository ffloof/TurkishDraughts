package ui

import (
	"TurkishDraughts/Board"

	"image/color"
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

const (
	Width = 1600
	Height = 900
)

var basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

func Init() {
	board.MaxDepth = 8
	b := board.CreateStartingBoard()
	
	cfg := pixelgl.WindowConfig{
		Title:  "Turkish Draughts Engine",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	
	//Search variables
	searching := false
	totalMoves := 0
	possibleMoves := []PossibleMove{} 

	output := make(chan PossibleMove)
	
	//Interacting variables
	selectedTileIndex := -1
	var moveMap map[int][]int 
	isTakeMap := false

	autoMoveWhite := false
	autoMoveBlack := false
	previousBoards := []board.BoardState{}
	var nextPrevBoard board.BoardState

	for !win.Closed() {
		//Pre drawing logic
		imd := imdraw.New(nil)

		if moveMap == nil {
			nextPrevBoard = b
			moveMap = ValidUiTakes(&b, -1, [2]int{0,0})
			isTakeMap = true
		}
		if len(moveMap) == 0 { 
			moveMap = ValidUiMoves(&b)
			isTakeMap = false
		}


		//Drawing logic
		win.Clear(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})

		currentTheme.DrawBoard(imd)	
		currentTheme.DrawSelected(imd, selectedTileIndex)
		if isTakeMap { currentTheme.DrawChecks(imd, moveMap) }
		clicked, released, tileIndex := currentTheme.GetMouseData(win)
		currentTheme.DrawMoves(imd, selectedTileIndex, moveMap)
		currentTheme.DrawPieces(imd, &b)
		drawControls(imd, win, autoMoveBlack, autoMoveWhite)

		if win.JustPressed(pixelgl.Key1) { autoMoveBlack = !autoMoveBlack }
		if win.JustPressed(pixelgl.Key2) { autoMoveWhite = !autoMoveWhite }
		if win.JustPressed(pixelgl.KeyMinus) {
			board.MaxDepth -= 1
			if board.MaxDepth < 0 {
				board.MaxDepth = 0
			}
		}
		if win.JustPressed(pixelgl.KeyEqual) { board.MaxDepth += 1 }
		if win.JustPressed(pixelgl.KeyZ) {
			if len(previousBoards) > 0 {
				b = previousBoards[len(previousBoards)-1]
				previousBoards = previousBoards[0:len(previousBoards)-1]
				selectedTileIndex = -1
				moveMap = nil
			}
		}

		imd.Draw(win)
		win.Update()

		gameWon, _, gameDraw := b.PlayerHasWon()
		if gameWon || gameDraw { continue }

		if (!autoMoveWhite && b.Turn == board.White) || (!autoMoveBlack && b.Turn == board.Black) {
			//User input
			if contains(moveMap[selectedTileIndex], tileIndex) {
				if clicked || released {
					if contains(moveMap[selectedTileIndex], tileIndex) {
						swapTeams, prevDirection := tryMove(&b, selectedTileIndex, tileIndex)
						moveMap = ValidUiTakes(&b, tileIndex, prevDirection)
						if swapTeams || len(moveMap) == 0 {
							selectedTileIndex = -1
							moveMap = nil
							b.SwapTeam()
							previousBoards = append(previousBoards, nextPrevBoard)
						} else {
							selectedTileIndex = tileIndex
						}
					}
				}
			} else {
				if clicked {
					selectedTileIndex = tileIndex
				}
			}

		} else {
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
				totalMoves = Search(b, output)
				//Start searching board states
			}

			//Add auto pick move logic
			if totalMoves == len(possibleMoves) {
				searching = false

				var bestMove PossibleMove 

				for i, checkMove := range possibleMoves {
					if i == 0 || (b.Turn == board.White && checkMove.value > bestMove.value) || (b.Turn == board.Black && checkMove.value < bestMove.value) {
						bestMove = checkMove
					}
				}
				b = bestMove.board
				selectedTileIndex = -1
				moveMap = nil
				previousBoards = append(previousBoards, nextPrevBoard)
			}
		}
	}
}

func contains(a []int, b int) bool {
	for _, v := range a {
		if b == v { return true }
	}
	return false
}

func tryMove(b *board.BoardState, fromIndex, toIndex int) (bool, [2]int) {
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

	moveDir := [2]int{changeIndex%8, changeIndex/8}
	moveDir = [2]int{ moveDir[0]/abs(moveDir[0]), moveDir[1]/abs(moveDir[1]) }
	return swapTeam, moveDir
}

func abs(a int) int {
	if a == 0 { return 1 }
	if a < 0 { return -a }
	return a
}

func drawControls(imd *imdraw.IMDraw, win *pixelgl.Window, black bool, white bool){
	basicTxt := text.New(pixel.V(Height+20, Height-30), basicAtlas)
	basicTxt.Color = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	fmt.Fprintln(basicTxt, "[+,-] AI Depth:", board.MaxDepth)
	fmt.Fprintln(basicTxt, "[1] Black AI Moves:", black)
	fmt.Fprintln(basicTxt, "[2] White AI Moves:", white)
	fmt.Fprintln(basicTxt, "[Z] Undo Move")
	//fmt.Fprintln(basicTxt, "[Z] Undo")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2))
}