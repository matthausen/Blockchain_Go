package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// A blockchain is a series of blocks
type BlockChain struct {
	blocks []*Block
}

// A block contains a hash, data and the previous hash
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Method to create the hash based on previous hash & Data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Func to create a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Method to add block to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Generate genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Initiate blockchain with genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
