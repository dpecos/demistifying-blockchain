package main

import (
	"fmt"
	"time"
)

type BlockChain struct {
	Blocks    map[string]*Block
	LastBlock *Block
}

func (blockchain *BlockChain) Append(data string) {
	block := blockchain.createUnminedBlock(data)
	block.Mine()
	appendBlock(blockchain, block)
}

func (blockchain *BlockChain) Print() {
	for previous := blockchain.LastBlock.Hash; previous != ""; {
		currentBlock := blockchain.Blocks[previous]
		fmt.Printf("%s\n\n", currentBlock)
		previous = currentBlock.Previous
	}
}

func (blockchain *BlockChain) createUnminedBlock(data string) *Block {

	block := Block{
		Timestamp: time.Now(),
		Data:      data,
	}

	if blockchain.LastBlock != nil {
		block.Number = blockchain.LastBlock.Number + 1
		block.Previous = blockchain.LastBlock.Hash
	} else {
		block.Number = 0
	}

	return &block
}

func appendBlock(blockchain *BlockChain, block *Block) {
	blockchain.Blocks[block.Hash] = block
	blockchain.LastBlock = block
}
