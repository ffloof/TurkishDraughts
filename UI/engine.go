package ui

import (
	"runtime/debug"
	"TurkishDraughts/Board"
)

type PossibleMove struct {
	board board.BoardState
	value float32 
}

//Channel for sending move weight values without freezing rendering thread
func Search(b board.BoardState, output chan PossibleMove) int {
	options := b.MaxTakeBoards()
	if len(options) == 0 {
		options = b.AllMoveBoards()
	}

	for _, branch := range options{
		branch.SwapTeam()
		go analyzeBranch(branch, board.NewTable(), output)
	}

	return len(options)
}

func analyzeBranch (branch board.BoardState, table *board.TransposTable, output chan PossibleMove) {
	output <- PossibleMove {branch, branch.MinMax(0, -board.AlphaBetaMax, board.AlphaBetaMax, table)}
	debug.FreeOSMemory()
}

//Creates a map of moves that result in board takes with the maximum amount of pieces taken
func ValidUiTakes(bs *board.BoardState, forcedIndex int, lastDir [2]int) map[int][]int {
	bestTake := 1 //Filters boards with no jumps
	validTakes := map[int][]int{}

	//Loop through every possible position
	for i:=0;i<64;i++ {
		//forced index is used when a piece has already taken another piece so it limits the move search to just that piece, for subsequent takes.
		if forcedIndex != -1 { i = forcedIndex } 
		
		//Only check pieces of moving team and not empty tiles
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == board.Empty || piece.Team != bs.Turn { continue }
		
		//Get possible takes from this postion
		var takes int
		var validTakePos []int
		if piece.King == board.King {
			takes, _, validTakePos = bs.FindKingTakes(i%8,i/8,0,lastDir)
		} else {
			takes, _, validTakePos = bs.FindPawnTakes(i%8,i/8,0)
		}

		//Depending on amount of pieces taken
		if takes > bestTake { //If its more than previous max, forget about it
			bestTake = takes
			validTakes = map[int][]int{
				i: validTakePos,
			}
		} else if takes == bestTake { //If its equal add it to the map of possible takes
			validTakes[i] = validTakePos
		}
		if forcedIndex != -1 { break }
	}
	return validTakes
}

//Creates a map of all possible normal moves
func ValidUiMoves(bs *board.BoardState) map[int][]int {
	validMoves := map[int][]int{}

	//Loop through every position on the board
	for a:=0;a<64;a++ {
		//Only check pieces of moving team and not empty tiles
		piece, _ := bs.GetBoardTile(a%8,a/8)
		if piece.Full == board.Empty || piece.Team != bs.Turn { continue }

		//Depending on the type of piece how far to limit the search in each direction
		stepMax := 1
		if piece.King == board.King { stepMax = 7 }
		x, y := a%8, a/8
		moveList := []int{} 

		//Search each possible direction
		for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
			//For each direction add possible moves until we hit a piece/go off the board
			for b:=1;b<=stepMax;b++ {
				if piece.King == board.Pawn {
					//Dont allow pawns to go backward
					if piece.Team == board.White && (direction[0] == 0 && direction[1] == 1) { continue } //Down (black only)
					if piece.Team == board.Black && (direction[0] == 0 && direction[1] == -1) { continue } //Up (white only)
				}
				moveX := (direction[0]*b) + x
				moveY := (direction[1]*b) + y
				moveTile, onBoard := bs.GetBoardTile(moveX,moveY)
				if moveTile.Full == board.Empty && onBoard { //Check that we haven't hit edge of board or another piece
					moveList = append(moveList, (moveY*8)+moveX)
				} else {
					break //Stops going in this direction after it hits wall or piece
				}
			}
		}
		if len(moveList) != 0 {
			validMoves[a] = moveList
		}
	}
	return validMoves
}