package main

import (
	"time"
	"crypto/sha256"
	"strconv"
	"bytes"
)

// Create blocks that store valuable information
type Block struct {
	Timestamp 		int64	// When block is created
	Data			[]byte	// Valuable information in block
	PrevBlockHash 	[]byte	// Previous block hash
	Hash			[]byte 	// Block headers
}

// Create func that handles calculating hashes for blocks
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// Create new block via previous block hash 
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
