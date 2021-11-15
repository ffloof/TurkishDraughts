package ui

func (bs *BoardState) MaxTakeBoards() map[int] {
	possibleMaxTakeBoards := []BoardState{}
	bestTake := 1 //Filters boards with no jumps

	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == Empty || piece.Team != bs.Turn { continue }
		var takes int
		var possibleTakeBoards []BoardState
		if piece.King == King {
			takes, possibleTakeBoards = bs.FindKingTakes(i%8,i/8,0,[2]int{0,0})
		} else {
			takes, possibleTakeBoards = bs.FindPawnTakes(i%8,i/8,0)
		}
		if takes > bestTake {
				bestTake = takes
				possibleMaxTakeBoards = possibleTakeBoards
		} else if takes == bestTake {
			possibleMaxTakeBoards = append(possibleMaxTakeBoards, possibleTakeBoards...)
		}
		
	}
	return possibleMaxTakeBoards
}