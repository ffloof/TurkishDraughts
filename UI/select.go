package ui

import (
	"TurkishDraughts/Board"
)

func ValidUiTakes(bs *board.BoardState) map[int][]int {
	bestTake := 1 //Filters boards with no jumps
	validTakes := map[int][]int{}

	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == board.Empty || piece.Team != bs.Turn { continue }
		var takes int
		var validTakePos []int
		if piece.King == board.King {
			takes, _, validTakePos = bs.FindKingTakes(i%8,i/8,0,[2]int{0,0})
		} else {
			takes, _, validTakePos = bs.FindPawnTakes(i%8,i/8,0)
		}
		if takes > bestTake {
			validTakes = map[int][]int{
				i: validTakePos,
			}
		} else if takes == bestTake {
			validTakes[i] = validTakePos
		}
		
	}
	return validTakes
}