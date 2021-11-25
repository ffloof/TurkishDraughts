package board

//Takes a king at a certain position and finds all the possible resulting boards
func (bs *BoardState) AllMovesKing(x, y int) []BoardState {
	boards := []BoardState{}
	checkingTile, _ := bs.GetBoardTile(x,y)

	//Loops through every direction it could move
	for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} { //Right
		for i:=1;i<8;i++ {
			//Since its a king it can move i in any direction
			move := [2]int{direction[0]*i, direction[1]*i}
			moveTile, onBoard := bs.GetBoardTile(x + move[0],y + move[1])

			//If the tile its trying to move to is empty add the valid move 
			//otherwise we hit a wall, stop and try the next direction
			if moveTile.Full == Empty && onBoard {
				newBS := *bs
				newBS.SetBoardTile(x + move[0], y + move[1], checkingTile)
				newBS.SetBoardTile(x, y, Tile{})
				boards = append(boards, newBS)
			} else {
				break //Hit wall, next direction
			}
		}
	}

	return boards
}

//Takes a pawn at a certain position and finds all the possible resulting boards
func (bs *BoardState) AllMovesPawn(x, y int) []BoardState {
	boards := []BoardState{}
	checkingTile, _ := bs.GetBoardTile(x,y)

	//Look at each possible direction it could move
	for _, move := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		//Check that pawns aren't moving backwards
		if checkingTile.Team == White && move[1] == 1 { continue }
		if checkingTile.Team == Black && move[1] == -1 { continue }
		
		//Checks if the move is valid and appends the possible board
		moveTile, onBoard := bs.GetBoardTile(x + move[0],y + move[1])
		if moveTile.Full == Empty && onBoard {
			newBS := *bs
			newBS.SetBoardTile(x + move[0], y + move[1], checkingTile)
			newBS.SetBoardTile(x, y, Tile{})
			boards = append(boards, newBS)
		}
	}

	return boards
}


//Gets all possible moves that could be made on a turn
func (bs *BoardState) AllMoveBoards() []BoardState {
	possibleMoveBoards := []BoardState {}
	//Loop through all pieces
	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		//Skips pieces that are empty or are from the team color that isn't moving
		if piece.Full == Empty || piece.Team != bs.Turn { continue } 
		
		//Adds possible moves each piece could make
		if piece.King == King {
			possibleMoveBoards = append(possibleMoveBoards, bs.AllMovesKing(i%8,i/8)...)
		} else {
			possibleMoveBoards = append(possibleMoveBoards, bs.AllMovesPawn(i%8,i/8)...)
		}
		
	}
	return possibleMoveBoards
}