package board

import (
	"math"
)

const (
	AlphaBetaMax = 1000.0
	WinWeight = 100.0
	KingWeight = 5.0
	PawnWeight = 1.0
	//Heuristic weight of how far advanced a sides pawn pieces are 0.1 per piece per tile away from side
	AdvanceWeight = 0.1 
)

var (
	Searches = 0
)


func (bs *BoardState) BoardValue(depth int, alpha float64, beta float64 /*, table *TransposTable */) float64 {
	
	/* alreadyEval, prevValue := table.Load(bs)
	if alreadyEval {
		return prevValue
	} */

	Searches += 1
	//add a check for winner here
	playerWon, winWhite := bs.PlayerHasWon()
	
	if playerWon {
		if winWhite == White {
			return WinWeight
		} 
		return -WinWeight 
	}

	if depth == 0 {
		return bs.RawBoardValue()
	}

	options := bs.MaxTakeBoards()
	if len(options) == 0 {
		options = bs.AllMoveBoards()

		if len(options) == 0 { //No legal move check
			if bs.Turn == White { return -WinWeight 
			} else { return WinWeight }
		}
	}
	
	var bestValue float64

	if bs.Turn == White {
		bestValue = -AlphaBetaMax
		for _, branch := range options {
			branch.SwapTeam()
			value := branch.BoardValue(depth-1, alpha, beta /*, table*/)
			bestValue = math.Max(bestValue, value)
			
			alpha = math.Max(alpha, value)
			if beta <= alpha { break }
		}
	} else {
		bestValue = AlphaBetaMax
		for _, branch := range options {
			branch.SwapTeam()
			value := branch.BoardValue(depth-1, alpha, beta /*, table*/)
			bestValue = math.Min(bestValue, value)

			beta = math.Min(beta, value)
			if beta <= alpha { break }

		}
	}

	/* for _, branch := range options {
		branch.SwapTeam()
		table.Store(&branch, bestValue)
	} */

	return bestValue
}


func (bs *BoardState) PlayerHasWon() (bool, TileTeam) { 
	//If either player is out of pieces they lose
	var wKings uint8 = 0
	var wPieces uint8 = 0

	var bKings uint8 = 0
	var bPieces uint8 = 0

	for i:=0;i<64;i++ {
		piece, _ := bs.GetBoardTile(i%8,i/8)
		if piece.Full == Empty { continue }
		if piece.Team == White {
			wKings += uint8(piece.King)
			wPieces += 1
		} else {
			bKings += uint8(piece.King)
			bPieces += 1
		}		
	}

	//If a player has no moves they lose lol
	if wPieces == 0 {
		return true, Black
	} 

	if bPieces == 0 {
		return true, White
	}

	//If one player has a king and the other has one piece they lose
	if wPieces == 1 {
		if bKings > 0 {
			return true, Black
		}
	}

	if bPieces == 1 {
		if wKings > 0 {
			return true, White
		}
	}

	return false, 0 //No winner
	//If a player has no playable moves they lose (checked in another part of the code)
}

func (bs *BoardState) PlayersDrawed() bool {
	//Check if players are in a stalemate / draw
	return false
}


func (bs *BoardState) RawBoardValue() float64 { //Game is always from whites perspective
	value := 0.0

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			piece, _ := bs.GetBoardTile(x,y)
			if piece.Full == Empty { continue }
			if piece.Team == White {
				if piece.King == King {
					value += KingWeight
				} else {
					value += PawnWeight
					value += float64(8-y) * PawnWeight
				}
			} else {
				if piece.King == King {
					value -= KingWeight
				} else {
					value -= PawnWeight
					value -= float64(y) * PawnWeight
				}
			}
		}
	}
	return value
}