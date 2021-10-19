package board

func (bs *BoardState) MaxTakeBoards(turnTeam Team) []BoardState {
	possibleMaxTakeBoards := []BoardState{}
	bestTake := 1 //Filters boards with no jumps

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			if turnTeam != bs[(y*8)+x].Team { continue }
			var takes int
			var possibleTakeBoards []BoardState
			if bs[(y*8)+ x].King {
				takes, possibleTakeBoards = bs.FindKingTakes(x,y,0,[2]int{0,0})
			} else {
				takes, possibleTakeBoards = bs.FindPawnTakes(x,y,0)
			}
			if takes > bestTake {
					bestTake = takes
					possibleMaxTakeBoards = possibleTakeBoards
			} else if takes == bestTake {
				possibleMaxTakeBoards = append(possibleMaxTakeBoards, possibleTakeBoards...)
			}
		}
	}

	return possibleMaxTakeBoards
}

func (bs *BoardState) FindKingTakes(x int, y int, currentTakes int, lastDir [2]int) (int, []BoardState) {
	boards := []BoardState{ *bs }
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y) //TODO: add error checks for not on board and empty tiles

	for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		if -lastDir[0] == direction[0] && -lastDir[1] == direction[1] { continue } //Check to not go backwards in a direction

		for i:=0;i<8;i++ {
			jumpPos := [2]int{direction[0]*i,direction[1]*i}
			jumpOverTile, onBoard := bs.GetBoardTile(jumpPos[0],jumpPos[1])
			if !onBoard { break }
			if jumpOverTile.Team == Empty { continue }
			if attackingTile.Team == jumpOverTile.Team { break }
			if attackingTile.Team != jumpOverTile.Team {
				//We have a jump
				for i=i+1;i<8;i++{
					landingPos := [2]int{direction[0]*i, direction[1]*i} 
					landingTile, onBoard := bs.GetBoardTile(landingPos[0], landingPos[1])
					if !onBoard { 
						break
					} else {
						if landingTile.Team != Empty { break }
						newBS := *bs
						newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
						newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{Empty, false})
						newBS.SetBoardTile(x,y, Tile{Empty, false})

						takes, possibleBoards := newBS.FindKingTakes(landingPos[0], landingPos[1],currentTakes+1, direction)
						
						if takes > bestTake {
							bestTake = takes
							boards = possibleBoards
						}
						if takes == bestTake {
							boards = append(boards, possibleBoards...)
						}
					}
				}
			}
		}		
	}

	return bestTake, boards

	return 0, []BoardState{}
}

func (bs *BoardState) FindPawnTakes(x int, y int, currentTakes int) (int, []BoardState) {
	boards := []BoardState{ *bs }
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y) //TODO: add error checks for not on board and empty tiles

	
	for _, move := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		if attackingTile.Team == White && !(move[0] == 0 && move[1] == 1) { continue } //Down (black only)
		if attackingTile.Team == Black && !(move[0] == 0 && move[1] == -1) { continue } //Up (white only)

		jumpPos := [2]int{ x+move[0],y+move[1] }
		landingPos := [2]int { x+(2*move[0]) , y+(2*move[1]) }
		jumpOverTile, onBoard1 := bs.GetBoardTile(jumpPos[0], jumpPos[1])
		landingTile, onBoard2 := bs.GetBoardTile(landingPos[0], landingPos[1])
		if onBoard1 && onBoard2 {
			if landingTile.Team == Empty && jumpOverTile.Team != Empty && attackingTile.Team != jumpOverTile.Team {
				newBS := *bs
				newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
				newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{Empty, false})
				newBS.SetBoardTile(x,y, Tile{Empty, false})
				
				takes, possibleBoards := newBS.FindPawnTakes(landingPos[0], landingPos[1],currentTakes+1)
				if takes > bestTake {
					bestTake = takes
					boards = possibleBoards
				}
				if takes == bestTake {
					boards = append(boards, possibleBoards...)
				}
			}
		}
	}

	return bestTake, boards
}