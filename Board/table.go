package board

type storedState struct {
	board BoardState //Stores board state since hash will have collisions
	value float32 //Stores previously evaluated value
	depth int32 //Stores depth to ensure that a position isnt too shallowly evaluated
}

type TransposTable struct {
	internal map[uint64]storedState
	//Set maximum depth for hashes to reduce memory by only saving computationally expensive hashes
	//Lower values lead to less memory consumption but slower computer performance
	maxHashDepth int32 //default a few layers short of full depth

	//When set to above 0 it will allow the transposition table to get values evaluated at lower depths
	//i.e. at = 2, for a 6 ply evaluation it can use previous 4 ply evaluation 
	//This introduces inaccuracy but has a massive performance gain
	//To minimize inaccuracy use a low MaximumHashDepth and a low TableDepthAllowedInaccuracy
	depthInaccuracy int32 //default 0 or 2
}

//Creates a new table and returns a pointer
func NewTable(maxhash, inaccuracy int32) *TransposTable {
	return &TransposTable{
		internal: make(map[uint64]storedState),
		maxHashDepth: maxhash,
		depthInaccuracy: inaccuracy,
	}
}

//Returns if the board has been previously evaluated, and its value
func (table *TransposTable) Request(board BoardState, depth int32) (bool, float32) {
	//Hash board state and load entry
	hash := board.hashBoard()
	entry, exists := table.internal[hash]

	if exists { //Check if it exists and if its the exact same board not just a collision 
		if entry.board == board {
			//Checks if the board was evaluated at a depth greater than or within a certain range
			//Prevents using too shallowly evaluated moves
			if entry.depth - table.depthInaccuracy <= depth { 
				return true, entry.value
			}
		}
	}
	return false, 0.0 //Otherwise returns that it didn't find anything

}

//Sets the board
func (table *TransposTable) Set(board BoardState, value float32, depth int32){
	if depth >= table.maxHashDepth { return } 

	//Hash board state and write to table
	hash := board.hashBoard()

	//Replace only if greater depth, ie more computationally expensive
	entry, exists := table.internal[hash]
	if !exists || depth <= entry.depth {
		//By saving deeper explored branches not only do we save the most time saving possiblity, we also perform far fewer writes increasing efficiency. 
		table.internal[hash] = storedState{board, value, depth}
	}
}

//TODO: write unit test(s) for this
func (table *TransposTable) Turn(){
	for a,b := range table.internal {
		b.depth += 1
		if b.depth > table.maxHashDepth {
			delete(table.internal, a)
		} else {
			table.internal[a] = b
		}
	}
}

//63 bits of the hash are made of where pieces are (no other info about them) and 1 bit for whos turn it is
//This results in many collisions but this is necessary as no computer comes close to having enough memory to being able to address the whole 64bit range
//We only really care about storing a few expensive searches so this is ideal
func (board *BoardState) hashBoard() uint64 {
	return (board.Full << 1) | uint64(board.Turn)
}


