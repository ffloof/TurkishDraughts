package board

func (bs *BoardState) AllMovesKing(x int, y int) []BoardState {
	boards := []BoardState{}
	checkingTile, _ := bs.GetBoardTile(x,y)

	for _, direction := range [4][2]int {{0,1},{0,-1},{-1,0},{1,0},} { //Right
		for i:=1;i<8;i++ {
			move := [2]int{direction[0]*i, direction[1]*i}
			moveTile, onBoard := bs.GetBoardTile(x + move[0],y + move[1])
			if moveTile.Team == Empty && onBoard {
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
		if moveTile.Team == Empty && onBoard {
			newBS := *bs
			king := checkingTile.King
			if checkingTile.Team == White && y + move[1] == 0 { //Promote to king condition
				king = true
			} else if checkingTile.Team == Black && y + move[1] == 7 {
				king = true
			}
			newBS.SetBoardTile(x + move[0], y + move[1], Tile{checkingTile.Team, king})
			newBS.SetBoardTile(x, y, Tile{})
			boards = append(boards, newBS)
		}
	}

	return boards
}

func (bs *BoardState) AllMoveBoards(turnTeam Team) []BoardState {
	possibleMoveBoards := []BoardState {}
	for y := 0; y<8; y++ {
		for x := 0; x<8; x++ {
			if turnTeam != bs[(y*8)+x].Team { continue }
			if bs[(y*8)+x].King {
				possibleMoveBoards = append(possibleMoveBoards, bs.AllMovesKing(x, y)...)
			} else {
				possibleMoveBoards = append(possibleMoveBoards, bs.AllMovesPawn(x, y)...)
			}
		}
	}
	return possibleMoveBoards
}