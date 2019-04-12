package blockchain

import (
	. "github.com/LowEntropyBody/basic-blockchain/block"
)

// Blockchain struct
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain - construct
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// AddBlock - add block to chain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
