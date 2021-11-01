package board

import (
	//"sync"
)

type Entry struct {
	Exists bool
	Board BoardState
	Value float64
}

type TransposTable struct {
	//sync.RWMutex
	internal []Entry 
}

const size = 131072 // 2^17, 17 bit keys


func NewTable() *TransposTable {
	return &TransposTable{
		internal: make([]Entry, size),
	}
}

func (table *TransposTable) read(key uint64) Entry {
	//table.RLock()
	entry := table.internal[key]
	//table.RUnlock()
	return entry
}

func (table *TransposTable) delete(key uint64) {
	//table.Lock()
	table.internal[key] = Entry{}
	//table.Unlock()
}

func (table *TransposTable) write(key uint64, entry Entry) {
	//table.Lock()
	table.internal[key] = entry
	//table.Unlock()
}

func (b *BoardState) hashBoard() uint64 { //Does a hash based on columns and rows of filled positions % 2
	var hash uint64

	var x uint64
	var y uint64
	for y=0;y<8;y++ {
		var sum uint64 = 0
		for x=0;x<8;x++ {
			sum += (b.Full >> (y*8+x)) & 1
		}
		sum %= 2
		hash = hash << 1
		hash |= sum
	}

	for x=0;x<8;x++ {
		var sum uint64 = 0
		for y=0;y<8;y++ {
			sum += (b.Full >> (y*8+x)) & 1
		}
		sum %= 2
		hash = hash << 1
		hash |= sum
	}

	hash = hash << 1
	hash |= uint64(b.Turn)
	return hash
}

func (table *TransposTable) Load(b *BoardState) (bool, float64) {
	//Hash board state and request
	hash := b.hashBoard()
	entry := table.read(hash)

	//Check if it exists and if the stored board states are identical
	if entry.Exists && *b == entry.Board {
		return true, entry.Value 
	}
	return false, 0.0
}

func (table *TransposTable) Store(b *BoardState, value float64){
	//Hash board state and write to table
	hash := b.hashBoard()
	table.write(hash, Entry {true, *b, value})
}