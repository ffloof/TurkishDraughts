package board

const (
	//Weights for pieces/wins
	//Win should always be greater than the theoretical max of value of a board where one side gets 16 kings
	AlphaBetaMax float32 = 999.0
	WinWeight float32 = 100.0
	KingWeight float32 = 5.0
	PawnWeight float32 = 1.0
)

var (
	//What depth minmax searches to
	MaxDepth int32 = 10 //default 10
	
	//Heuristic weight of how far advanced a sides pawn pieces are from promotion //TODO: make it increase as piece count decreases?
	//When its not 0.0 it makes ab pruning much slower put pushes the engine to play better in the long term
	AdvanceWeight float32 = 0.1 //default 0.0 or 0.1

	//Set maximum depth for hashes to reduce memory by only saving computationally expensive hashes
	//Lower values lead to less memory consumption but slower computer performance
	MaximumHashDepth int32 = 7 //default around 7

	//When set to above 0 it will allow the transposition table to get values evaluated at lower depths
	//i.e. at = 2, for a 6 ply evaluation it can use previous 4 ply evaluation 
	//This introduces inaccuracy but has a massive performance gain
	//To minimize inaccuracy use a low MaximumHashDepth and a low TableDepthAllowedInaccuracy
	TableDepthAllowedInaccuracy int32 = 0 //default 0 or 2
)



//Evaluates recursively the value of a board using the minmax algorithm
//Board value is always when in whites favor positive and blacks favor negative
func (bs *BoardState) MinMax(depth int32, alpha, beta float32, table *TransposTable) float32 {
	Hits += 1

	//Checks table to see if theres already an entry for this board
	if alreadyChecked, prevValue := table.Request(bs, depth); alreadyChecked {
		return prevValue //If there is one no need to research this branch of the tree
	}

	//Checks if the board is won or if players have drawed and returns accordingly
	playerWon, winWhite, playerDrew := bs.PlayerHasWon()
	
	if playerWon {
		if winWhite == White {
			return WinWeight
		} 
		return -WinWeight 
	} else if playerDrew {
		return 0.0
	}

	if depth == MaxDepth {
		return bs.RawBoardValue()
	}	

	Searches += 1

	options := bs.ValidPlays()
	if len(options) == 0 { //If a payer has no legal moves they lose
		if bs.Turn == White { 
			return -WinWeight 
		} else { 
			return WinWeight 
		}
	}
	
	var bestValue float32

	//Search for the best possible value move
	if bs.Turn == White {
		bestValue = -AlphaBetaMax
		for _, branch := range options { //Search each possible move with minmax
			branch.SwapTeam()
			value := branch.MinMax(depth+1, alpha, beta, table)
			
			//Cache move in the table to speed up later searches of identical situations
			if depth <= MaximumHashDepth { table.Set(&branch, value, depth+1) }
			
			//AB pruning to speed up tree search
			if value > bestValue { bestValue = value }
			if value > alpha { alpha = value }
			if beta <= alpha { break }
		}
	} else { //Same just from black's perspective
		bestValue = AlphaBetaMax
		for _, branch := range options {
			branch.SwapTeam()
			value := branch.MinMax(depth+1, alpha, beta, table)
			if depth <= MaximumHashDepth { table.Set(&branch, value, depth+1) }

			if value < bestValue { bestValue = value }
			if value < beta { beta = value }
			if beta <= alpha { break }
		}
	}
	
	//Return final best value of a move found from this board
	return bestValue
}

//Gets the value of the board by summing piece weights and how far advanced a sides pieces are
func (bs *BoardState) RawBoardValue() float32 { 
	wPawns := 0
	wKings := 0
	bPawns := 0
	bKings := 0
	netAdvance := 0

	//Sum amount of pieces and how advanced they are up the board
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

	//Some arithmetic to get the final board value
	var value float32 = 0.0

	value += PawnWeight * float32(wPawns - bPawns)
	value += KingWeight * float32(wKings - bKings)
	value += AdvanceWeight * float32(netAdvance)

	return value
}