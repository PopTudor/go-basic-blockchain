package main

import (
	"fmt"
	"strings"
)

const DIFFICULTY = "0000"

type Blockchain struct {
	chain []Block
	queue []Transaction
	nodes []Node
}

func (b *Blockchain) NewTransaction(transaction Transaction) int {
	b.queue = append(b.queue, transaction)
	return len(b.queue)
}

func (b *Blockchain) Chain() []Block {
	return b.chain
}

func (b *Blockchain) Init() {
	fmt.Printf("init...")
	block := NewBlockFromPrimitive("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks.", b.queue)
	b.ProofOfWork(block)
	b.AddBlock(block)
	fmt.Printf("end init...")
}

func (b *Blockchain) Mine() Block {
	prevHash := b.LastBlock().PrevHash
	bl := NewBlockFromPrimitive(prevHash, b.queue)
	b.ProofOfWork(bl)
	b.AddBlock(bl)
	return *bl
}

func (b *Blockchain) ProofOfWork(block *Block) {
	for nonce := 0; !strings.HasPrefix(block.Hash, DIFFICULTY); nonce++ {
		block.CalculateHash(nonce)
	}
}

func (b *Blockchain) RegisterNode(node Node) {
	b.nodes = append(b.nodes, node)
}

func (b *Blockchain) isValidChain() {
	// todo
}

func (b *Blockchain) resolveFork() {
	// todo
}

func (b *Blockchain) LastBlock() Block {
	return b.chain[b.lastChainIndex()]
}

func (b *Blockchain) AddBlock(block *Block) {
	block.Index = b.nextBlockIndex()
	b.chain = append(b.chain, *block)
}

func (b *Blockchain) lastChainIndex() int {
	return len(b.chain) - 1
}

func (b *Blockchain) nextBlockIndex() int {
	return len(b.chain)
}

func (b *Blockchain) validChain() bool {
	//	for each block in the chain
	//  check that the previousHash of the current block == the hash of the previous block
	return false
}

func (b *Blockchain) resolveConsensus() {
	for node := range b.nodes {
		// make request to node and ask for it's chain
		// convert the json to structure
		// check if the length is greater than our own and if it's valid
		// if both are true, replace ours
	}
}
