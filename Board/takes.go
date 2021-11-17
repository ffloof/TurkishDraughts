package board

func (bs *BoardState) MaxTakeBoards() []BoardState {
	possibleMaxTakeBoards := []BoardState{}
	bestTake := 1 //Filters boards with no jumps

	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == Empty || piece.Team != bs.Turn { continue }
		var takes int
		var possibleTakeBoards []BoardState
		if piece.King == King {
			takes, possibleTakeBoards, _ = bs.FindKingTakes(i%8,i/8,0,[2]int{0,0})
		} else {
			takes, possibleTakeBoards, _ = bs.FindPawnTakes(i%8,i/8,0)
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

func (bs *BoardState) FindKingTakes(x int, y int, currentTakes int, lastDir [2]int) (int, []BoardState, []int) {
	boards := []BoardState{}
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y)
	validTakePos := []int{}

	for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		if -lastDir[0] == direction[0] && -lastDir[1] == direction[1] { continue } //Check to not go backwards in a direction

		var jumpPos [2]int
		var jumpOverTile Tile

		i:=1
		exit := false

		for i<8{
			var onBoard bool
			jumpPos = [2]int{x+(direction[0]*i),y+(direction[1]*i)}
			jumpOverTile, onBoard = bs.GetBoardTile(jumpPos[0],jumpPos[1])

			i++

			if !onBoard || attackingTile.Team == jumpOverTile.Team { 
				exit = true
				break
			}
			if jumpOverTile.Full == Empty { continue }
			break
		}
		
		if exit { continue }
		
		//We have a jump
		for i<8 {
			landingPos := [2]int{x+(direction[0]*i), y+(direction[1]*i)} 
			landingTile, onBoard := bs.GetBoardTile(landingPos[0], landingPos[1])

			i++

			if !onBoard || landingTile.Full == Filled { 
				exit = true
				break 
			}
			
			newBS := *bs
			newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
			newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{})
			newBS.SetBoardTile(x,y, Tile{})

			takes, possibleBoards, _ := newBS.FindKingTakes(landingPos[0], landingPos[1],currentTakes+1, direction)
			
			if takes > bestTake {
				if currentTakes == 0 {
					validTakePos = []int{(landingPos[1]*8)+landingPos[0]}
				}
				bestTake = takes
				boards = possibleBoards
			} else if takes == bestTake {
				if currentTakes == 0 {
					validTakePos = append(validTakePos, (landingPos[1]*8)+landingPos[0])
				}
				boards = append(boards, possibleBoards...)
			}
		}		
	}

	if len(boards) == 0 {
		return bestTake, []BoardState{ *bs }, validTakePos
	}
	return bestTake, boards, validTakePos
}

func (bs *BoardState) FindPawnTakes(x int, y int, currentTakes int) (int, []BoardState, []int) {
	boards := []BoardState{}
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y)
	validTakePos := []int{}
	
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
				
				takes, possibleBoards, _ := newBS.FindPawnTakes(landingPos[0], landingPos[1],currentTakes+1)
				if takes > bestTake {
					if currentTakes == 0 {
						validTakePos = []int{(landingPos[1]*8)+landingPos[0]}
					}
					bestTake = takes
					boards = possibleBoards
				} else if takes == bestTake {
					if currentTakes == 0 {
						validTakePos = append(validTakePos, (landingPos[1]*8)+landingPos[0])
					}
					boards = append(boards, possibleBoards...)
				}
			}
		}
	}

	if len(boards) == 0 {
		return bestTake, []BoardState{ *bs }, validTakePos
	}
	return bestTake, boards, validTakePos
	
}