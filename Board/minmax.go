package board

import (
	"math"
)

const (
	AlphaBetaMax = 1000.0
	WinWeight = 100.0
	KingWeight = 5.0
	PawnWeight = 1.0
)


//TODO: add multiple possibility return
func (bs *BoardState) BoardValue(depth int, alpha float64, beta float64, turnTeam Team) float64 {
	//add a check for winner here
	winState := bs.PlayerHasWon()
	if winState == White { 
		return WinWeight
	} else if winState == Black { 
		return -WinWeight 
	}

	if depth == 0 {
		return bs.RawBoardValue()
	}

	options := bs.MaxTakeBoards(turnTeam)
	if len(options) == 0 {
		options = bs.AllMoveBoards(turnTeam)

		if len(options) == 0 { //No legal move check
			if turnTeam == White { return -WinWeight }
			if turnTeam == Black { return WinWeight }
		}
	}
	
	var bestValue float64

	if turnTeam == White {
		bestValue = -AlphaBetaMax
		for _, branch := range options {
			value := branch.BoardValue(depth-1, alpha, beta, Black)
			bestValue = math.Max(bestValue, value)
			
			alpha = math.Max(alpha, value)
			if beta <= alpha { break }
		}
	} else if turnTeam == Black {
		bestValue = AlphaBetaMax
		for _, branch := range options {
			value := branch.BoardValue(depth-1, alpha, beta, White)
			bestValue = math.Min(bestValue, value)

			beta = math.Min(beta, value)
			if beta <= alpha { break }

		}
	}
	return bestValue
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
				value += KingWeight
			} else {
				value += PawnWeight
			}
		} else if piece.Team == Black {
			if piece.King {
				value -= KingWeight
			} else {
				value -= PawnWeight
			}
		}
	}
	return value
}