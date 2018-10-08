package main

import (
	"fmt"
)

type BlockChain struct {
	Blocks map[string]*Block
	LastID int
	Last   string
}

func (blockchain *BlockChain) Append(data string) {
	block := blockchain.createUnminedBlock(data)
	block.Mine()
	appendBlock(blockchain, block)
}

func (blockchain *BlockChain) Print() {
	for previous := blockchain.Last; previous != ""; {
		currentBlock := blockchain.Blocks[previous]
		fmt.Println(currentBlock)
		previous = currentBlock.Previous
	}
}

func (blockchain *BlockChain) createUnminedBlock(data string) *Block {
	block := Block{
		ID:       blockchain.LastID + 1,
		Data:     data,
		Previous: blockchain.Last,
	}
	return &block
}

func appendBlock(blockchain *BlockChain, block *Block) {
	blockchain.Blocks[block.Hash] = block
	blockchain.Last = block.Hash
	blockchain.LastID = block.ID
}
