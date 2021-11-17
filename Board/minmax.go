package board

const (
	AlphaBetaMax = 999.0 //default 999.0
	WinWeight = 100.0 //default 100.0
	KingWeight = 5.0 //default 5.0
	PawnWeight = 1.0 //default 1.0
	
	//Heuristic weight of how far advanced a sides pawn pieces are from promotion //TODO: make it increase as piece count decreases?
	//When its not 0.0 it makes ab pruning much slower put pushes the engine to play better in the long term
	AdvanceWeight = 0.1 //default 0.1

	//Set minimum depth for hashes to reduce memory by only saving computationally expensive hashes
	//Greater values lead to less memory consumption but slower computer performance (0 - depth-1)
	//MinimumHashDepth = 3  //default 2
	MaximumHashDepth = 7


	//DANGER ZONE BELOW:
	//When set to above 0 it will allow the transposition table to get values only evaluated at lower depths
	//i.e. at = 2, it can use at depth 6 a previous evaluation at depth 4 
	//This introduces inaccuracy but has a massive performance gain
	//To minimize inaccuracy use a low MaximumHashDepth and a low TableDepthAllowedInaccuracy
	TableDepthAllowedInaccuracy = 2 //default 0
)

var (
	Hits = 0
	Searches = 0
)

//TODO: check if promotion check is applied at end of take as well
func (bs *BoardState) MinMax(depth uint32, alpha float32, beta float32, table *TransposTable) float32 {
	Hits += 1

	if alreadyChecked, prevValue := table.request(bs, depth); alreadyChecked {
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
			value := branch.MinMax(depth-1, alpha, beta, table)
			
			if value > bestValue { bestValue = value }
			if value > alpha { alpha = value }
			if beta <= alpha { break }
		}
	} else {
		bestValue = AlphaBetaMax
		for _, branch := range options {
			branch.SwapTeam()
			value := branch.MinMax(depth-1, alpha, beta, table)

			if value < bestValue { bestValue = value }
			if value < beta { beta = value }
			if beta <= alpha { break }

		}
	}

	if depth <= MaximumHashDepth {
		for _, branch := range options {
			//branch.SwapTeam()
			table.set(&branch, bestValue, depth)
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