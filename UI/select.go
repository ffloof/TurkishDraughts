package ui

import (
	"TurkishDraughts/Board"
)

func ValidUiTakes(bs *board.BoardState, forcedIndex int, lastDir [2]int) map[int][]int {
	bestTake := 1 //Filters boards with no jumps
	validTakes := map[int][]int{}

	for i:=0;i<64;i++ {
		if forcedIndex != -1 { i = forcedIndex }
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == board.Empty || piece.Team != bs.Turn { continue }
		var takes int
		var validTakePos []int
		if piece.King == board.King {
			takes, _, validTakePos = bs.FindKingTakes(i%8,i/8,0,lastDir)
		} else {
			takes, _, validTakePos = bs.FindPawnTakes(i%8,i/8,0)
		}
		if takes > bestTake {
			bestTake = takes
			validTakes = map[int][]int{
				i: validTakePos,
			}
		} else if takes == bestTake {
			validTakes[i] = validTakePos
		}
		if forcedIndex != -1 { break }
	}
	return validTakes
}

func ValidUiMoves(bs *board.BoardState) map[int][]int {
	validMoves := map[int][]int{}

	for a:=0;a<64;a++ {
		piece, _ := bs.GetBoardTile(a%8,a/8)
		if piece.Full == board.Empty || piece.Team != bs.Turn { continue }

		stepMax := 1
		if piece.King == board.King { stepMax = 7 }
		
		x, y := a%8, a/8
		moveList := []int{} 
		for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
			for b:=1;b<=stepMax;b++ {
				if piece.King == board.Pawn {
					if piece.Team == board.White && (direction[0] == 0 && direction[1] == 1) { continue } //Down (black only)
					if piece.Team == board.Black && (direction[0] == 0 && direction[1] == -1) { continue } //Up (white only)
				}
				moveX := (direction[0]*b) + x
				moveY := (direction[1]*b) + y
				moveTile, onBoard := bs.GetBoardTile(moveX,moveY)
				if moveTile.Full == board.Empty && onBoard {
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