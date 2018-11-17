package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
	"time"
)

type Block struct {
	Index        int           // index in the array
	Nonce        int           // unique number for Proof of Work
	Hash         string        // Hash(block+Nonce), nonce != 0
	PrevHash     string        // block i-1
	Timestamp    int64         // unix long
	Transactions []Transaction // included transactions
}

func NewBlock() (b *Block) {
	return &Block{
		Timestamp: time.Now().Unix(),
		Nonce:     0,
	}
}

func NewBlockFromPrimitive(prevHash string, transaction []Transaction) *Block {
	block := NewBlock()
	block.PrevHash = prevHash
	block.Transactions = transaction
	return block
}
func (b *Block) CalculateHash(nonce int) {
	b.Nonce = nonce
	b.Hash = b.CalculateBlockHash()
}

func (b *Block) CalculateBlockHash() string {
	hash := make([]byte, 64)
	sha3.ShakeSum256(hash, []byte(fmt.Sprintf("%v", b)))
	return fmt.Sprintf("%x", hash)
}
