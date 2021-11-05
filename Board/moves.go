package board

func (bs *BoardState) AllMovesKing(x int, y int) []BoardState {
	boards := []BoardState{}
	checkingTile, _ := bs.GetBoardTile(x,y)

	for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} { //Right
		for i:=1;i<8;i++ {
			move := [2]int{direction[0]*i, direction[1]*i}
			moveTile, onBoard := bs.GetBoardTile(x + move[0],y + move[1])
			if moveTile.Full == Empty && onBoard {
				newBS := *bs
				newBS.SetBoardTile(x + move[0], y + move[1], checkingTile)
				newBS.SetBoardTile(x, y, Tile{})
				boards = append(boards, newBS)
			} else {
				break //Stops going in this direction after it hits wall or piece
			}
		}
	}

	return boards
}

func (bs *BoardState) AllMovesPawn(x int, y int) []BoardState {
	boards := []BoardState{}
	checkingTile, _ := bs.GetBoardTile(x,y)

	for _, move := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} {
		if checkingTile.Team == White && move[0] == 0 && move[1] == 1 { continue }
		if checkingTile.Team == Black && move[0] == 0 && move[1] == -1 { continue }
		
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

func (bs *BoardState) AllMoveBoards() []BoardState {
	possibleMoveBoards := []BoardState {}
	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == Empty { continue }
		if piece.Team != bs.Turn { continue }
		if piece.King == King {
			possibleMoveBoards = append(possibleMoveBoards, bs.AllMovesKing(i%8,i/8)...)
		} else {
			possibleMoveBoards = append(possibleMoveBoards, bs.AllMovesPawn(i%8,i/8)...)
		}
		
	}
	return possibleMoveBoards
}