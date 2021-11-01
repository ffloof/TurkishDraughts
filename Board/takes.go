package board

func (bs *BoardState) MaxTakeBoards() []BoardState {
	possibleMaxTakeBoards := []BoardState{}
	bestTake := 1 //Filters boards with no jumps

	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == Empty || piece.Team != bs.Turn { continue }
		var takes int
		var possibleTakeBoards []BoardState
		//if piece.King == King {
		//	takes, possibleTakeBoards = bs.FindKingTakes(i%8,i/8,0,[2]int{0,0})
		//} else {
			takes, possibleTakeBoards = bs.FindPawnTakes(i%8,i/8,0)
		//}
		if takes > bestTake {
				bestTake = takes
				possibleMaxTakeBoards = possibleTakeBoards
		} else if takes == bestTake {
			possibleMaxTakeBoards = append(possibleMaxTakeBoards, possibleTakeBoards...)
		}
	}
	return possibleMaxTakeBoards
}

func (bs *BoardState) FindKingTakes(x int, y int, currentTakes int, lastDir [2]int) (int, []BoardState) {
	boards := []BoardState{}
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y)

	for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		if -lastDir[0] == direction[0] && -lastDir[1] == direction[1] { continue } //Check to not go backwards in a direction

		for i:=0;i<8;i++ {
			jumpPos := [2]int{direction[0]*i,direction[1]*i}
			jumpOverTile, onBoard := bs.GetBoardTile(jumpPos[0],jumpPos[1])
			if !onBoard { break }
			if jumpOverTile.Full == Empty { continue }
			if attackingTile.Team == jumpOverTile.Team { break }
			if attackingTile.Team != jumpOverTile.Team {
				//We have a jump
				for i=i+1;i<8;i++{
					landingPos := [2]int{direction[0]*i, direction[1]*i} 
					landingTile, onBoard := bs.GetBoardTile(landingPos[0], landingPos[1])
					if !onBoard { 
						break
					} else {
						if landingTile.Full == Filled { break }
						newBS := *bs
						newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
						newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{})
						newBS.SetBoardTile(x,y, Tile{})

						takes, possibleBoards := newBS.FindKingTakes(landingPos[0], landingPos[1],currentTakes+1, direction)
						
						if takes > bestTake {
							bestTake = takes
							boards = possibleBoards
						} else if takes == bestTake {
							boards = append(boards, possibleBoards...)
						}
					}
				}
			}
		}		
	}

	if len(boards) == 0 {
		return bestTake, []BoardState{ *bs }
	}
	return bestTake, boards
}

func (bs *BoardState) FindPawnTakes(x int, y int, currentTakes int) (int, []BoardState) {
	boards := []BoardState{}
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y)

	
	for _, move := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		if attackingTile.Team == White && (move[0] == 0 && move[1] == 1) { continue } //Down (black only)
		if attackingTile.Team == Black && (move[0] == 0 && move[1] == -1) { continue } //Up (white only)

		jumpPos := [2]int{ x+move[0],y+move[1] }
		landingPos := [2]int { x+(2*move[0]) , y+(2*move[1]) }
		jumpOverTile, onBoard1 := bs.GetBoardTile(jumpPos[0], jumpPos[1])
		landingTile, onBoard2 := bs.GetBoardTile(landingPos[0], landingPos[1])
		if onBoard1 && onBoard2 {
			if landingTile.Full == Empty && jumpOverTile.Full == Filled && attackingTile.Team != jumpOverTile.Team {
				newBS := *bs
				newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
				newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{})
				newBS.SetBoardTile(x,y, Tile{})
				
				takes, possibleBoards := newBS.FindPawnTakes(landingPos[0], landingPos[1],currentTakes+1)
				if takes > bestTake {
					bestTake = takes
					boards = possibleBoards
				} else if takes == bestTake {
					boards = append(boards, possibleBoards...)
				}
			}
		}
	}

	if len(boards) == 0 {
		return bestTake, []BoardState{ *bs }
	}
	return bestTake, boards
	
}