package board

//Find all possible takes that result in the maximum amount
func (bs *BoardState) MaxTakeBoards() []BoardState {
	possibleMaxTakeBoards := []BoardState{}
	bestTake := 1 //Filters boards with no jumps

	//Loop through all pieces
	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		//Skip pieces that are empty or from the team not moving
		if piece.Full == Empty || piece.Team != bs.Turn { continue }
		var takes int
		var possibleTakeBoards []BoardState
		//Recursively find final outcomes of every take combination
		if piece.King == King {
			takes, possibleTakeBoards, _ = bs.FindKingTakes(i%8,i/8,0,[2]int{0,0})
		} else {
			takes, possibleTakeBoards, _ = bs.FindPawnTakes(i%8,i/8,0)
		}

		//Only combinations with the most takes are kept
		if takes > bestTake {
				bestTake = takes
				possibleMaxTakeBoards = possibleTakeBoards
		} else if takes == bestTake {
			possibleMaxTakeBoards = append(possibleMaxTakeBoards, possibleTakeBoards...)
		}
	}
	return possibleMaxTakeBoards
}

//Finds all possible takes from a king at a given position
//Returns # of takes, final boards, and if at starting depth what moves lead to max take outcomes (used in ui code)
func (bs *BoardState) FindKingTakes(x int, y int, currentTakes int, lastDir [2]int) (int, []BoardState, []int) {
	newTakeBoards := []BoardState{}
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y)
	validTakePos := []int{}

	//Loop through each possible direction king can take in
	for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		//King can't take in a 180 from previous take (so it cant take, north then take south)
		if -lastDir[0] == direction[0] && -lastDir[1] == direction[1] { continue } //Check to not go backwards in a direction

		var jumpPos [2]int
		var jumpOverTile Tile

		//Searches in the direction until it finds a piece to jump over
		i:=1
		exit := false //Exit indicates its either hit a friendly piece or a wall and that it should search the next direction
		for i<8{
			var onBoard bool
			jumpPos = [2]int{x+(direction[0]*i),y+(direction[1]*i)}
			jumpOverTile, onBoard = bs.GetBoardTile(jumpPos[0],jumpPos[1])

			i++

			if jumpOverTile.Full == Empty { continue }
			//Hit friendly piece or is searching off the board
			if !onBoard || attackingTile.Team == jumpOverTile.Team { exit = true }
			break
		}
		
		if exit { continue } //Search next direction
		
		//We have a possible piece to take
		//Search each tile behind it to see if theres valid tile to land on, without jumping over multiple pieces
		for i<8 {
			landingPos := [2]int{x+(direction[0]*i), y+(direction[1]*i)} 
			landingTile, onBoard := bs.GetBoardTile(landingPos[0], landingPos[1])

			i++

			//Cant jump over multiple pieces or off the board so stops
			if !onBoard || landingTile.Full == Filled { break }
			
			//Otherwise it has a valid take position and stores the board
			newBS := *bs
			newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
			newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{})
			newBS.SetBoardTile(x,y, Tile{})

			//Check if from that position there are any new takes recusrively
			takes, possibleBoards, _ := newBS.FindKingTakes(landingPos[0], landingPos[1],currentTakes+1, direction)
			
			//Store only the boards that had the most amount of takes
			if takes > bestTake {
				if currentTakes == 0 { //If this is the starting node store the move made (for ui)
					validTakePos = []int{(landingPos[1]*8)+landingPos[0]}
				}
				bestTake = takes
				newTakeBoards = possibleBoards
			} else if takes == bestTake {
				if currentTakes == 0 { //If this is the starting node store the move made (for ui)
					validTakePos = append(validTakePos, (landingPos[1]*8)+landingPos[0])
				}
				newTakeBoards = append(newTakeBoards, possibleBoards...)
			}
		}		
	}

	if len(newTakeBoards) == 0 { //If there weren't any new takes found return end node
		return bestTake, []BoardState{ *bs }, validTakePos
	}
	//Otherwise return all the takes that resulted in the max amount of pieces taken searched and the amount
	return bestTake, newTakeBoards, validTakePos 
}

//Finds recursively all possible takes from a pawn at a given position
//Returns # of takes, final boards, and if at starting depth what moves lead to max take outcomes (used in ui code) 
func (bs *BoardState) FindPawnTakes(x int, y int, currentTakes int) (int, []BoardState, []int) {
	newTakeBoards := []BoardState{}
	bestTake := currentTakes
	attackingTile, _ := bs.GetBoardTile(x,y)
	validTakePos := []int{}
	
	//Loops through each direction a pawn could take in
	for _, move := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		//Pawns can't take backwards
		if attackingTile.Team == White && (move[0] == 0 && move[1] == 1) { continue } //Down (black only)
		if attackingTile.Team == Black && (move[0] == 0 && move[1] == -1) { continue } //Up (white only)

		//Gets positions and tiles that are being jumped over and to
		jumpPos := [2]int{ x+move[0],y+move[1] }
		landingPos := [2]int { x+(2*move[0]) , y+(2*move[1]) }
		jumpOverTile, onBoard1 := bs.GetBoardTile(jumpPos[0], jumpPos[1])
		landingTile, onBoard2 := bs.GetBoardTile(landingPos[0], landingPos[1])

		//Check if the take is valid
		if onBoard1 && onBoard2 {
			if landingTile.Full == Empty && jumpOverTile.Full == Filled && attackingTile.Team != jumpOverTile.Team {
				//If it is create a new board that stores the outcome of that take
				newBS := *bs
				newBS.SetBoardTile(landingPos[0], landingPos[1], attackingTile)
				newBS.SetBoardTile(jumpPos[0], jumpPos[1], Tile{})
				newBS.SetBoardTile(x,y, Tile{})
				
				//Then check again if theres any possible takes from that board at the pawns new position
				takes, possibleBoards, _ := newBS.FindPawnTakes(landingPos[0], landingPos[1],currentTakes+1)

				//Only store the max take combinations
				if takes > bestTake {
					if currentTakes == 0 { //If this is the starting node store the move made (for ui)
						validTakePos = []int{(landingPos[1]*8)+landingPos[0]}
					}
					bestTake = takes
					newTakeBoards = possibleBoards
				} else if takes == bestTake {
					if currentTakes == 0 { //If this is the starting node store the move made (for ui)
						validTakePos = append(validTakePos, (landingPos[1]*8)+landingPos[0])
					}
					newTakeBoards = append(newTakeBoards, possibleBoards...)
				}
			}
		}
	}

	if len(newTakeBoards) == 0 { //If there weren't any new takes found return end node
		return bestTake, []BoardState{ *bs }, validTakePos
	}
	//Otherwise return all the takes that resulted in the max amount of pieces taken searched and the amount
	return bestTake, newTakeBoards, validTakePos
}