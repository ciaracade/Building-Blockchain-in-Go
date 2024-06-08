package main

import (

)

// Create Blockchain -> data of block structures
type Blockchain struct {
	blocks []*Block
}

// Adds blocks to blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Creates Genesis block (first block in the blockchain to add to)
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Creates block with Genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}