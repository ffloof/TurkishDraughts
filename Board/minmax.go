package board

const (
	AlphaBetaMax = 999.0
	WinWeight = 100.0
	KingWeight = 5.0
	PawnWeight = 1.0
	
	//Heuristic weight of how far advanced a sides pawn pieces are from promotion //TODO: make it increase as piece count decreases?
	//When its not 0.0 it makes ab pruning much slower put pushes the engine to play better in the long term
	AdvanceWeight = 0.2

	//Set minimum depth for hashes to reduce memory by only saving computationally expensive hashes
	//Greater values lead to less memory consumption but slower computer performance (0 - depth-1)
	MinimumHashDepth = 2
)

var (
	Hits = 0
	Searches = 0
)

//TODO: check if promotion check is applied at end of take as well
func (bs *BoardState) MinMax(depth uint32, alpha float32, beta float32, table *TransposTable) float32 {
	Hits += 1

	if alreadyChecked, prevValue := table.Request(bs); alreadyChecked {
		return prevValue
	}

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
	
	Searches += 1

	options := bs.MaxTakeBoards()
	if len(options) == 0 {
		options = bs.AllMoveBoards()

		if len(options) == 0 { //No legal move check
			if bs.Turn == White { return -WinWeight 
			} else { return WinWeight }
		}
	}
	
	var bestValue float32

	if bs.Turn == White {
		bestValue = -AlphaBetaMax
		for _, branch := range options {
			branch.SwapTeam()
			branch.TryPromotion()
			value := branch.MinMax(depth-1, alpha, beta, table)
			
			if value > bestValue { bestValue = value }
			if value > alpha { alpha = value }
			if beta <= alpha { break }
		}
	} else {
		bestValue = AlphaBetaMax
		for _, branch := range options {
			branch.SwapTeam()
			branch.TryPromotion()
			value := branch.MinMax(depth-1, alpha, beta, table)

			if value < bestValue { bestValue = value }
			if value < beta { beta = value }
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


func (bs *BoardState) RawBoardValue() float32 { //Game is always from whites perspective

	wPawns := 0
	wKings := 0
	bPawns := 0
	bKings := 0
	netAdvance := 0

	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			piece, _ := bs.GetBoardTile(x,y)
			if piece.Full == Empty { continue }
			if piece.Team == White {
				if piece.King == King {
					wKings += 1
				} else {
					wPawns += 1
					netAdvance += (7-y)
				}
			} else {
				if piece.King == King {
					bKings += 1
				} else {
					bPawns += 1
					netAdvance -= (y)
				}
			}
		}
	}
	var value float32 = 0.0

	value += PawnWeight * float32(wPawns - bPawns)
	value += KingWeight * float32(wKings - bKings)
	value += AdvanceWeight * float32(netAdvance)

	return value
}