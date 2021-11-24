package ui

import (
	"TurkishDraughts/Board"
	"TurkishDraughts/UI/Theme"

	"image/color"

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

var basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII) //Font
var themes = []DrawTheme{ theme.LichessTheme{}, theme.WikipediaTheme{}} //Different themes to cycle through

func Init() {
	board.MaxDepth = 8 //Depth to search to for AI
	
	//Create game window and starting board
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

	
	//AI move search variables
	searching := false
	totalMoves := 0
	possibleMoves := []PossibleMove{} //When the amount of total moves = the amount of possibleMoves, the search has completed

	output := make(chan PossibleMove) //Used to pass evaluations back to this rendering thread
	
	//Human player interacting variables
	selectedTileIndex := -1 //Tile player is currently clicked on
	var moveMap map[int][]int //Valid options player can currently select
	isTakeMap := false

	//Controls variables
	autoMoveWhite := false //If ai plays for white
	autoMoveBlack := false //If ai plays for black
	themeIndex := 1 //Current theme selected
	previousBoards := []board.BoardState{} //Previous boards for undo feature
	var nextPrevBoard board.BoardState
	var lastEval *PossibleMove

	//Draw loop
	for !win.Closed() {
		//If a list of valid moves from this board hasnt been created, create one
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
		
		//New blank window
		imd := imdraw.New(nil)
		win.Clear(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})

		//Select theme and then draw board and pieces
		currentTheme := themes[themeIndex] 
		currentTheme.DrawBoard(imd)	
		if selectedTileIndex != -1 { currentTheme.DrawSelected(imd, selectedTileIndex) }
		if isTakeMap { currentTheme.DrawChecks(imd, moveMap) }
		//Mouse data is theme specific since tiles can be in slightly different alignments
		if selectedTileIndex != -1 { currentTheme.DrawMoves(imd, selectedTileIndex, moveMap) }
		clicked, released, tileIndex := currentTheme.GetMouseData(win) 
		currentTheme.DrawPieces(imd, &b)
		drawControls(imd, win, autoMoveBlack, autoMoveWhite, lastEval)

		//Control logic
		if win.JustPressed(pixelgl.Key1) { autoMoveBlack = !autoMoveBlack } //Toggle ai white
		if win.JustPressed(pixelgl.Key2) { autoMoveWhite = !autoMoveWhite } //Toggle ai black
		if win.JustPressed(pixelgl.KeyMinus) { //Decrement search depth
			board.MaxDepth -= 1
			if board.MaxDepth < 0 {
				board.MaxDepth = 0
			}
		}
		if win.JustPressed(pixelgl.KeyEqual) { board.MaxDepth += 1 } //Increment search depth
		if win.JustPressed(pixelgl.KeyZ) { //Undo move
			if len(previousBoards) > 0 {
				b = previousBoards[len(previousBoards)-1]
				previousBoards = previousBoards[0:len(previousBoards)-1]
				selectedTileIndex = -1
				moveMap = nil
				lastEval = nil
			}
		}

		if win.JustPressed(pixelgl.KeyLeft) { //Cycle prev theme
			themeIndex--
			if themeIndex < 0 {
				themeIndex = len(themes) - 1
			}
		}
		if win.JustPressed(pixelgl.KeyRight) { //Cycle next theme
			themeIndex++
			if themeIndex >= len(themes) {
				themeIndex = 0
			}
		}

		//Finish drawing
		imd.Draw(win)
		win.Update()

		//Check if the game is won, if it is don't allow any user input or ai moves
		gameWon, _, gameDraw := b.PlayerHasWon()
		if gameWon || gameDraw { continue }

		//User input

		//Check that ai isn't playing for the current side
		if (!autoMoveWhite && b.Turn == board.White) || (!autoMoveBlack && b.Turn == board.Black) {
			//Select a tile if we haven't clicked on a move square 
			if contains(moveMap[selectedTileIndex], tileIndex) {
				if clicked || released {
					//Make the move
					swapTeams, prevDirection := tryMove(&b, selectedTileIndex, tileIndex)
					moveMap = ValidUiTakes(&b, tileIndex, prevDirection)
					if swapTeams || len(moveMap) == 0 {
						selectedTileIndex = -1
						moveMap = nil
						lastEval = nil
						b.SwapTeam()
						previousBoards = append(previousBoards, nextPrevBoard)
					} else {
						selectedTileIndex = tileIndex
					}
				}
			} else {
				if clicked {
					selectedTileIndex = tileIndex
				}
			}
		} else { //In the event that the ai is playing
			//Engine logic

			//If we are still expecting more moves to be evaluated do nothing
			if totalMoves != len(possibleMoves) {
				//Check if theres a new result
				select {
				case pMove := <-output:
					//Add it to the list of known results
					possibleMoves = append(possibleMoves, pMove)
				}
			} else if !searching { //Otherwise if we haven't started searching, start
				searching = true
				possibleMoves = []PossibleMove{}
				totalMoves = Search(b, output)
				//Start searching board states
			} 

			//Once all moves are on pick best move for the given side
			if totalMoves == len(possibleMoves) {
				searching = false

				var bestMove PossibleMove 

				for i, checkMove := range possibleMoves {
					if i == 0 || (b.Turn == board.White && checkMove.value > bestMove.value) || (b.Turn == board.Black && checkMove.value < bestMove.value) {
						bestMove = checkMove
					}
				}

				//Update to new board, and reset some variables
				b = bestMove.board
				selectedTileIndex = -1
				moveMap = nil
				lastEval = &bestMove
				previousBoards = append(previousBoards, nextPrevBoard)
			}
		}
	}
}



//Moves a piece and if it took a piece it returns that it can move again, and the direction of last move
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

//Absolute value function for integers
func abs(a int) int {
	if a == 0 { return 1 }
	if a < 0 { return -a }
	return a
}

//Just a quick function to check if a slice contains an entry
func contains(a []int, b int) bool {
	for _, v := range a {
		if b == v { return true }
	}
	return false
}