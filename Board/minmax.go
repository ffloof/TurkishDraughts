package board

import (
	"math"
)

const (
	AlphaBetaMax = 9999.0
	WinWeight = 1000.0
	KingWeight = 50.0
	PawnWeight = 10.0
	
	//Heuristic weight of how far advanced a sides pawn pieces are 0.1 per piece per tile away from side
	//For some reason when its not 0.0 it makes ab pruning less efficient, but incentivises agressive play
	AdvanceWeight = 2.0

	//Set minimum depth for hashes to reduce memory by only saving computationally expensive hashes
	//Greater values lead to less memory consumption but slower computer performance (0 - depth-1)
	MinimumHashDepth = 3
)

var (
	Searches = 0
)

//Minimax function TODO: convert to negamax?
func (bs *BoardState) MinMax(depth uint32, alpha float64, beta float64, table *TransposTable) float64 {
	if alreadyChecked, prevValue := table.Request(bs); alreadyChecked {
		return prevValue
	}

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
			value := branch.MinMax(depth-1, alpha, beta, table)
			bestValue = math.Max(bestValue, value)
			
			alpha = math.Max(alpha, value)
			if beta <= alpha { break }
		}
	} else {
		bestValue = AlphaBetaMax
		for _, branch := range options {
			branch.SwapTeam()
			value := branch.MinMax(depth-1, alpha, beta, table)
			bestValue = math.Min(bestValue, value)

			beta = math.Min(beta, value)
			if beta <= alpha { break }

		}
	}

	if depth > MinimumHashDepth {
		for _, branch := range options {
			branch.SwapTeam()
			table.Set(&branch, bestValue, depth)
		}
	}
	

	return bestValue
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
					value += float64(7-y) * AdvanceWeight
				}
			} else {
				if piece.King == King {
					value -= KingWeight
				} else {
					value -= PawnWeight
					value -= float64(y) * AdvanceWeight
				}
			}
		}
	}
	return value
}