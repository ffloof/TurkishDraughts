package board

import (
	"sync"
)

type storedState struct {
	board BoardState
	value float64
}

type TransposTable struct {
	lock sync.RWMutex
	internal map[uint64]storedState
}

func NewTable() *TransposTable {
	return &TransposTable{
		internal: make(map[uint64]storedState),
	}
}

func (table *TransposTable) Request(board *BoardState) (bool, float64) {
	//Hash board state and load entry
	hash := board.hashBoard()
	table.lock.RLock()
	entry, exists := table.internal[hash]
	table.lock.RUnlock()

	if exists {
		if entry.board == *board {
			return true, entry.value
		}
	}
	return false, 0.0

}

func (table *TransposTable) Set(board *BoardState, value float64){
	//Hash board state and write to table
	hash := board.hashBoard()
	table.lock.Lock()
	table.internal[hash] = storedState{*board, value}
	table.lock.Unlock()
}


func (board *BoardState) hashBoard2() uint64 {
	return 0
}

func (board *BoardState) hashBoard() uint64 {
	hash := board.Full << 1
	hash |= uint64(board.Turn)
	return hash
}