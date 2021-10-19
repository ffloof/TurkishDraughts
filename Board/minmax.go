package board

//TODO: add multiple possibility return
func (bs *BoardState) BoardValue(depth int, turnTeam Team) (float64, BoardState) {
	//add a check for winner here
	winState := bs.PlayerHasWon()
	if winState == White { 
		return 100.0, *bs 
	} else if winState == Black { 
		return -100.0, *bs 
	}

	if depth == 0 {
		return bs.RawBoardValue(), *bs
	}

	options := bs.MaxTakeBoards(turnTeam)
	if len(options) == 0 {
		options = bs.AllMoveBoards(turnTeam)

		if len(options) == 0 { //No legal move check
			if turnTeam == White { return -100.0, *bs }
			if turnTeam == Black { return 100.0, *bs }
		}
	}
	

	var bestValue float64
	var bestBranch BoardState

	for i, branch := range options{
		if turnTeam == White {
			value, _ := branch.BoardValue(depth-1, Black)
			if i==0 || value >= bestValue {
				bestValue = value //White tries to maximize value
				bestBranch = branch
			}
		} else if turnTeam == Black {
			value, _ := branch.BoardValue(depth-1, White)
			if i==0 || value <= bestValue {
				bestValue = value //Black tries to minimize value
				bestBranch = branch
			}
		}
	}
	
	return bestValue, bestBranch;
}


func (bs *BoardState) PlayerHasWon() Team { //0 = no winner 1 = white wins 2 = black wins
	//If either player is out of pieces they lose
	wKings := 0
	wPieces := 0

	bKings := 0
	bPieces := 0

	for _, piece := range bs {
		if piece.Team == White {
			if piece.King {
				wKings += 1
			}
			wPieces += 1
		} else if piece.Team == Black {
			if piece.King {
				bKings += 1
			}
			bPieces += 1
		}
	}

	//If a player has no moves they lose lol
	if wPieces == 0 {
		return Black
	} 

	if bPieces == 0 {
		return White
	}

	//If one player has a king and the other has one piece they lose
	if wPieces == 1 {
		if bKings > 0 {
			return Black
		}
	}

	if bPieces == 1 {
		if wKings > 0 {
			return White
		}
	}

	return Empty //No winner
	//If a player has no playable moves they lose (checked in another part of the code)
}

//TODO: Draw check would optimize returning 0 instead of worthless move searches
func (bs *BoardState) PlayersDrawed() bool {
	//Check if players are in a stalemate / draw
	return false
}

func (bs *BoardState) RawBoardValue() float64 { //Game is always from whites perspective
	value := 0.0
	for _, piece := range bs {
		if piece.Team == White {
			if piece.King {
				value += 5.0
			} else {
				value += 1.0
			}
		} else if piece.Team == Black {
			if piece.King {
				value -= 5.0
			} else {
				value -= 1.0
			}
		}
	}
	return value
}